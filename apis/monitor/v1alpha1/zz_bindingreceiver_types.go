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

type BindingReceiverObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type BindingReceiverParameters struct {

	// Policy group ID for binding receivers.
	// +kubebuilder:validation:Required
	GroupID *float64 `json:"groupId" tf:"group_id,omitempty"`

	// A list of receivers(will overwrite the configuration of the server or other resources). Each element contains the following attributes:
	// +kubebuilder:validation:Optional
	Receivers []ReceiversParameters `json:"receivers,omitempty" tf:"receivers,omitempty"`
}

type ReceiversObservation struct {
}

type ReceiversParameters struct {

	// End of alarm period. Meaning with `start_time`.
	// +kubebuilder:validation:Optional
	EndTime *float64 `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// Method of warning notification.Optional `CALL`,`EMAIL`,`SITE`,`SMS`,`WECHAT`.
	// +kubebuilder:validation:Required
	NotifyWay []*string `json:"notifyWay" tf:"notify_way,omitempty"`

	// Alert sending language. Optional `en-US`,`zh-CN`.
	// +kubebuilder:validation:Optional
	ReceiveLanguage *string `json:"receiveLanguage,omitempty" tf:"receive_language,omitempty"`

	// Alarm receive group ID list.
	// +kubebuilder:validation:Optional
	ReceiverGroupList []*float64 `json:"receiverGroupList,omitempty" tf:"receiver_group_list,omitempty"`

	// Receive type. Optional `group`,`user`.
	// +kubebuilder:validation:Required
	ReceiverType *string `json:"receiverType" tf:"receiver_type,omitempty"`

	// Alarm receiver ID list.
	// +kubebuilder:validation:Optional
	ReceiverUserList []*float64 `json:"receiverUserList,omitempty" tf:"receiver_user_list,omitempty"`

	// Alarm period start time. Valid value ranges: (0~86399). which removes the date after it is converted to Beijing time as a Unix timestamp, for example 7200 means '10:0:0'.
	// +kubebuilder:validation:Optional
	StartTime *float64 `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

// BindingReceiverSpec defines the desired state of BindingReceiver
type BindingReceiverSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BindingReceiverParameters `json:"forProvider"`
}

// BindingReceiverStatus defines the observed state of BindingReceiver.
type BindingReceiverStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BindingReceiverObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BindingReceiver is the Schema for the BindingReceivers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type BindingReceiver struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BindingReceiverSpec   `json:"spec"`
	Status            BindingReceiverStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BindingReceiverList contains a list of BindingReceivers
type BindingReceiverList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BindingReceiver `json:"items"`
}

// Repository type metadata.
var (
	BindingReceiver_Kind             = "BindingReceiver"
	BindingReceiver_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BindingReceiver_Kind}.String()
	BindingReceiver_KindAPIVersion   = BindingReceiver_Kind + "." + CRDGroupVersion.String()
	BindingReceiver_GroupVersionKind = CRDGroupVersion.WithKind(BindingReceiver_Kind)
)

func init() {
	SchemeBuilder.Register(&BindingReceiver{}, &BindingReceiverList{})
}
