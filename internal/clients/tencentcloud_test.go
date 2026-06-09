/*
Copyright 2026 PingCAP.
*/

package clients

import (
	"context"
	"reflect"
	"testing"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	"github.com/crossplane-contrib/provider-tencentcloud/apis/v1beta1"
	vpcv1alpha1 "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1"
)

func TestTerraformSetupBuilderIncludesSecurityToken(t *testing.T) {
	setup := setupWithCredentials(t, `{"secret_id":"sid","secret_key":"skey","security_token":"token","region":"ap-guangzhou"}`)

	want := map[string]interface{}{
		keySecretID:      "sid",
		keySecretKey:     "skey",
		keySecurityToken: "token",
		keyRegion:        "ap-guangzhou",
	}
	assertConfiguration(t, setup.Configuration, want)
}

func TestTerraformSetupBuilderPassesThroughProviderAuthenticationConfiguration(t *testing.T) {
	setup := setupWithCredentials(t, `{
		"region": "ap-guangzhou",
		"profile": "premium-prod",
		"shared_credentials_dir": "/var/run/tencentcloud/credentials",
		"cam_role_name": "premium-cvm-role",
		"enable_pod_oidc": true,
		"allowed_account_ids": ["100000000001"],
		"domain": "tencentcloudapi.com",
		"protocol": "HTTPS",
		"cos_domain": "https://cos.ap-guangzhou.myqcloud.com",
		"ignored": "not-forwarded"
	}`)

	want := map[string]interface{}{
		keyRegion:               "ap-guangzhou",
		keyProfile:              "premium-prod",
		keySharedCredentialsDir: "/var/run/tencentcloud/credentials",
		keyCAMRoleName:          "premium-cvm-role",
		keyEnablePodOIDC:        true,
		keyAllowedAccountIDs:    []interface{}{"100000000001"},
		keyDomain:               "tencentcloudapi.com",
		keyProtocol:             "HTTPS",
		keyCOSDomain:            "https://cos.ap-guangzhou.myqcloud.com",
	}
	assertConfiguration(t, setup.Configuration, want)
}

func TestTerraformSetupBuilderPassesThroughAssumeRole(t *testing.T) {
	setup := setupWithCredentials(t, `{
		"secret_id": "sid",
		"secret_key": "skey",
		"assume_role": {
			"role_arn": "qcs::cam::uin/100000000001:roleName/premium-role",
			"session_name": "premium-crossplane",
			"session_duration": 7200,
			"external_id": "premium"
		}
	}`)

	want := map[string]interface{}{
		keySecretID:  "sid",
		keySecretKey: "skey",
		keyAssumeRole: []interface{}{
			map[string]interface{}{
				"role_arn":         "qcs::cam::uin/100000000001:roleName/premium-role",
				"session_name":     "premium-crossplane",
				"session_duration": float64(7200),
				"external_id":      "premium",
			},
		},
	}
	assertConfiguration(t, setup.Configuration, want)
}

func TestTerraformSetupBuilderPassesThroughAssumeRoleWithWebIdentity(t *testing.T) {
	setup := setupWithCredentials(t, `{
		"region": "ap-guangzhou",
		"assume_role_with_web_identity": [{
			"provider_id": "OIDC",
			"role_arn": "qcs::cam::uin/100000000001:roleName/premium-oidc-role",
			"session_name": "premium-crossplane",
			"web_identity_token_file": "/var/run/secrets/tke-oidc/token"
		}]
	}`)

	want := map[string]interface{}{
		keyRegion: "ap-guangzhou",
		keyAssumeRoleWithWebIdentity: []interface{}{
			map[string]interface{}{
				"provider_id":             "OIDC",
				"role_arn":                "qcs::cam::uin/100000000001:roleName/premium-oidc-role",
				"session_name":            "premium-crossplane",
				"web_identity_token_file": "/var/run/secrets/tke-oidc/token",
			},
		},
	}
	assertConfiguration(t, setup.Configuration, want)
}

