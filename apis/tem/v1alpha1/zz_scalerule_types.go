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

type AutoscalerObservation struct {
}

type AutoscalerParameters struct {

	// name.
	// +kubebuilder:validation:Required
	AutoscalerName *string `json:"autoscalerName" tf:"autoscaler_name,omitempty"`

	// scaler based on cron configuration.
	// +kubebuilder:validation:Optional
	CronHorizontalAutoscaler []CronHorizontalAutoscalerParameters `json:"cronHorizontalAutoscaler,omitempty" tf:"cron_horizontal_autoscaler,omitempty"`

	// description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// enable AutoScaler.
	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// scaler based on metrics.
	// +kubebuilder:validation:Optional
	HorizontalAutoscaler []HorizontalAutoscalerParameters `json:"horizontalAutoscaler,omitempty" tf:"horizontal_autoscaler,omitempty"`

	// maximal replica number.
	// +kubebuilder:validation:Required
	MaxReplicas *float64 `json:"maxReplicas" tf:"max_replicas,omitempty"`

	// minimal replica number.
	// +kubebuilder:validation:Required
	MinReplicas *float64 `json:"minReplicas" tf:"min_replicas,omitempty"`
}

type CronHorizontalAutoscalerObservation struct {
}

type CronHorizontalAutoscalerParameters struct {

	// enable scaler.
	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// name.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// period.
	// +kubebuilder:validation:Required
	Period *string `json:"period" tf:"period,omitempty"`

	// priority.
	// +kubebuilder:validation:Required
	Priority *float64 `json:"priority" tf:"priority,omitempty"`

	// schedule payload.
	// +kubebuilder:validation:Required
	Schedules []SchedulesParameters `json:"schedules" tf:"schedules,omitempty"`
}

type HorizontalAutoscalerObservation struct {
}

type HorizontalAutoscalerParameters struct {

	// enable scaler.
	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// maximal replica number.
	// +kubebuilder:validation:Required
	MaxReplicas *float64 `json:"maxReplicas" tf:"max_replicas,omitempty"`

	// metric name.
	// +kubebuilder:validation:Required
	Metrics *string `json:"metrics" tf:"metrics,omitempty"`

	// minimal replica number.
	// +kubebuilder:validation:Required
	MinReplicas *float64 `json:"minReplicas" tf:"min_replicas,omitempty"`

	// metric threshold.
	// +kubebuilder:validation:Required
	Threshold *float64 `json:"threshold" tf:"threshold,omitempty"`
}

type ScaleRuleObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ScaleRuleParameters struct {

	// application ID.
	// +crossplane:generate:reference:type=Application
	// +kubebuilder:validation:Optional
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// +kubebuilder:validation:Optional
	ApplicationIDRef *v1.Reference `json:"applicationIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ApplicationIDSelector *v1.Selector `json:"applicationIdSelector,omitempty" tf:"-"`

	// .
	// +kubebuilder:validation:Required
	Autoscaler []AutoscalerParameters `json:"autoscaler" tf:"autoscaler,omitempty"`

	// environment ID.
	// +crossplane:generate:reference:type=Environment
	// +kubebuilder:validation:Optional
	EnvironmentID *string `json:"environmentId,omitempty" tf:"environment_id,omitempty"`

	// +kubebuilder:validation:Optional
	EnvironmentIDRef *v1.Reference `json:"environmentIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	EnvironmentIDSelector *v1.Selector `json:"environmentIdSelector,omitempty" tf:"-"`

	// application ID, which is combined by environment ID and application ID, like `en-o5edaepv#app-3j29aa2p`.
	// +crossplane:generate:reference:type=Workload
	// +kubebuilder:validation:Optional
	WorkloadID *string `json:"workloadId,omitempty" tf:"workload_id,omitempty"`

	// +kubebuilder:validation:Optional
	WorkloadIDRef *v1.Reference `json:"workloadIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	WorkloadIDSelector *v1.Selector `json:"workloadIdSelector,omitempty" tf:"-"`
}

type SchedulesObservation struct {
}

type SchedulesParameters struct {

	// start time.
	// +kubebuilder:validation:Required
	StartAt *string `json:"startAt" tf:"start_at,omitempty"`

	// target replica number.
	// +kubebuilder:validation:Required
	TargetReplicas *float64 `json:"targetReplicas" tf:"target_replicas,omitempty"`
}

// ScaleRuleSpec defines the desired state of ScaleRule
type ScaleRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ScaleRuleParameters `json:"forProvider"`
}

// ScaleRuleStatus defines the observed state of ScaleRule.
type ScaleRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ScaleRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ScaleRule is the Schema for the ScaleRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type ScaleRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ScaleRuleSpec   `json:"spec"`
	Status            ScaleRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ScaleRuleList contains a list of ScaleRules
type ScaleRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ScaleRule `json:"items"`
}

// Repository type metadata.
var (
	ScaleRule_Kind             = "ScaleRule"
	ScaleRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ScaleRule_Kind}.String()
	ScaleRule_KindAPIVersion   = ScaleRule_Kind + "." + CRDGroupVersion.String()
	ScaleRule_GroupVersionKind = CRDGroupVersion.WithKind(ScaleRule_Kind)
)

func init() {
	SchemeBuilder.Register(&ScaleRule{}, &ScaleRuleList{})
}
