package instance

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/upjet/pkg/terraform"
	"github.com/pkg/errors"
	tccommon "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	sdkerrors "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	tcr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tcr/v20190924"
	"sigs.k8s.io/controller-runtime/pkg/client"

	v1alpha1 "github.com/crossplane-contrib/provider-tencentcloud/apis/tcr/v1alpha1"
)

const (
	errTCRCleanupSetup     = "cannot build TencentCloud TCR cleanup client"
	errTCRCleanupConfigure = "cannot read TencentCloud provider configuration for TCR cleanup"

	tcrDeleteRetryInterval = 5 * time.Second
	tcrDeleteRetryTimeout  = 2 * time.Minute
)

const (
	tcrKeyAssumeRole                = "assume_role"
	tcrKeyAssumeRoleWithWebIdentity = "assume_role_with_web_identity"
	tcrKeyCAMRoleName               = "cam_role_name"
	tcrKeyEnablePodOIDC             = "enable_pod_oidc"
	tcrKeyRegion                    = "region"
	tcrKeySecretID                  = "secret_id"
	tcrKeySecretKey                 = "secret_key"
	tcrKeySecurityToken             = "security_token"
)

type tcrInstanceCleanupConnector struct {
	base  managed.ExternalConnecter
	kube  client.Client
	setup terraform.SetupFn
	log   logging.Logger
}

type tcrInstanceCleanupExternal struct {
	managed.ExternalClient
	kube  client.Client
	setup terraform.SetupFn
	log   logging.Logger
}

func newTCRInstanceCleanupConnector(base managed.ExternalConnecter, kube client.Client, setup terraform.SetupFn, log logging.Logger) managed.ExternalConnecter {
	return &tcrInstanceCleanupConnector{
		base:  base,
		kube:  kube,
		setup: setup,
		log:   log.WithValues("cleanup", "tcr-instance"),
	}
}

func (c *tcrInstanceCleanupConnector) Connect(ctx context.Context, mg xpresource.Managed) (managed.ExternalClient, error) {
	ec, err := c.base.Connect(ctx, mg)
	if err != nil {
		return nil, err
	}
	return &tcrInstanceCleanupExternal{
		ExternalClient: ec,
		kube:           c.kube,
		setup:          c.setup,
		log:            c.log.WithValues("name", mg.GetName(), "uid", mg.GetUID()),
	}, nil
}

func (e *tcrInstanceCleanupExternal) Delete(ctx context.Context, mg xpresource.Managed) error {
	return e.cleanupAndDelete(ctx, mg)
}

func (e *tcrInstanceCleanupExternal) cleanupAndDelete(ctx context.Context, mg xpresource.Managed) error {
	annotations := mg.GetAnnotations()
	if annotations == nil {
		return nil
	}
	registryID := annotations["crossplane.io/external-name"]
	if registryID == "" {
		return nil
	}

	client, region, err := e.newTCRClient(ctx, mg)
	if err != nil {
		return errors.Wrap(err, errTCRCleanupSetup)
	}
	logger := e.log.WithValues("registryID", registryID, "region", region)

	if err := cleanupTCRTokens(ctx, client, registryID); err != nil {
		return errors.Wrap(err, "cannot cleanup TCR tokens")
	}
	if err := cleanupTCRInternalEndpoints(ctx, client, registryID, region); err != nil {
		return errors.Wrap(err, "cannot cleanup TCR internal endpoints")
	}
	if err := cleanupTCRSecurityPolicies(ctx, client, registryID); err != nil {
		return errors.Wrap(err, "cannot cleanup TCR security policies")
	}
	if err := cleanupTCRRepositories(ctx, client, registryID); err != nil {
		return errors.Wrap(err, "cannot cleanup TCR repositories")
	}
	if err := cleanupTCRNamespaces(ctx, client, registryID); err != nil {
		return errors.Wrap(err, "cannot cleanup TCR namespaces")
	}

	logger.Debug("TCR instance dependencies cleanup completed")
	if err := deleteTCRInstanceAndWait(ctx, client, registryID, tcrDeleteBucket(mg)); err != nil {
		return errors.Wrap(err, "cannot delete TCR instance")
	}
	return nil
}

