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

type EmailAddressObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type EmailAddressParameters struct {

	// Your sender address. (You can create up to 10 sender addresses for each domain.).
	// +kubebuilder:validation:Required
	EmailAddress *string `json:"emailAddress" tf:"email_address,omitempty"`

	// Sender name.
	// +kubebuilder:validation:Optional
	EmailSenderName *string `json:"emailSenderName,omitempty" tf:"email_sender_name,omitempty"`
}

// EmailAddressSpec defines the desired state of EmailAddress
type EmailAddressSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EmailAddressParameters `json:"forProvider"`
}

// EmailAddressStatus defines the observed state of EmailAddress.
type EmailAddressStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EmailAddressObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EmailAddress is the Schema for the EmailAddresss API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type EmailAddress struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EmailAddressSpec   `json:"spec"`
	Status            EmailAddressStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EmailAddressList contains a list of EmailAddresss
type EmailAddressList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EmailAddress `json:"items"`
}

// Repository type metadata.
var (
	EmailAddress_Kind             = "EmailAddress"
	EmailAddress_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EmailAddress_Kind}.String()
	EmailAddress_KindAPIVersion   = EmailAddress_Kind + "." + CRDGroupVersion.String()
	EmailAddress_GroupVersionKind = CRDGroupVersion.WithKind(EmailAddress_Kind)
)

func init() {
	SchemeBuilder.Register(&EmailAddress{}, &EmailAddressList{})
}
