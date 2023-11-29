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

type NotificationAlertManagerObservation struct {
}

type NotificationAlertManagerParameters struct {

	// Cluster id.
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Cluster type.
	// +kubebuilder:validation:Optional
	ClusterType *string `json:"clusterType,omitempty" tf:"cluster_type,omitempty"`

	// Alert manager url.
	// +kubebuilder:validation:Required
	URL *string `json:"url" tf:"url,omitempty"`
}

type TmpTkeGlobalNotificationNotificationObservation struct {
}

type TmpTkeGlobalNotificationNotificationParameters struct {

	// Alert manager, if Type is `alertmanager`, this field is required.
	// +kubebuilder:validation:Optional
	AlertManager []NotificationAlertManagerParameters `json:"alertManager,omitempty" tf:"alert_manager,omitempty"`

	// Alarm notification switch.
	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// Alarm notification method, Valid values: `SMS`, `EMAIL`, `CALL`, `WECHAT`.
	// +kubebuilder:validation:Optional
	NotifyWay []*string `json:"notifyWay,omitempty" tf:"notify_way,omitempty"`

	// Phone Alarm Reach Notification, NotifyWay is `CALL`, and this parameter is used.
	// +kubebuilder:validation:Optional
	PhoneArriveNotice *bool `json:"phoneArriveNotice,omitempty" tf:"phone_arrive_notice,omitempty"`

	// Telephone alarm off-wheel interval, NotifyWay is `CALL`, and this parameter is used.
	// +kubebuilder:validation:Optional
	PhoneCircleInterval *float64 `json:"phoneCircleInterval,omitempty" tf:"phone_circle_interval,omitempty"`

	// Number of phone alerts (user group), NotifyWay is `CALL`, and this parameter is used.
	// +kubebuilder:validation:Optional
	PhoneCircleTimes *float64 `json:"phoneCircleTimes,omitempty" tf:"phone_circle_times,omitempty"`

	// Interval between telephone alarm rounds, NotifyWay is `CALL`, and this parameter is used.
	// +kubebuilder:validation:Optional
	PhoneInnerInterval *float64 `json:"phoneInnerInterval,omitempty" tf:"phone_inner_interval,omitempty"`

	// Phone alert sequence, NotifyWay is `CALL`, and this parameter is used.
	// +kubebuilder:validation:Optional
	PhoneNotifyOrder []*float64 `json:"phoneNotifyOrder,omitempty" tf:"phone_notify_order,omitempty"`

	// Alarm receiving group(user group).
	// +kubebuilder:validation:Optional
	ReceiverGroups []*string `json:"receiverGroups,omitempty" tf:"receiver_groups,omitempty"`

	// Convergence time.
	// +kubebuilder:validation:Optional
	RepeatInterval *string `json:"repeatInterval,omitempty" tf:"repeat_interval,omitempty"`

	// Effective end time.
	// +kubebuilder:validation:Optional
	TimeRangeEnd *string `json:"timeRangeEnd,omitempty" tf:"time_range_end,omitempty"`

	// Effective start time.
	// +kubebuilder:validation:Optional
	TimeRangeStart *string `json:"timeRangeStart,omitempty" tf:"time_range_start,omitempty"`

	// Alarm notification type, Valid values: `amp`, `webhook`, `alertmanager`.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// Web hook, if Type is `webhook`, this field is required.
	// +kubebuilder:validation:Optional
	WebHook *string `json:"webHook,omitempty" tf:"web_hook,omitempty"`
}

type TmpTkeGlobalNotificationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TmpTkeGlobalNotificationParameters struct {

	// Instance Id.
	// +crossplane:generate:reference:type=TmpInstance
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// Alarm notification channels.
	// +kubebuilder:validation:Required
	Notification []TmpTkeGlobalNotificationNotificationParameters `json:"notification" tf:"notification,omitempty"`
}

// TmpTkeGlobalNotificationSpec defines the desired state of TmpTkeGlobalNotification
type TmpTkeGlobalNotificationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TmpTkeGlobalNotificationParameters `json:"forProvider"`
}

// TmpTkeGlobalNotificationStatus defines the observed state of TmpTkeGlobalNotification.
type TmpTkeGlobalNotificationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TmpTkeGlobalNotificationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TmpTkeGlobalNotification is the Schema for the TmpTkeGlobalNotifications API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type TmpTkeGlobalNotification struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TmpTkeGlobalNotificationSpec   `json:"spec"`
	Status            TmpTkeGlobalNotificationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TmpTkeGlobalNotificationList contains a list of TmpTkeGlobalNotifications
type TmpTkeGlobalNotificationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TmpTkeGlobalNotification `json:"items"`
}

// Repository type metadata.
var (
	TmpTkeGlobalNotification_Kind             = "TmpTkeGlobalNotification"
	TmpTkeGlobalNotification_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TmpTkeGlobalNotification_Kind}.String()
	TmpTkeGlobalNotification_KindAPIVersion   = TmpTkeGlobalNotification_Kind + "." + CRDGroupVersion.String()
	TmpTkeGlobalNotification_GroupVersionKind = CRDGroupVersion.WithKind(TmpTkeGlobalNotification_Kind)
)

func init() {
	SchemeBuilder.Register(&TmpTkeGlobalNotification{}, &TmpTkeGlobalNotificationList{})
}