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

type BucketDomainCertificateAttachmentObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type BucketDomainCertificateAttachmentParameters struct {

	// Bucket name.
	// +crossplane:generate:reference:type=Bucket
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// The certificate of specified doamin.
	// +kubebuilder:validation:Required
	DomainCertificate []DomainCertificateParameters `json:"domainCertificate" tf:"domain_certificate,omitempty"`
}

type CertificateObservation struct {
}

type CertificateParameters struct {

	// Certificate type.
	// +kubebuilder:validation:Required
	CertType *string `json:"certType" tf:"cert_type,omitempty"`

	// Custom certificate.
	// +kubebuilder:validation:Required
	CustomCert []CustomCertParameters `json:"customCert" tf:"custom_cert,omitempty"`
}

type CustomCertObservation struct {
}

type CustomCertParameters struct {

	// Public key of certificate.
	// +kubebuilder:validation:Required
	Cert *string `json:"cert" tf:"cert,omitempty"`

	// Private key of certificate.
	// +kubebuilder:validation:Required
	PrivateKey *string `json:"privateKey" tf:"private_key,omitempty"`
}

type DomainCertificateObservation struct {
}

type DomainCertificateParameters struct {

	// Certificate info.
	// +kubebuilder:validation:Required
	Certificate []CertificateParameters `json:"certificate" tf:"certificate,omitempty"`

	// The name of domain.
	// +kubebuilder:validation:Required
	Domain *string `json:"domain" tf:"domain,omitempty"`
}

// BucketDomainCertificateAttachmentSpec defines the desired state of BucketDomainCertificateAttachment
type BucketDomainCertificateAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketDomainCertificateAttachmentParameters `json:"forProvider"`
}

// BucketDomainCertificateAttachmentStatus defines the observed state of BucketDomainCertificateAttachment.
type BucketDomainCertificateAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketDomainCertificateAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BucketDomainCertificateAttachment is the Schema for the BucketDomainCertificateAttachments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type BucketDomainCertificateAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketDomainCertificateAttachmentSpec   `json:"spec"`
	Status            BucketDomainCertificateAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketDomainCertificateAttachmentList contains a list of BucketDomainCertificateAttachments
type BucketDomainCertificateAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketDomainCertificateAttachment `json:"items"`
}

// Repository type metadata.
var (
	BucketDomainCertificateAttachment_Kind             = "BucketDomainCertificateAttachment"
	BucketDomainCertificateAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketDomainCertificateAttachment_Kind}.String()
	BucketDomainCertificateAttachment_KindAPIVersion   = BucketDomainCertificateAttachment_Kind + "." + CRDGroupVersion.String()
	BucketDomainCertificateAttachment_GroupVersionKind = CRDGroupVersion.WithKind(BucketDomainCertificateAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketDomainCertificateAttachment{}, &BucketDomainCertificateAttachmentList{})
}
