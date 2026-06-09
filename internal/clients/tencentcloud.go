/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/upjet/pkg/terraform"

	"github.com/crossplane-contrib/provider-tencentcloud/apis/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal tencentcloud credentials as JSON"
)

const (
	keyAllowedAccountIDs         = "allowed_account_ids"
	keyAssumeRole                = "assume_role"
	keyAssumeRoleWithSAML        = "assume_role_with_saml"
	keyAssumeRoleWithWebIdentity = "assume_role_with_web_identity"
	keyCAMRoleName               = "cam_role_name"
	keyCOSDomain                 = "cos_domain"
	keyDomain                    = "domain"
	keyEnablePodOIDC             = "enable_pod_oidc"
	keyForbiddenAccountIDs       = "forbidden_account_ids"
	keyMFACertification          = "mfa_certification"
	keyProfile                   = "profile"
	keyProtocol                  = "protocol"
	keyRegion                    = "region"
	keySecretID                  = "secret_id"
	keySecretKey                 = "secret_key"
	keySecurityToken             = "security_token"
	keySharedCredentialsDir      = "shared_credentials_dir"
)

var providerConfigurationKeys = map[string]struct{}{
	keyAllowedAccountIDs:         {},
	keyAssumeRole:                {},
	keyAssumeRoleWithSAML:        {},
	keyAssumeRoleWithWebIdentity: {},
	keyCAMRoleName:               {},
	keyCOSDomain:                 {},
	keyDomain:                    {},
	keyEnablePodOIDC:             {},
	keyForbiddenAccountIDs:       {},
	keyMFACertification:          {},
	keyProfile:                   {},
	keyProtocol:                  {},
	keyRegion:                    {},
	keySecretID:                  {},
	keySecretKey:                 {},
	keySecurityToken:             {},
	keySharedCredentialsDir:      {},
}

var singletonBlockKeys = map[string]struct{}{
	keyAssumeRole:                {},
	keyAssumeRoleWithSAML:        {},
	keyAssumeRoleWithWebIdentity: {},
	keyMFACertification:          {},
}

func buildTerraformProviderConfiguration(creds map[string]interface{}) map[string]interface{} {
	config := make(map[string]interface{}, len(creds))
	for key, value := range creds {
		if _, ok := providerConfigurationKeys[key]; !ok {
			continue
		}
		if _, ok := singletonBlockKeys[key]; ok {
			config[key] = normalizeSingletonBlock(value)
			continue
		}
		config[key] = value
	}
	return config
}

func normalizeSingletonBlock(value interface{}) interface{} {
	switch v := value.(type) {
	case []interface{}:
		return v
	case map[string]interface{}:
		return []interface{}{v}
	default:
		return value
	}
}

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		creds := map[string]interface{}{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		ps.Configuration = buildTerraformProviderConfiguration(creds)
		return ps, nil
	}
}
