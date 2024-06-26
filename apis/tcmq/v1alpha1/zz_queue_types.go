/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type QueueInitParameters struct {

	// Dead letter queue name.
	// Dead letter queue name.
	DeadLetterQueueName *string `json:"deadLetterQueueName,omitempty" tf:"dead_letter_queue_name,omitempty"`

	// First lookback interval.
	// First lookback interval.
	FirstQueryInterval *float64 `json:"firstQueryInterval,omitempty" tf:"first_query_interval,omitempty"`

	// Maximum number of heaped messages. The value range is 1,000,000-10,000,000 during the beta test and can be 1,000,000-1,000,000,000 after the product is officially released. The default value is 10,000,000 during the beta test and will be 100,000,000 after the product is officially released.
	// Maximum number of heaped messages. The value range is 1,000,000-10,000,000 during the beta test and can be 1,000,000-1,000,000,000 after the product is officially released. The default value is 10,000,000 during the beta test and will be 100,000,000 after the product is officially released.
	MaxMsgHeapNum *float64 `json:"maxMsgHeapNum,omitempty" tf:"max_msg_heap_num,omitempty"`

	// Maximum message length. Value range: 1024-65536 bytes (i.e., 1-64 KB). Default value: 65536.
	// Maximum message length. Value range: 1024-65536 bytes (i.e., 1-64 KB). Default value: 65536.
	MaxMsgSize *float64 `json:"maxMsgSize,omitempty" tf:"max_msg_size,omitempty"`

	// Maximum number of lookbacks.
	// Maximum number of lookbacks.
	MaxQueryCount *float64 `json:"maxQueryCount,omitempty" tf:"max_query_count,omitempty"`

	// Maximum receipt times. Value range: 1-1000.
	// Maximum receipt times. Value range: 1-1000.
	MaxReceiveCount *float64 `json:"maxReceiveCount,omitempty" tf:"max_receive_count,omitempty"`

	// Maximum period in seconds before an unconsumed message expires, which is required if policy is 1. Value range: 300-43200. This value should be smaller than msgRetentionSeconds (maximum message retention period).
	// Maximum period in seconds before an unconsumed message expires, which is required if `policy` is 1. Value range: 300-43200. This value should be smaller than `msgRetentionSeconds` (maximum message retention period).
	MaxTimeToLive *float64 `json:"maxTimeToLive,omitempty" tf:"max_time_to_live,omitempty"`

	// The max period during which a message is retained before it is automatically acknowledged. Value range: 30-43,200 seconds (30 seconds to 12 hours). Default value: 3600 seconds (1 hour).
	// The max period during which a message is retained before it is automatically acknowledged. Value range: 30-43,200 seconds (30 seconds to 12 hours). Default value: 3600 seconds (1 hour).
	MsgRetentionSeconds *float64 `json:"msgRetentionSeconds,omitempty" tf:"msg_retention_seconds,omitempty"`

	// Dead letter policy. 0: message has been consumed multiple times but not deleted; 1: Time-To-Live has elapsed.
	// Dead letter policy. 0: message has been consumed multiple times but not deleted; 1: `Time-To-Live` has elapsed.
	Policy *float64 `json:"policy,omitempty" tf:"policy,omitempty"`

	// Long polling wait time for message reception. Value range: 0-30 seconds. Default value: 0.
	// Long polling wait time for message reception. Value range: 0-30 seconds. Default value: 0.
	PollingWaitSeconds *float64 `json:"pollingWaitSeconds,omitempty" tf:"polling_wait_seconds,omitempty"`

	// Queue name, which must be unique under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	// Queue name, which must be unique under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	QueueName *string `json:"queueName,omitempty" tf:"queue_name,omitempty"`

	// Queue storage space configured for message rewind. Value range: 10,240-512,000 MB (if message rewind is enabled). The value 0 indicates that message rewind is not enabled.
	// Queue storage space configured for message rewind. Value range: 10,240-512,000 MB (if message rewind is enabled). The value `0` indicates that message rewind is not enabled.
	RetentionSizeInMb *float64 `json:"retentionSizeInMb,omitempty" tf:"retention_size_in_mb,omitempty"`

	// Rewindable time of messages in the queue. Value range: 0-1,296,000s (if message rewind is enabled). The value 0 indicates that message rewind is not enabled.
	// Rewindable time of messages in the queue. Value range: 0-1,296,000s (if message rewind is enabled). The value `0` indicates that message rewind is not enabled.
	RewindSeconds *float64 `json:"rewindSeconds,omitempty" tf:"rewind_seconds,omitempty"`

	// Whether to enable message trace. true: yes; false: no. If this field is not configured, the feature will not be enabled.
	// Whether to enable message trace. true: yes; false: no. If this field is not configured, the feature will not be enabled.
	Trace *bool `json:"trace,omitempty" tf:"trace,omitempty"`

	// 1: transaction queue; 0: general queue.
	// 1: transaction queue; 0: general queue.
	Transaction *float64 `json:"transaction,omitempty" tf:"transaction,omitempty"`

	// Message visibility timeout period. Value range: 1-43200 seconds (i.e., 12 hours). Default value: 30.
	// Message visibility timeout period. Value range: 1-43200 seconds (i.e., 12 hours). Default value: 30.
	VisibilityTimeout *float64 `json:"visibilityTimeout,omitempty" tf:"visibility_timeout,omitempty"`
}

