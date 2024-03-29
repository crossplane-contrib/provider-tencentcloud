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

type RetentionPolicyObservation struct {
}

type RetentionPolicyParameters struct {

	// the size of message to retain.
	// +kubebuilder:validation:Optional
	SizeInMb *float64 `json:"sizeInMb,omitempty" tf:"size_in_mb,omitempty"`

	// the time of message to retain.
	// +kubebuilder:validation:Optional
	TimeInMinutes *float64 `json:"timeInMinutes,omitempty" tf:"time_in_minutes,omitempty"`
}

type TdmqNamespaceObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TdmqNamespaceParameters struct {

	// The Dedicated Cluster Id.
	// +crossplane:generate:reference:type=Instance
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// The name of namespace to be created.
	// +kubebuilder:validation:Required
	EnvironName *string `json:"environName" tf:"environ_name,omitempty"`

	// The expiration time of unconsumed message.
	// +kubebuilder:validation:Required
	MsgTTL *float64 `json:"msgTtl" tf:"msg_ttl,omitempty"`

	// Description of the namespace.
	// +kubebuilder:validation:Optional
	Remark *string `json:"remark,omitempty" tf:"remark,omitempty"`

	// The Policy of message to retain. Format like: `{time_in_minutes: Int, size_in_mb: Int}`. `time_in_minutes`: the time of message to retain; `size_in_mb`: the size of message to retain.
	// +kubebuilder:validation:Optional
	RetentionPolicy []RetentionPolicyParameters `json:"retentionPolicy,omitempty" tf:"retention_policy,omitempty"`
}

// TdmqNamespaceSpec defines the desired state of TdmqNamespace
type TdmqNamespaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TdmqNamespaceParameters `json:"forProvider"`
}

// TdmqNamespaceStatus defines the observed state of TdmqNamespace.
type TdmqNamespaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TdmqNamespaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TdmqNamespace is the Schema for the TdmqNamespaces API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type TdmqNamespace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TdmqNamespaceSpec   `json:"spec"`
	Status            TdmqNamespaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TdmqNamespaceList contains a list of TdmqNamespaces
type TdmqNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TdmqNamespace `json:"items"`
}

// Repository type metadata.
var (
	TdmqNamespace_Kind             = "TdmqNamespace"
	TdmqNamespace_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TdmqNamespace_Kind}.String()
	TdmqNamespace_KindAPIVersion   = TdmqNamespace_Kind + "." + CRDGroupVersion.String()
	TdmqNamespace_GroupVersionKind = CRDGroupVersion.WithKind(TdmqNamespace_Kind)
)

func init() {
	SchemeBuilder.Register(&TdmqNamespace{}, &TdmqNamespaceList{})
}