func TestTerraformSetupBuilderPassesThroughSAMLAndMFA(t *testing.T) {
	setup := setupWithCredentials(t, `{
		"assume_role_with_saml": {
			"principal_arn": "qcs::cam::uin/100000000001:saml-provider/premium",
			"role_arn": "qcs::cam::uin/100000000001:roleName/premium-saml-role",
			"saml_assertion": "base64-assertion",
			"session_name": "premium-crossplane"
		},
		"mfa_certification": {
			"serial_number": "qcs::cam::uin/100000000001::mfa/softToken",
			"token_code": "123456",
			"duration_seconds": 1800
		}
	}`)

	want := map[string]interface{}{
		keyAssumeRoleWithSAML: []interface{}{
			map[string]interface{}{
				"principal_arn":  "qcs::cam::uin/100000000001:saml-provider/premium",
				"role_arn":       "qcs::cam::uin/100000000001:roleName/premium-saml-role",
				"saml_assertion": "base64-assertion",
				"session_name":   "premium-crossplane",
			},
		},
		keyMFACertification: []interface{}{
			map[string]interface{}{
				"serial_number":    "qcs::cam::uin/100000000001::mfa/softToken",
				"token_code":       "123456",
				"duration_seconds": float64(1800),
			},
		},
	}
	assertConfiguration(t, setup.Configuration, want)
}

func setupWithCredentials(t *testing.T, credentialsJSON string) struct{ Configuration map[string]interface{} } {
	t.Helper()

	const (
		providerConfigName = "default"
		credentialsEnv     = "TEST_TENCENTCLOUD_CREDS"
	)

	t.Setenv(credentialsEnv, credentialsJSON)

	scheme := runtime.NewScheme()
	if err := v1beta1.SchemeBuilder.AddToScheme(scheme); err != nil {
		t.Fatalf("cannot add providerconfig scheme: %v", err)
	}
	if err := vpcv1alpha1.AddToScheme(scheme); err != nil {
		t.Fatalf("cannot add vpc scheme: %v", err)
	}

	pc := &v1beta1.ProviderConfig{
		ObjectMeta: metav1.ObjectMeta{Name: providerConfigName},
		Spec: v1beta1.ProviderConfigSpec{
			Credentials: v1beta1.ProviderCredentials{
				Source: xpv1.CredentialsSourceEnvironment,
				CommonCredentialSelectors: xpv1.CommonCredentialSelectors{
					Env: &xpv1.EnvSelector{Name: credentialsEnv},
				},
			},
		},
	}
	mg := &vpcv1alpha1.VPC{
		TypeMeta: metav1.TypeMeta{
			APIVersion: vpcv1alpha1.CRDGroupVersion.String(),
			Kind:       vpcv1alpha1.VPC_Kind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-vpc",
			UID:  types.UID("test-vpc-uid"),
		},
		Spec: vpcv1alpha1.VPCSpec{
			ResourceSpec: xpv1.ResourceSpec{
				ProviderConfigReference: &xpv1.Reference{Name: providerConfigName},
			},
		},
	}

	client := fake.NewClientBuilder().
		WithScheme(scheme).
		WithObjects(pc).
		Build()
	setup, err := TerraformSetupBuilder("1.2.1", "tencentcloudstack/tencentcloud", "1.81.185")(context.Background(), client, mg)
	if err != nil {
		t.Fatalf("TerraformSetupBuilder returned error: %v", err)
	}
	return struct{ Configuration map[string]interface{} }{Configuration: setup.Configuration}
}

func assertConfiguration(t *testing.T, got, want map[string]interface{}) {
	t.Helper()

	for key, wantValue := range want {
		if gotValue := got[key]; !reflect.DeepEqual(gotValue, wantValue) {
			t.Fatalf("unexpected terraform provider configuration %q: want %#v, got %#v", key, wantValue, gotValue)
		}
	}
	if gotLen, wantLen := len(got), len(want); gotLen != wantLen {
		t.Fatalf("unexpected terraform provider configuration length: want %d, got %d; got %#v", wantLen, gotLen, got)
	}
}
