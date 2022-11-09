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

type VPVAclAttachmentObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type VPVAclAttachmentParameters struct {

	// ID of the attached ACL.
	// +crossplane:generate:reference:type=VPCAcl
	// +kubebuilder:validation:Optional
	ACLID *string `json:"aclId,omitempty" tf:"acl_id,omitempty"`

	// +kubebuilder:validation:Optional
	ACLIDRef *v1.Reference `json:"aclidRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ACLIDSelector *v1.Selector `json:"aclidSelector,omitempty" tf:"-"`

	// The Subnet instance ID.
	// +crossplane:generate:reference:type=Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

// VPVAclAttachmentSpec defines the desired state of VPVAclAttachment
type VPVAclAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPVAclAttachmentParameters `json:"forProvider"`
}

// VPVAclAttachmentStatus defines the observed state of VPVAclAttachment.
type VPVAclAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPVAclAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VPVAclAttachment is the Schema for the VPVAclAttachments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type VPVAclAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPVAclAttachmentSpec   `json:"spec"`
	Status            VPVAclAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPVAclAttachmentList contains a list of VPVAclAttachments
type VPVAclAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPVAclAttachment `json:"items"`
}

// Repository type metadata.
var (
	VPVAclAttachment_Kind             = "VPVAclAttachment"
	VPVAclAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPVAclAttachment_Kind}.String()
	VPVAclAttachment_KindAPIVersion   = VPVAclAttachment_Kind + "." + CRDGroupVersion.String()
	VPVAclAttachment_GroupVersionKind = CRDGroupVersion.WithKind(VPVAclAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&VPVAclAttachment{}, &VPVAclAttachmentList{})
}