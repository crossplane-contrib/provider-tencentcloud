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

type TmpTkeRecordRuleYamlObservation struct {
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	TemplateID *string `json:"templateId,omitempty" tf:"template_id,omitempty"`

	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type TmpTkeRecordRuleYamlParameters struct {

	// Contents of record rules in yaml format.
	// +kubebuilder:validation:Required
	Content *string `json:"content" tf:"content,omitempty"`

	// Instance Id.
	// +crossplane:generate:reference:type=TmpInstance
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`
}

// TmpTkeRecordRuleYamlSpec defines the desired state of TmpTkeRecordRuleYaml
type TmpTkeRecordRuleYamlSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TmpTkeRecordRuleYamlParameters `json:"forProvider"`
}

// TmpTkeRecordRuleYamlStatus defines the observed state of TmpTkeRecordRuleYaml.
type TmpTkeRecordRuleYamlStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TmpTkeRecordRuleYamlObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TmpTkeRecordRuleYaml is the Schema for the TmpTkeRecordRuleYamls API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type TmpTkeRecordRuleYaml struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TmpTkeRecordRuleYamlSpec   `json:"spec"`
	Status            TmpTkeRecordRuleYamlStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TmpTkeRecordRuleYamlList contains a list of TmpTkeRecordRuleYamls
type TmpTkeRecordRuleYamlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TmpTkeRecordRuleYaml `json:"items"`
}

// Repository type metadata.
var (
	TmpTkeRecordRuleYaml_Kind             = "TmpTkeRecordRuleYaml"
	TmpTkeRecordRuleYaml_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TmpTkeRecordRuleYaml_Kind}.String()
	TmpTkeRecordRuleYaml_KindAPIVersion   = TmpTkeRecordRuleYaml_Kind + "." + CRDGroupVersion.String()
	TmpTkeRecordRuleYaml_GroupVersionKind = CRDGroupVersion.WithKind(TmpTkeRecordRuleYaml_Kind)
)

func init() {
	SchemeBuilder.Register(&TmpTkeRecordRuleYaml{}, &TmpTkeRecordRuleYamlList{})
}
