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

type ProtocolTemplateObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ProtocolTemplateParameters struct {

	// Name of the protocol template.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Protocol list. Valid protocols are  `tcp`, `udp`, `icmp`, `gre`. Single port(tcp:80), multi-port(tcp:80,443), port range(tcp:3306-20000), all(tcp:all) format are support. Protocol `icmp` and `gre` cannot specify port.
	// +kubebuilder:validation:Required
	Protocols []*string `json:"protocols" tf:"protocols,omitempty"`
}

// ProtocolTemplateSpec defines the desired state of ProtocolTemplate
type ProtocolTemplateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProtocolTemplateParameters `json:"forProvider"`
}

// ProtocolTemplateStatus defines the observed state of ProtocolTemplate.
type ProtocolTemplateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProtocolTemplateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ProtocolTemplate is the Schema for the ProtocolTemplates API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type ProtocolTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProtocolTemplateSpec   `json:"spec"`
	Status            ProtocolTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProtocolTemplateList contains a list of ProtocolTemplates
type ProtocolTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProtocolTemplate `json:"items"`
}

// Repository type metadata.
var (
	ProtocolTemplate_Kind             = "ProtocolTemplate"
	ProtocolTemplate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProtocolTemplate_Kind}.String()
	ProtocolTemplate_KindAPIVersion   = ProtocolTemplate_Kind + "." + CRDGroupVersion.String()
	ProtocolTemplate_GroupVersionKind = CRDGroupVersion.WithKind(ProtocolTemplate_Kind)
)

func init() {
	SchemeBuilder.Register(&ProtocolTemplate{}, &ProtocolTemplateList{})
}
