/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ExternalKeyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	KeyState *string `json:"keyState,omitempty" tf:"key_state,omitempty"`
}

type ExternalKeyParameters struct {

	// Name of CMK. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
	// +kubebuilder:validation:Required
	Alias *string `json:"alias" tf:"alias,omitempty"`

	// Description of CMK. The maximum is 1024 bytes.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specify whether to archive key. Default value is `false`. This field is conflict with `is_enabled`, valid when key_state is `Enabled`, `Disabled`, `Archived`.
	// +kubebuilder:validation:Optional
	IsArchived *bool `json:"isArchived,omitempty" tf:"is_archived,omitempty"`

	// Specify whether to enable key. Default value is `false`. This field is conflict with `is_archived`, valid when key_state is `Enabled`, `Disabled`, `Archived`.
	// +kubebuilder:validation:Optional
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled,omitempty"`

	// The base64-encoded key material encrypted with the public_key. For regions using the national secret version, the length of the imported key material is required to be 128 bits, and for regions using the FIPS version, the length of the imported key material is required to be 256 bits.
	// +kubebuilder:validation:Optional
	KeyMaterialBase64SecretRef *v1.SecretKeySelector `json:"keyMaterialBase64SecretRef,omitempty" tf:"-"`

	// Duration in days after which the key is deleted after destruction of the resource, must be between 7 and 30 days. Defaults to 7 days.
	// +kubebuilder:validation:Optional
	PendingDeleteWindowInDays *float64 `json:"pendingDeleteWindowInDays,omitempty" tf:"pending_delete_window_in_days,omitempty"`

	// Tags of CMK.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// This value means the effective timestamp of the key material, 0 means it does not expire. Need to be greater than the current timestamp, the maximum support is 2147443200.
	// +kubebuilder:validation:Optional
	ValidTo *float64 `json:"validTo,omitempty" tf:"valid_to,omitempty"`

	// The algorithm for encrypting key material. Available values include `RSAES_PKCS1_V1_5`, `RSAES_OAEP_SHA_1` and `RSAES_OAEP_SHA_256`. Default value is `RSAES_PKCS1_V1_5`.
	// +kubebuilder:validation:Optional
	WrappingAlgorithm *string `json:"wrappingAlgorithm,omitempty" tf:"wrapping_algorithm,omitempty"`
}

// ExternalKeySpec defines the desired state of ExternalKey
type ExternalKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ExternalKeyParameters `json:"forProvider"`
}

// ExternalKeyStatus defines the observed state of ExternalKey.
type ExternalKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ExternalKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ExternalKey is the Schema for the ExternalKeys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type ExternalKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ExternalKeySpec   `json:"spec"`
	Status            ExternalKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ExternalKeyList contains a list of ExternalKeys
type ExternalKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExternalKey `json:"items"`
}

// Repository type metadata.
var (
	ExternalKey_Kind             = "ExternalKey"
	ExternalKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ExternalKey_Kind}.String()
	ExternalKey_KindAPIVersion   = ExternalKey_Kind + "." + CRDGroupVersion.String()
	ExternalKey_GroupVersionKind = CRDGroupVersion.WithKind(ExternalKey_Kind)
)

func init() {
	SchemeBuilder.Register(&ExternalKey{}, &ExternalKeyList{})
}
