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

type ProtocolTemplateGroupObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ProtocolTemplateGroupParameters struct {

	// Name of the protocol template group.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Service template ID list.
	// +crossplane:generate:reference:type=ProtocolTemplate
	// +kubebuilder:validation:Optional
	TemplateIds []*string `json:"templateIds,omitempty" tf:"template_ids,omitempty"`

	// +kubebuilder:validation:Optional
	TemplateIdsRefs []v1.Reference `json:"templateIdsRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	TemplateIdsSelector *v1.Selector `json:"templateIdsSelector,omitempty" tf:"-"`
}

// ProtocolTemplateGroupSpec defines the desired state of ProtocolTemplateGroup
type ProtocolTemplateGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProtocolTemplateGroupParameters `json:"forProvider"`
}

// ProtocolTemplateGroupStatus defines the observed state of ProtocolTemplateGroup.
type ProtocolTemplateGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProtocolTemplateGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ProtocolTemplateGroup is the Schema for the ProtocolTemplateGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type ProtocolTemplateGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProtocolTemplateGroupSpec   `json:"spec"`
	Status            ProtocolTemplateGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProtocolTemplateGroupList contains a list of ProtocolTemplateGroups
type ProtocolTemplateGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProtocolTemplateGroup `json:"items"`
}

// Repository type metadata.
var (
	ProtocolTemplateGroup_Kind             = "ProtocolTemplateGroup"
	ProtocolTemplateGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProtocolTemplateGroup_Kind}.String()
	ProtocolTemplateGroup_KindAPIVersion   = ProtocolTemplateGroup_Kind + "." + CRDGroupVersion.String()
	ProtocolTemplateGroup_GroupVersionKind = CRDGroupVersion.WithKind(ProtocolTemplateGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&ProtocolTemplateGroup{}, &ProtocolTemplateGroupList{})
}