func (e *tcrInstanceCleanupExternal) newTCRClient(ctx context.Context, mg xpresource.Managed) (*tcr.Client, string, error) {
	ts, err := e.setup(ctx, e.kube, mg)
	if err != nil {
		return nil, "", err
	}

	region, ok := stringConfig(ts.Configuration, tcrKeyRegion)
	if !ok {
		return nil, "", errors.New(errTCRCleanupConfigure + ": missing region")
	}
	credential, err := tcrCredentialFromConfig(region, ts.Configuration)
	if err != nil {
		return nil, "", err
	}
	clientProfile := profile.NewClientProfile()
	clientProfile.HttpProfile.ReqTimeout = 60

	client, err := tcr.NewClient(credential, region, clientProfile)
	return client, region, err
}

func tcrCredentialFromConfig(region string, config map[string]interface{}) (tccommon.CredentialIface, error) {
	if oidcEnabled(config) {
		provider, err := tccommon.DefaultTkeOIDCRoleArnProvider()
		if err != nil {
			return nil, err
		}
		return provider.GetCredential()
	}
	if block := singletonBlockConfig(config, tcrKeyAssumeRoleWithWebIdentity); block != nil {
		return credentialFromAssumeRoleWithWebIdentity(region, block)
	}
	if block := singletonBlockConfig(config, tcrKeyAssumeRole); block != nil {
		return credentialFromAssumeRole(config, block)
	}
	if secretID, hasSecretID := stringConfig(config, tcrKeySecretID); hasSecretID {
		secretKey, hasSecretKey := stringConfig(config, tcrKeySecretKey)
		if !hasSecretKey {
			return nil, errors.New(errTCRCleanupConfigure + ": missing secret_key")
		}
		token, _ := stringConfig(config, tcrKeySecurityToken)
		if token != "" {
			return tccommon.NewTokenCredential(secretID, secretKey, token), nil
		}
		return tccommon.NewCredential(secretID, secretKey), nil
	}
	if _, hasSecretKey := stringConfig(config, tcrKeySecretKey); hasSecretKey {
		return nil, errors.New(errTCRCleanupConfigure + ": missing secret_id")
	}

	if camRoleName, ok := stringConfig(config, tcrKeyCAMRoleName); ok {
		return tccommon.NewCvmRoleProvider(camRoleName).GetCredential()
	}
	return tccommon.DefaultProviderChain().GetCredential()
}

func credentialFromAssumeRole(config map[string]interface{}, block map[string]interface{}) (tccommon.CredentialIface, error) {
	secretID, ok := stringConfig(config, tcrKeySecretID)
	if !ok {
		return nil, errors.New(errTCRCleanupConfigure + ": assume_role requires secret_id")
	}
	secretKey, ok := stringConfig(config, tcrKeySecretKey)
	if !ok {
		return nil, errors.New(errTCRCleanupConfigure + ": assume_role requires secret_key")
	}
	roleARN, ok := stringFromMap(block, "role_arn")
	if !ok {
		return nil, errors.New(errTCRCleanupConfigure + ": assume_role requires role_arn")
	}
	sessionName, ok := stringFromMap(block, "session_name")
	if !ok {
		sessionName = "premium-crossplane-tcr-cleanup"
	}
	duration := int64FromMap(block, "session_duration", 7200)
	return tccommon.NewRoleArnProvider(secretID, secretKey, roleARN, sessionName, duration).GetCredential()
}

