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

type RecordRulesObservation struct {
}

type RecordRulesParameters struct {

	// Config.
	// +kubebuilder:validation:Required
	Config *string `json:"config" tf:"config,omitempty"`

	// Name.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Used for the argument, if the configuration comes to the template, the template id.
	// +kubebuilder:validation:Optional
	TemplateID *string `json:"templateId,omitempty" tf:"template_id,omitempty"`
}

type TemplateObservation struct {
}

type TemplateParameters struct {

	// Template description.
	// +kubebuilder:validation:Optional
	Describe *string `json:"describe,omitempty" tf:"describe,omitempty"`

	// Whether the system-supplied default template is used for outgoing references.
	// +kubebuilder:validation:Optional
	IsDefault *bool `json:"isDefault,omitempty" tf:"is_default,omitempty"`

	// Template dimensions, the following types are supported `instance` instance level, `cluster` cluster level.
	// +kubebuilder:validation:Required
	Level *string `json:"level" tf:"level,omitempty"`

	// Template name.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Effective when Level is a cluster, A list of PodMonitors rules in the template.
	// +kubebuilder:validation:Optional
	PodMonitors []TemplatePodMonitorsParameters `json:"podMonitors,omitempty" tf:"pod_monitors,omitempty"`

	// Effective when Level is a cluster, A list of RawJobs rules in the template.
	// +kubebuilder:validation:Optional
	RawJobs []TemplateRawJobsParameters `json:"rawJobs,omitempty" tf:"raw_jobs,omitempty"`

	// Effective when Level is instance, A list of aggregation rules in the template.
	// +kubebuilder:validation:Optional
	RecordRules []RecordRulesParameters `json:"recordRules,omitempty" tf:"record_rules,omitempty"`

	// Effective when Level is a cluster, A list of ServiceMonitor rules in the template.
	// +kubebuilder:validation:Optional
	ServiceMonitors []TemplateServiceMonitorsParameters `json:"serviceMonitors,omitempty" tf:"service_monitors,omitempty"`

	// The ID of the template, which is used for the outgoing reference.
	// +kubebuilder:validation:Optional
	TemplateID *string `json:"templateId,omitempty" tf:"template_id,omitempty"`

	// Last updated, for outgoing references.
	// +kubebuilder:validation:Optional
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`

	// Whether the system-supplied default template is used for outgoing references.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type TemplatePodMonitorsObservation struct {
}

type TemplatePodMonitorsParameters struct {

	// Config.
	// +kubebuilder:validation:Required
	Config *string `json:"config" tf:"config,omitempty"`

	// Name.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Used for the argument, if the configuration comes to the template, the template id.
	// +kubebuilder:validation:Optional
	TemplateID *string `json:"templateId,omitempty" tf:"template_id,omitempty"`
}

type TemplateRawJobsObservation struct {
}

type TemplateRawJobsParameters struct {

	// Config.
	// +kubebuilder:validation:Required
	Config *string `json:"config" tf:"config,omitempty"`

	// Name.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Used for the argument, if the configuration comes to the template, the template id.
	// +kubebuilder:validation:Optional
	TemplateID *string `json:"templateId,omitempty" tf:"template_id,omitempty"`
}

type TemplateServiceMonitorsObservation struct {
}

type TemplateServiceMonitorsParameters struct {

	// Config.
	// +kubebuilder:validation:Required
	Config *string `json:"config" tf:"config,omitempty"`

	// Name.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Used for the argument, if the configuration comes to the template, the template id.
	// +kubebuilder:validation:Optional
	TemplateID *string `json:"templateId,omitempty" tf:"template_id,omitempty"`
}

type TmpTkeTemplateObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TmpTkeTemplateParameters struct {

	// Template settings.
	// +kubebuilder:validation:Required
	Template []TemplateParameters `json:"template" tf:"template,omitempty"`
}

// TmpTkeTemplateSpec defines the desired state of TmpTkeTemplate
type TmpTkeTemplateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TmpTkeTemplateParameters `json:"forProvider"`
}

// TmpTkeTemplateStatus defines the observed state of TmpTkeTemplate.
type TmpTkeTemplateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TmpTkeTemplateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TmpTkeTemplate is the Schema for the TmpTkeTemplates API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type TmpTkeTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TmpTkeTemplateSpec   `json:"spec"`
	Status            TmpTkeTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TmpTkeTemplateList contains a list of TmpTkeTemplates
type TmpTkeTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TmpTkeTemplate `json:"items"`
}

// Repository type metadata.
var (
	TmpTkeTemplate_Kind             = "TmpTkeTemplate"
	TmpTkeTemplate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TmpTkeTemplate_Kind}.String()
	TmpTkeTemplate_KindAPIVersion   = TmpTkeTemplate_Kind + "." + CRDGroupVersion.String()
	TmpTkeTemplate_GroupVersionKind = CRDGroupVersion.WithKind(TmpTkeTemplate_Kind)
)

func init() {
	SchemeBuilder.Register(&TmpTkeTemplate{}, &TmpTkeTemplateList{})
}
