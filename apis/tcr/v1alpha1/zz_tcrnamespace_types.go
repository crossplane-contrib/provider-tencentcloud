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

type CveWhitelistItemsObservation struct {
}

type CveWhitelistItemsParameters struct {

	// Vulnerability Whitelist ID.
	// +kubebuilder:validation:Optional
	CveID *string `json:"cveId,omitempty" tf:"cve_id,omitempty"`
}

type TcrNamespaceObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TcrNamespaceParameters struct {

	// Vulnerability Whitelist.
	// +kubebuilder:validation:Optional
	CveWhitelistItems []CveWhitelistItemsParameters `json:"cveWhitelistItems,omitempty" tf:"cve_whitelist_items,omitempty"`

	// ID of the TCR instance.
	// +kubebuilder:validation:Required
	InstanceID *string `json:"instanceId" tf:"instance_id,omitempty"`

	// Scanning level, `True` is automatic, `False` is manual. Default is `false`.
	// +kubebuilder:validation:Optional
	IsAutoScan *bool `json:"isAutoScan,omitempty" tf:"is_auto_scan,omitempty"`

	// Blocking switch, `True` is open, `False` is closed. Default is `false`.
	// +kubebuilder:validation:Optional
	IsPreventVul *bool `json:"isPreventVul,omitempty" tf:"is_prevent_vul,omitempty"`

	// Indicate that the namespace is public or not. Default is `false`.
	// +kubebuilder:validation:Optional
	IsPublic *bool `json:"isPublic,omitempty" tf:"is_public,omitempty"`

	// Name of the TCR namespace. Valid length is [2~30]. It can only contain lowercase letters, numbers and separators (`.`, `_`, `-`), and cannot start, end or continue with separators.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Block vulnerability level, currently only supports `low`, `medium`, `high`.
	// +kubebuilder:validation:Optional
	Severity *string `json:"severity,omitempty" tf:"severity,omitempty"`
}

// TcrNamespaceSpec defines the desired state of TcrNamespace
type TcrNamespaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TcrNamespaceParameters `json:"forProvider"`
}

// TcrNamespaceStatus defines the observed state of TcrNamespace.
type TcrNamespaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TcrNamespaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TcrNamespace is the Schema for the TcrNamespaces API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type TcrNamespace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TcrNamespaceSpec   `json:"spec"`
	Status            TcrNamespaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TcrNamespaceList contains a list of TcrNamespaces
type TcrNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TcrNamespace `json:"items"`
}

// Repository type metadata.
var (
	TcrNamespace_Kind             = "TcrNamespace"
	TcrNamespace_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TcrNamespace_Kind}.String()
	TcrNamespace_KindAPIVersion   = TcrNamespace_Kind + "." + CRDGroupVersion.String()
	TcrNamespace_GroupVersionKind = CRDGroupVersion.WithKind(TcrNamespace_Kind)
)

func init() {
	SchemeBuilder.Register(&TcrNamespace{}, &TcrNamespaceList{})
}