func credentialFromAssumeRoleWithWebIdentity(region string, block map[string]interface{}) (tccommon.CredentialIface, error) {
	providerID, ok := stringFromMap(block, "provider_id")
	if !ok {
		return nil, errors.New(errTCRCleanupConfigure + ": assume_role_with_web_identity requires provider_id")
	}
	roleARN, ok := stringFromMap(block, "role_arn")
	if !ok {
		return nil, errors.New(errTCRCleanupConfigure + ": assume_role_with_web_identity requires role_arn")
	}
	sessionName, ok := stringFromMap(block, "session_name")
	if !ok {
		sessionName = "premium-crossplane-tcr-cleanup"
	}
	token, ok := stringFromMap(block, "web_identity_token")
	if !ok {
		tokenFile, ok := stringFromMap(block, "web_identity_token_file")
		if !ok {
			return nil, errors.New(errTCRCleanupConfigure + ": assume_role_with_web_identity requires web_identity_token or web_identity_token_file")
		}
		tokenBytes, err := os.ReadFile(tokenFile)
		if err != nil {
			return nil, err
		}
		token = string(tokenBytes)
	}
	duration := int64FromMap(block, "duration_seconds", 7200)
	return tccommon.NewOIDCRoleArnProvider(region, providerID, token, roleARN, sessionName, duration).GetCredential()
}

func stringConfig(config map[string]interface{}, key string) (string, bool) {
	v, ok := config[key]
	if !ok || v == nil {
		return "", false
	}
	s, ok := v.(string)
	if !ok || s == "" {
		return "", false
	}
	return s, true
}

func oidcEnabled(config map[string]interface{}) bool {
	v, ok := config[tcrKeyEnablePodOIDC]
	if !ok || v == nil {
		return false
	}
	enabled, ok := v.(bool)
	return ok && enabled
}

func singletonBlockConfig(config map[string]interface{}, key string) map[string]interface{} {
	v, ok := config[key]
	if !ok || v == nil {
		return nil
	}
	if block, ok := v.(map[string]interface{}); ok {
		return block
	}
	list, ok := v.([]interface{})
	if !ok || len(list) == 0 || list[0] == nil {
		return nil
	}
	block, _ := list[0].(map[string]interface{})
	return block
}

func stringFromMap(config map[string]interface{}, key string) (string, bool) {
	v, ok := config[key]
	if !ok || v == nil {
		return "", false
	}
	s, ok := v.(string)
	if !ok || s == "" {
		return "", false
	}
	return s, true
}

func int64FromMap(config map[string]interface{}, key string, fallback int64) int64 {
	v, ok := config[key]
	if !ok || v == nil {
		return fallback
	}
	switch n := v.(type) {
	case int64:
		return n
	case int:
		return int64(n)
	case float64:
		return int64(n)
	default:
		return fallback
	}
}

func cleanupTCRTokens(ctx context.Context, c *tcr.Client, registryID string) error {
	var offset int64
	for {
		req := tcr.NewDescribeInstanceTokenRequest()
		req.RegistryId = stringPtr(registryID)
		req.Offset = int64Ptr(offset)
		req.Limit = int64Ptr(100)
		resp, err := c.DescribeInstanceTokenWithContext(ctx, req)
		if ignoreTencentCloudNotFound(err) {
			return nil
		}
		if err != nil {
			return err
		}
		tokens := resp.Response.Tokens
		for _, token := range tokens {
			if token == nil || token.Id == nil {
				continue
			}
			del := tcr.NewDeleteInstanceTokenRequest()
			del.RegistryId = stringPtr(registryID)
			del.TokenId = token.Id
			if err := retryTCR(ctx, func() error {
				_, err := c.DeleteInstanceTokenWithContext(ctx, del)
				return err
			}); err != nil && !ignoreTencentCloudNotFound(err) {
				return err
			}
		}
		if len(tokens) < 100 {
			return nil
		}
		offset += 100
	}
}

func cleanupTCRInternalEndpoints(ctx context.Context, c *tcr.Client, registryID, region string) error {
	req := tcr.NewDescribeInternalEndpointsRequest()
	req.RegistryId = stringPtr(registryID)
	resp, err := c.DescribeInternalEndpointsWithContext(ctx, req)
	if ignoreTencentCloudNotFound(err) {
		return nil
	}
	if err != nil {
		return err
	}
	for _, access := range resp.Response.AccessVpcSet {
		if access == nil || access.VpcId == nil || access.SubnetId == nil {
			continue
		}
		del := tcr.NewManageInternalEndpointRequest()
		del.RegistryId = stringPtr(registryID)
		del.Operation = stringPtr("Delete")
		del.VpcId = access.VpcId
		del.SubnetId = access.SubnetId
		del.RegionName = stringPtr(region)
		if err := retryTCR(ctx, func() error {
			_, err := c.ManageInternalEndpointWithContext(ctx, del)
			return err
		}); err != nil && !ignoreTencentCloudNotFound(err) {
			return err
		}
	}
	return nil
}