type QueueObservation struct {

	// Dead letter queue name.
	// Dead letter queue name.
	DeadLetterQueueName *string `json:"deadLetterQueueName,omitempty" tf:"dead_letter_queue_name,omitempty"`

	// First lookback interval.
	// First lookback interval.
	FirstQueryInterval *float64 `json:"firstQueryInterval,omitempty" tf:"first_query_interval,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Maximum number of heaped messages. The value range is 1,000,000-10,000,000 during the beta test and can be 1,000,000-1,000,000,000 after the product is officially released. The default value is 10,000,000 during the beta test and will be 100,000,000 after the product is officially released.
	// Maximum number of heaped messages. The value range is 1,000,000-10,000,000 during the beta test and can be 1,000,000-1,000,000,000 after the product is officially released. The default value is 10,000,000 during the beta test and will be 100,000,000 after the product is officially released.
	MaxMsgHeapNum *float64 `json:"maxMsgHeapNum,omitempty" tf:"max_msg_heap_num,omitempty"`

	// Maximum message length. Value range: 1024-65536 bytes (i.e., 1-64 KB). Default value: 65536.
	// Maximum message length. Value range: 1024-65536 bytes (i.e., 1-64 KB). Default value: 65536.
	MaxMsgSize *float64 `json:"maxMsgSize,omitempty" tf:"max_msg_size,omitempty"`

	// Maximum number of lookbacks.
	// Maximum number of lookbacks.
	MaxQueryCount *float64 `json:"maxQueryCount,omitempty" tf:"max_query_count,omitempty"`

	// Maximum receipt times. Value range: 1-1000.
	// Maximum receipt times. Value range: 1-1000.
	MaxReceiveCount *float64 `json:"maxReceiveCount,omitempty" tf:"max_receive_count,omitempty"`

	// Maximum period in seconds before an unconsumed message expires, which is required if policy is 1. Value range: 300-43200. This value should be smaller than msgRetentionSeconds (maximum message retention period).
	// Maximum period in seconds before an unconsumed message expires, which is required if `policy` is 1. Value range: 300-43200. This value should be smaller than `msgRetentionSeconds` (maximum message retention period).
	MaxTimeToLive *float64 `json:"maxTimeToLive,omitempty" tf:"max_time_to_live,omitempty"`

	// The max period during which a message is retained before it is automatically acknowledged. Value range: 30-43,200 seconds (30 seconds to 12 hours). Default value: 3600 seconds (1 hour).
	// The max period during which a message is retained before it is automatically acknowledged. Value range: 30-43,200 seconds (30 seconds to 12 hours). Default value: 3600 seconds (1 hour).
	MsgRetentionSeconds *float64 `json:"msgRetentionSeconds,omitempty" tf:"msg_retention_seconds,omitempty"`

	// Dead letter policy. 0: message has been consumed multiple times but not deleted; 1: Time-To-Live has elapsed.
	// Dead letter policy. 0: message has been consumed multiple times but not deleted; 1: `Time-To-Live` has elapsed.
	Policy *float64 `json:"policy,omitempty" tf:"policy,omitempty"`

	// Long polling wait time for message reception. Value range: 0-30 seconds. Default value: 0.
	// Long polling wait time for message reception. Value range: 0-30 seconds. Default value: 0.
	PollingWaitSeconds *float64 `json:"pollingWaitSeconds,omitempty" tf:"polling_wait_seconds,omitempty"`

	// Queue name, which must be unique under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	// Queue name, which must be unique under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	QueueName *string `json:"queueName,omitempty" tf:"queue_name,omitempty"`

	// Queue storage space configured for message rewind. Value range: 10,240-512,000 MB (if message rewind is enabled). The value 0 indicates that message rewind is not enabled.
	// Queue storage space configured for message rewind. Value range: 10,240-512,000 MB (if message rewind is enabled). The value `0` indicates that message rewind is not enabled.
	RetentionSizeInMb *float64 `json:"retentionSizeInMb,omitempty" tf:"retention_size_in_mb,omitempty"`

	// Rewindable time of messages in the queue. Value range: 0-1,296,000s (if message rewind is enabled). The value 0 indicates that message rewind is not enabled.
	// Rewindable time of messages in the queue. Value range: 0-1,296,000s (if message rewind is enabled). The value `0` indicates that message rewind is not enabled.
	RewindSeconds *float64 `json:"rewindSeconds,omitempty" tf:"rewind_seconds,omitempty"`

	// Whether to enable message trace. true: yes; false: no. If this field is not configured, the feature will not be enabled.
	// Whether to enable message trace. true: yes; false: no. If this field is not configured, the feature will not be enabled.
	Trace *bool `json:"trace,omitempty" tf:"trace,omitempty"`

	// 1: transaction queue; 0: general queue.
	// 1: transaction queue; 0: general queue.
	Transaction *float64 `json:"transaction,omitempty" tf:"transaction,omitempty"`

	// Message visibility timeout period. Value range: 1-43200 seconds (i.e., 12 hours). Default value: 30.
	// Message visibility timeout period. Value range: 1-43200 seconds (i.e., 12 hours). Default value: 30.
	VisibilityTimeout *float64 `json:"visibilityTimeout,omitempty" tf:"visibility_timeout,omitempty"`
}

type QueueParameters struct {

	// Dead letter queue name.
	// Dead letter queue name.
	// +kubebuilder:validation:Optional
	DeadLetterQueueName *string `json:"deadLetterQueueName,omitempty" tf:"dead_letter_queue_name,omitempty"`

	// First lookback interval.
	// First lookback interval.
	// +kubebuilder:validation:Optional
	FirstQueryInterval *float64 `json:"firstQueryInterval,omitempty" tf:"first_query_interval,omitempty"`

	// Maximum number of heaped messages. The value range is 1,000,000-10,000,000 during the beta test and can be 1,000,000-1,000,000,000 after the product is officially released. The default value is 10,000,000 during the beta test and will be 100,000,000 after the product is officially released.
	// Maximum number of heaped messages. The value range is 1,000,000-10,000,000 during the beta test and can be 1,000,000-1,000,000,000 after the product is officially released. The default value is 10,000,000 during the beta test and will be 100,000,000 after the product is officially released.
	// +kubebuilder:validation:Optional
	MaxMsgHeapNum *float64 `json:"maxMsgHeapNum,omitempty" tf:"max_msg_heap_num,omitempty"`

	// Maximum message length. Value range: 1024-65536 bytes (i.e., 1-64 KB). Default value: 65536.
	// Maximum message length. Value range: 1024-65536 bytes (i.e., 1-64 KB). Default value: 65536.
	// +kubebuilder:validation:Optional
	MaxMsgSize *float64 `json:"maxMsgSize,omitempty" tf:"max_msg_size,omitempty"`

	// Maximum number of lookbacks.
	// Maximum number of lookbacks.
	// +kubebuilder:validation:Optional
	MaxQueryCount *float64 `json:"maxQueryCount,omitempty" tf:"max_query_count,omitempty"`

	// Maximum receipt times. Value range: 1-1000.
	// Maximum receipt times. Value range: 1-1000.
	// +kubebuilder:validation:Optional
	MaxReceiveCount *float64 `json:"maxReceiveCount,omitempty" tf:"max_receive_count,omitempty"`

	// Maximum period in seconds before an unconsumed message expires, which is required if policy is 1. Value range: 300-43200. This value should be smaller than msgRetentionSeconds (maximum message retention period).
	// Maximum period in seconds before an unconsumed message expires, which is required if `policy` is 1. Value range: 300-43200. This value should be smaller than `msgRetentionSeconds` (maximum message retention period).
	// +kubebuilder:validation:Optional
	MaxTimeToLive *float64 `json:"maxTimeToLive,omitempty" tf:"max_time_to_live,omitempty"`

	// The max period during which a message is retained before it is automatically acknowledged. Value range: 30-43,200 seconds (30 seconds to 12 hours). Default value: 3600 seconds (1 hour).
	// The max period during which a message is retained before it is automatically acknowledged. Value range: 30-43,200 seconds (30 seconds to 12 hours). Default value: 3600 seconds (1 hour).
	// +kubebuilder:validation:Optional
	MsgRetentionSeconds *float64 `json:"msgRetentionSeconds,omitempty" tf:"msg_retention_seconds,omitempty"`

	// Dead letter policy. 0: message has been consumed multiple times but not deleted; 1: Time-To-Live has elapsed.
	// Dead letter policy. 0: message has been consumed multiple times but not deleted; 1: `Time-To-Live` has elapsed.
	// +kubebuilder:validation:Optional
	Policy *float64 `json:"policy,omitempty" tf:"policy,omitempty"`

	// Long polling wait time for message reception. Value range: 0-30 seconds. Default value: 0.
	// Long polling wait time for message reception. Value range: 0-30 seconds. Default value: 0.
	// +kubebuilder:validation:Optional
	PollingWaitSeconds *float64 `json:"pollingWaitSeconds,omitempty" tf:"polling_wait_seconds,omitempty"`

	// Queue name, which must be unique under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	// Queue name, which must be unique under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	// +kubebuilder:validation:Optional
	QueueName *string `json:"queueName,omitempty" tf:"queue_name,omitempty"`

	// Queue storage space configured for message rewind. Value range: 10,240-512,000 MB (if message rewind is enabled). The value 0 indicates that message rewind is not enabled.
	// Queue storage space configured for message rewind. Value range: 10,240-512,000 MB (if message rewind is enabled). The value `0` indicates that message rewind is not enabled.
	// +kubebuilder:validation:Optional
	RetentionSizeInMb *float64 `json:"retentionSizeInMb,omitempty" tf:"retention_size_in_mb,omitempty"`

	// Rewindable time of messages in the queue. Value range: 0-1,296,000s (if message rewind is enabled). The value 0 indicates that message rewind is not enabled.
	// Rewindable time of messages in the queue. Value range: 0-1,296,000s (if message rewind is enabled). The value `0` indicates that message rewind is not enabled.
	// +kubebuilder:validation:Optional
	RewindSeconds *float64 `json:"rewindSeconds,omitempty" tf:"rewind_seconds,omitempty"`

	// Whether to enable message trace. true: yes; false: no. If this field is not configured, the feature will not be enabled.
	// Whether to enable message trace. true: yes; false: no. If this field is not configured, the feature will not be enabled.
	// +kubebuilder:validation:Optional
	Trace *bool `json:"trace,omitempty" tf:"trace,omitempty"`

	// 1: transaction queue; 0: general queue.
	// 1: transaction queue; 0: general queue.
	// +kubebuilder:validation:Optional
	Transaction *float64 `json:"transaction,omitempty" tf:"transaction,omitempty"`

	// Message visibility timeout period. Value range: 1-43200 seconds (i.e., 12 hours). Default value: 30.
	// Message visibility timeout period. Value range: 1-43200 seconds (i.e., 12 hours). Default value: 30.
	// +kubebuilder:validation:Optional
	VisibilityTimeout *float64 `json:"visibilityTimeout,omitempty" tf:"visibility_timeout,omitempty"`
}

// QueueSpec defines the desired state of Queue
type QueueSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     QueueParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider QueueInitParameters `json:"initProvider,omitempty"`
}

// QueueStatus defines the observed state of Queue.
type QueueStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        QueueObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Queue is the Schema for the Queues API. Provides a resource to create a tcmq queue
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type Queue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.queueName) || (has(self.initProvider) && has(self.initProvider.queueName))",message="spec.forProvider.queueName is a required parameter"
	Spec   QueueSpec   `json:"spec"`
	Status QueueStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// QueueList contains a list of Queues
type QueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Queue `json:"items"`
}

// Repository type metadata.
var (
	Queue_Kind             = "Queue"
	Queue_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Queue_Kind}.String()
	Queue_KindAPIVersion   = Queue_Kind + "." + CRDGroupVersion.String()
	Queue_GroupVersionKind = CRDGroupVersion.WithKind(Queue_Kind)
)

func init() {
	SchemeBuilder.Register(&Queue{}, &QueueList{})
}
