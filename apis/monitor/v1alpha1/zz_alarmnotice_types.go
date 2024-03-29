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

type AlarmNoticeObservation struct {
	AmpConsumerID *string `json:"ampConsumerId,omitempty" tf:"amp_consumer_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IsPreset *float64 `json:"isPreset,omitempty" tf:"is_preset,omitempty"`

	PolicyIds []*string `json:"policyIds,omitempty" tf:"policy_ids,omitempty"`

	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`

	UpdatedBy *string `json:"updatedBy,omitempty" tf:"updated_by,omitempty"`
}

type AlarmNoticeParameters struct {

	// A maximum of one alarm notification can be pushed to the CLS service.
	// +kubebuilder:validation:Optional
	ClsNotices []ClsNoticesParameters `json:"clsNotices,omitempty" tf:"cls_notices,omitempty"`

	// Notification template name within 60.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Notification language zh-CN=Chinese en-US=English.
	// +kubebuilder:validation:Required
	NoticeLanguage *string `json:"noticeLanguage" tf:"notice_language,omitempty"`

	// Alarm notification type ALARM=Notification not restored OK=Notification restored ALL.
	// +kubebuilder:validation:Required
	NoticeType *string `json:"noticeType" tf:"notice_type,omitempty"`

	// The maximum number of callback notifications is 3.
	// +kubebuilder:validation:Optional
	URLNotices []URLNoticesParameters `json:"urlNotices,omitempty" tf:"url_notices,omitempty"`

	// Alarm notification template list.(At most five).
	// +kubebuilder:validation:Optional
	UserNotices []UserNoticesParameters `json:"userNotices,omitempty" tf:"user_notices,omitempty"`
}

type ClsNoticesObservation struct {
}

type ClsNoticesParameters struct {

	// Start-stop status, can not be transmitted, default enabled. 0= Disabled, 1= enabled.
	// +kubebuilder:validation:Optional
	Enable *float64 `json:"enable,omitempty" tf:"enable,omitempty"`

	// Log collection Id.
	// +kubebuilder:validation:Required
	LogSetID *string `json:"logSetId" tf:"log_set_id,omitempty"`

	// Regional.
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// Theme Id.
	// +kubebuilder:validation:Required
	TopicID *string `json:"topicId" tf:"topic_id,omitempty"`
}

type URLNoticesObservation struct {
}

type URLNoticesParameters struct {

	// Notification End Time Seconds at the start of a day.
	// +kubebuilder:validation:Optional
	EndTime *float64 `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// If passed verification `0` is no, `1` is yes. Default `0`.
	// +kubebuilder:validation:Optional
	IsValid *float64 `json:"isValid,omitempty" tf:"is_valid,omitempty"`

	// Notification Start Time Number of seconds at the start of a day.
	// +kubebuilder:validation:Optional
	StartTime *float64 `json:"startTime,omitempty" tf:"start_time,omitempty"`

	// Callback URL (limited to 256 characters).
	// +kubebuilder:validation:Required
	URL *string `json:"url" tf:"url,omitempty"`

	// Verification code.
	// +kubebuilder:validation:Optional
	ValidationCode *string `json:"validationCode,omitempty" tf:"validation_code,omitempty"`

	// Notification period 1-7 indicates Monday to Sunday.
	// +kubebuilder:validation:Optional
	Weekday []*float64 `json:"weekday,omitempty" tf:"weekday,omitempty"`
}

type UserNoticesObservation struct {
}

type UserNoticesParameters struct {

	// The number of seconds since the notification end time 00:00:00 (value range 0-86399).
	// +kubebuilder:validation:Required
	EndTime *float64 `json:"endTime" tf:"end_time,omitempty"`

	// User group ID list.
	// +kubebuilder:validation:Optional
	GroupIds []*float64 `json:"groupIds,omitempty" tf:"group_ids,omitempty"`

	// Contact notification required 0= No 1= Yes.
	// +kubebuilder:validation:Optional
	NeedPhoneArriveNotice *float64 `json:"needPhoneArriveNotice,omitempty" tf:"need_phone_arrive_notice,omitempty"`

	// Notification Channel List EMAIL=Mail SMS=SMS CALL=Telephone WECHAT=WeChat RTX=Enterprise WeChat.
	// +kubebuilder:validation:Required
	NoticeWay []*string `json:"noticeWay" tf:"notice_way,omitempty"`

	// Call type SYNC= Simultaneous call CIRCLE= Round call If this parameter is not specified, the default value is round call.
	// +kubebuilder:validation:Optional
	PhoneCallType *string `json:"phoneCallType,omitempty" tf:"phone_call_type,omitempty"`

	// Number of seconds between polls (value range: 60-900).
	// +kubebuilder:validation:Optional
	PhoneCircleInterval *float64 `json:"phoneCircleInterval,omitempty" tf:"phone_circle_interval,omitempty"`

	// Number of telephone polls (value range: 1-5).
	// +kubebuilder:validation:Optional
	PhoneCircleTimes *float64 `json:"phoneCircleTimes,omitempty" tf:"phone_circle_times,omitempty"`

	// Number of seconds between calls in a polling session (value range: 60-900).
	// +kubebuilder:validation:Optional
	PhoneInnerInterval *float64 `json:"phoneInnerInterval,omitempty" tf:"phone_inner_interval,omitempty"`

	// Telephone polling list.
	// +kubebuilder:validation:Optional
	PhoneOrder []*float64 `json:"phoneOrder,omitempty" tf:"phone_order,omitempty"`

	// Recipient Type USER=User GROUP=User Group.
	// +kubebuilder:validation:Required
	ReceiverType *string `json:"receiverType" tf:"receiver_type,omitempty"`

	// The number of seconds since the notification start time 00:00:00 (value range 0-86399).
	// +kubebuilder:validation:Required
	StartTime *float64 `json:"startTime" tf:"start_time,omitempty"`

	// User UID List.
	// +kubebuilder:validation:Optional
	UserIds []*float64 `json:"userIds,omitempty" tf:"user_ids,omitempty"`

	// Notification period 1-7 indicates Monday to Sunday.
	// +kubebuilder:validation:Optional
	Weekday []*float64 `json:"weekday,omitempty" tf:"weekday,omitempty"`
}

// AlarmNoticeSpec defines the desired state of AlarmNotice
type AlarmNoticeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AlarmNoticeParameters `json:"forProvider"`
}

// AlarmNoticeStatus defines the observed state of AlarmNotice.
type AlarmNoticeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AlarmNoticeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AlarmNotice is the Schema for the AlarmNotices API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type AlarmNotice struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AlarmNoticeSpec   `json:"spec"`
	Status            AlarmNoticeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AlarmNoticeList contains a list of AlarmNotices
type AlarmNoticeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AlarmNotice `json:"items"`
}

// Repository type metadata.
var (
	AlarmNotice_Kind             = "AlarmNotice"
	AlarmNotice_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AlarmNotice_Kind}.String()
	AlarmNotice_KindAPIVersion   = AlarmNotice_Kind + "." + CRDGroupVersion.String()
	AlarmNotice_GroupVersionKind = CRDGroupVersion.WithKind(AlarmNotice_Kind)
)

func init() {
	SchemeBuilder.Register(&AlarmNotice{}, &AlarmNoticeList{})
}
