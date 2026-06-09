package instance

import "testing"

func TestTCRCredentialFromConfigUsesExplicitTokenCredential(t *testing.T) {
	cred, err := tcrCredentialFromConfig("ap-guangzhou", map[string]interface{}{
		"secret_id":      "sid",
		"secret_key":     "skey",
		"security_token": "token",
	})
	if err != nil {
		t.Fatalf("tcrCredentialFromConfig returned error: %v", err)
	}

	gotID, gotKey, gotToken := cred.GetCredential()
	if gotID != "sid" || gotKey != "skey" || gotToken != "token" {
		t.Fatalf("unexpected credential: got id=%q key=%q token=%q", gotID, gotKey, gotToken)
	}
}

func TestTCRCredentialFromConfigRejectsPartialStaticCredential(t *testing.T) {
	_, err := tcrCredentialFromConfig("ap-guangzhou", map[string]interface{}{
		"secret_id": "sid",
	})
	if err == nil {
		t.Fatal("expected error for missing secret_key")
	}
}

func TestSingletonBlockConfigAcceptsObjectAndList(t *testing.T) {
	object := map[string]interface{}{
		"role_arn": "qcs::cam::uin/100000000001:roleName/premium-role",
	}
	if got := singletonBlockConfig(map[string]interface{}{tcrKeyAssumeRole: object}, tcrKeyAssumeRole); got["role_arn"] != object["role_arn"] {
		t.Fatalf("unexpected object block: %#v", got)
	}
	if got := singletonBlockConfig(map[string]interface{}{tcrKeyAssumeRole: []interface{}{object}}, tcrKeyAssumeRole); got["role_arn"] != object["role_arn"] {
		t.Fatalf("unexpected list block: %#v", got)
	}
}

func TestInt64FromMapAcceptsTerraformJSONNumbers(t *testing.T) {
	config := map[string]interface{}{
		"float": float64(3600),
		"int":   1800,
	}
	if got := int64FromMap(config, "float", 7200); got != 3600 {
		t.Fatalf("unexpected float duration: %d", got)
	}
	if got := int64FromMap(config, "int", 7200); got != 1800 {
		t.Fatalf("unexpected int duration: %d", got)
	}
	if got := int64FromMap(config, "missing", 7200); got != 7200 {
		t.Fatalf("unexpected fallback duration: %d", got)
	}
}