func cleanupTCRSecurityPolicies(ctx context.Context, c *tcr.Client, registryID string) error {
	req := tcr.NewDescribeSecurityPoliciesRequest()
	req.RegistryId = stringPtr(registryID)
	resp, err := c.DescribeSecurityPoliciesWithContext(ctx, req)
	if ignoreTencentCloudNotFound(err) {
		return nil
	}
	if err != nil {
		return err
	}
	for _, policy := range resp.Response.SecurityPolicySet {
		if policy == nil {
			continue
		}
		del := tcr.NewDeleteSecurityPolicyRequest()
		del.RegistryId = stringPtr(registryID)
		del.PolicyIndex = policy.PolicyIndex
		del.PolicyVersion = policy.PolicyVersion
		del.CidrBlock = policy.CidrBlock
		if err := retryTCR(ctx, func() error {
			_, err := c.DeleteSecurityPolicyWithContext(ctx, del)
			return err
		}); err != nil && !ignoreTencentCloudNotFound(err) {
			return err
		}
	}
	return nil
}

func cleanupTCRRepositories(ctx context.Context, c *tcr.Client, registryID string) error {
	var offset int64
	for {
		req := tcr.NewDescribeRepositoriesRequest()
		req.RegistryId = stringPtr(registryID)
		req.Offset = int64Ptr(offset)
		req.Limit = int64Ptr(100)
		resp, err := c.DescribeRepositoriesWithContext(ctx, req)
		if ignoreTencentCloudNotFound(err) {
			return nil
		}
		if err != nil {
			return err
		}
		repositories := resp.Response.RepositoryList
		for _, repo := range repositories {
			namespace, name := repoNamespaceAndName(repo)
			if namespace == "" || name == "" {
				continue
			}
			del := tcr.NewDeleteRepositoryRequest()
			del.RegistryId = stringPtr(registryID)
			del.NamespaceName = stringPtr(namespace)
			del.RepositoryName = stringPtr(name)
			if err := retryTCR(ctx, func() error {
				_, err := c.DeleteRepositoryWithContext(ctx, del)
				return err
			}); err != nil && !ignoreTencentCloudNotFound(err) {
				return err
			}
		}
		if len(repositories) < 100 {
			return nil
		}
		offset += 100
	}
}

func cleanupTCRNamespaces(ctx context.Context, c *tcr.Client, registryID string) error {
	deadline := time.Now().Add(tcrDeleteRetryTimeout)
	for {
		namespaces, err := listTCRNamespaces(ctx, c, registryID)
		if ignoreTencentCloudNotFound(err) {
			return nil
		}
		if err != nil {
			return err
		}
		if len(namespaces) == 0 {
			return nil
		}
		for _, ns := range namespaces {
			if ns == nil || ns.Name == nil || *ns.Name == "" {
				continue
			}
			del := tcr.NewDeleteNamespaceRequest()
			del.RegistryId = stringPtr(registryID)
			del.NamespaceName = ns.Name
			if err := retryTCR(ctx, func() error {
				_, err := c.DeleteNamespaceWithContext(ctx, del)
				return err
			}); err != nil && !ignoreTencentCloudNotFound(err) {
				return err
			}
		}
		if time.Now().After(deadline) {
			return fmt.Errorf("timed out waiting for TCR namespaces to be deleted")
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(tcrDeleteRetryInterval):
		}
	}
}

