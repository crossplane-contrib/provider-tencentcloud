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

type ScalingPolicyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ScalingPolicyParameters struct {

	// Specifies whether the adjustment is an absolute number or a percentage of the current capacity. Valid values: `CHANGE_IN_CAPACITY`, `EXACT_CAPACITY` and `PERCENT_CHANGE_IN_CAPACITY`.
	// +kubebuilder:validation:Required
	AdjustmentType *string `json:"adjustmentType" tf:"adjustment_type,omitempty"`

	// Define the number of instances by which to scale.For `CHANGE_IN_CAPACITY` type or PERCENT_CHANGE_IN_CAPACITY, a positive increment adds to the current capacity and a negative value removes from the current capacity. For `EXACT_CAPACITY` type, it defines an absolute number of the existing Auto Scaling group size.
	// +kubebuilder:validation:Required
	AdjustmentValue *float64 `json:"adjustmentValue" tf:"adjustment_value,omitempty"`

	// Comparison operator. Valid values: `GREATER_THAN`, `GREATER_THAN_OR_EQUAL_TO`, `LESS_THAN`, `LESS_THAN_OR_EQUAL_TO`, `EQUAL_TO` and `NOT_EQUAL_TO`.
	// +kubebuilder:validation:Required
	ComparisonOperator *string `json:"comparisonOperator" tf:"comparison_operator,omitempty"`

	// Retry times. Valid value ranges: (1~10).
	// +kubebuilder:validation:Required
	ContinuousTime *float64 `json:"continuousTime" tf:"continuous_time,omitempty"`

	// Cooldwon time in second. Default is `30`0.
	// +kubebuilder:validation:Optional
	Cooldown *float64 `json:"cooldown,omitempty" tf:"cooldown,omitempty"`

	// Name of an indicator. Valid values: `CPU_UTILIZATION`, `MEM_UTILIZATION`, `LAN_TRAFFIC_OUT`, `LAN_TRAFFIC_IN`, `WAN_TRAFFIC_OUT` and `WAN_TRAFFIC_IN`.
	// +kubebuilder:validation:Required
	MetricName *string `json:"metricName" tf:"metric_name,omitempty"`

	// An ID group of users to be notified when an alarm is triggered.
	// +kubebuilder:validation:Optional
	NotificationUserGroupIds []*string `json:"notificationUserGroupIds,omitempty" tf:"notification_user_group_ids,omitempty"`

	// Time period in second. Valid values: `60` and `300`.
	// +kubebuilder:validation:Required
	Period *float64 `json:"period" tf:"period,omitempty"`

	// Name of a policy used to define a reaction when an alarm is triggered.
	// +kubebuilder:validation:Required
	PolicyName *string `json:"policyName" tf:"policy_name,omitempty"`

	// ID of a scaling group.
	// +crossplane:generate:reference:type=ScalingGroup
	// +kubebuilder:validation:Optional
	ScalingGroupID *string `json:"scalingGroupId,omitempty" tf:"scaling_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	ScalingGroupIDRef *v1.Reference `json:"scalingGroupIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ScalingGroupIDSelector *v1.Selector `json:"scalingGroupIdSelector,omitempty" tf:"-"`

	// Statistic types. Valid values: `AVERAGE`, `MAXIMUM` and `MINIMUM`. Default is `AVERAGE`.
	// +kubebuilder:validation:Optional
	Statistic *string `json:"statistic,omitempty" tf:"statistic,omitempty"`

	// Alarm threshold.
	// +kubebuilder:validation:Required
	Threshold *float64 `json:"threshold" tf:"threshold,omitempty"`
}

// ScalingPolicySpec defines the desired state of ScalingPolicy
type ScalingPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ScalingPolicyParameters `json:"forProvider"`
}

// ScalingPolicyStatus defines the observed state of ScalingPolicy.
type ScalingPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ScalingPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ScalingPolicy is the Schema for the ScalingPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type ScalingPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ScalingPolicySpec   `json:"spec"`
	Status            ScalingPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ScalingPolicyList contains a list of ScalingPolicys
type ScalingPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ScalingPolicy `json:"items"`
}

// Repository type metadata.
var (
	ScalingPolicy_Kind             = "ScalingPolicy"
	ScalingPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ScalingPolicy_Kind}.String()
	ScalingPolicy_KindAPIVersion   = ScalingPolicy_Kind + "." + CRDGroupVersion.String()
	ScalingPolicy_GroupVersionKind = CRDGroupVersion.WithKind(ScalingPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&ScalingPolicy{}, &ScalingPolicyList{})
}
