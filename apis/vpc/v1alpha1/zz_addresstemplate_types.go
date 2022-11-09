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

type AddressTemplateObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AddressTemplateParameters struct {

	// Address list. IP(`10.0.0.1`), CIDR(`10.0.1.0/24`), IP range(`10.0.0.1-10.0.0.100`) format are supported.
	// +kubebuilder:validation:Required
	Addresses []*string `json:"addresses" tf:"addresses,omitempty"`

	// Name of the address template.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

// AddressTemplateSpec defines the desired state of AddressTemplate
type AddressTemplateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AddressTemplateParameters `json:"forProvider"`
}

// AddressTemplateStatus defines the observed state of AddressTemplate.
type AddressTemplateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AddressTemplateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AddressTemplate is the Schema for the AddressTemplates API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type AddressTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AddressTemplateSpec   `json:"spec"`
	Status            AddressTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AddressTemplateList contains a list of AddressTemplates
type AddressTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AddressTemplate `json:"items"`
}

// Repository type metadata.
var (
	AddressTemplate_Kind             = "AddressTemplate"
	AddressTemplate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AddressTemplate_Kind}.String()
	AddressTemplate_KindAPIVersion   = AddressTemplate_Kind + "." + CRDGroupVersion.String()
	AddressTemplate_GroupVersionKind = CRDGroupVersion.WithKind(AddressTemplate_Kind)
)

func init() {
	SchemeBuilder.Register(&AddressTemplate{}, &AddressTemplateList{})
}