func deleteTCRInstanceAndWait(ctx context.Context, c *tcr.Client, registryID string, deleteBucket bool) error {
	req := tcr.NewDeleteInstanceRequest()
	req.RegistryId = stringPtr(registryID)
	req.DeleteBucket = boolPtr(deleteBucket)
	if err := retryTCR(ctx, func() error {
		_, err := c.DeleteInstanceWithContext(ctx, req)
		return err
	}); err != nil && !ignoreTencentCloudNotFound(err) {
		return err
	}

	deadline := time.Now().Add(tcrDeleteRetryTimeout)
	for {
		exists, err := tcrInstanceExists(ctx, c, registryID)
		if ignoreTencentCloudNotFound(err) {
			return nil
		}
		if err != nil {
			return err
		}
		if !exists {
			return nil
		}
		if time.Now().After(deadline) {
			return fmt.Errorf("timed out waiting for TCR instance %s to be deleted", registryID)
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(tcrDeleteRetryInterval):
		}
	}
}

func tcrInstanceExists(ctx context.Context, c *tcr.Client, registryID string) (bool, error) {
	req := tcr.NewDescribeInstancesRequest()
	req.Registryids = []*string{stringPtr(registryID)}
	req.Limit = int64Ptr(20)
	resp, err := c.DescribeInstancesWithContext(ctx, req)
	if err != nil {
		return false, err
	}
	return len(resp.Response.Registries) > 0, nil
}

func tcrDeleteBucket(mg xpresource.Managed) bool {
	inst, ok := mg.(*v1alpha1.Instance)
	if !ok || inst.Spec.ForProvider.DeleteBucket == nil {
		return false
	}
	return *inst.Spec.ForProvider.DeleteBucket
}

func listTCRNamespaces(ctx context.Context, c *tcr.Client, registryID string) ([]*tcr.TcrNamespaceInfo, error) {
	var all []*tcr.TcrNamespaceInfo
	var offset int64
	for {
		req := tcr.NewDescribeNamespacesRequest()
		req.RegistryId = stringPtr(registryID)
		req.Offset = int64Ptr(offset)
		req.Limit = int64Ptr(100)
		req.All = boolPtr(true)
		resp, err := c.DescribeNamespacesWithContext(ctx, req)
		if err != nil {
			return nil, err
		}
		namespaces := resp.Response.NamespaceList
		all = append(all, namespaces...)
		if len(namespaces) < 100 {
			return all, nil
		}
		offset += 100
	}
}

func repoNamespaceAndName(repo *tcr.TcrRepositoryInfo) (string, string) {
	if repo == nil || repo.Name == nil {
		return "", ""
	}
	namespace := ""
	if repo.Namespace != nil {
		namespace = *repo.Namespace
	}
	name := *repo.Name
	if namespace != "" {
		name = strings.TrimPrefix(name, namespace+"/")
		return namespace, name
	}
	parts := strings.SplitN(name, "/", 2)
	if len(parts) != 2 {
		return "", name
	}
	return parts[0], parts[1]
}

func retryTCR(ctx context.Context, fn func() error) error {
	deadline := time.Now().Add(tcrDeleteRetryTimeout)
	for {
		err := fn()
		if err == nil || ignoreTencentCloudNotFound(err) {
			return nil
		}
		if !isTencentCloudRetryable(err) || time.Now().After(deadline) {
			return err
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(tcrDeleteRetryInterval):
		}
	}
}

func isTencentCloudRetryable(err error) bool {
	te, ok := err.(*sdkerrors.TencentCloudSDKError)
	if !ok {
		return false
	}
	code := te.Code
	return strings.HasPrefix(code, "InternalError") ||
		strings.HasPrefix(code, "FailedOperation") ||
		code == "RequestLimitExceeded" ||
		code == "ResourceUnavailable"
}

func ignoreTencentCloudNotFound(err error) bool {
	if err == nil {
		return false
	}
	te, ok := err.(*sdkerrors.TencentCloudSDKError)
	if !ok {
		return false
	}
	code := te.Code
	return strings.HasPrefix(code, "ResourceNotFound") ||
		strings.Contains(code, "ErrNoNamespace") ||
		strings.Contains(code, "ErrNoRepository")
}

func stringPtr(v string) *string {
	return &v
}

func int64Ptr(v int64) *int64 {
	return &v
}

func boolPtr(v bool) *bool {
	return &v
}
