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

type TmpInstanceInitParameters struct {

	// Data retention time(in days). Value range: 15, 30, 45, 90, 180, 360, 720.
	// Data retention time(in days). Value range: 15, 30, 45, 90, 180, 360, 720.
	DataRetentionTime *float64 `json:"dataRetentionTime,omitempty" tf:"data_retention_time,omitempty"`

	// Instance name.
	// Instance name.
	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name,omitempty"`

	// Subnet Id.
	// Subnet Id.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.Subnet
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// Tag description list.
	// Tag description list.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Vpc Id.
	// Vpc Id.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.VPC
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`

	// Available zone.
	// Available zone.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type TmpInstanceObservation struct {

	// Prometheus HTTP API root address.
	// Prometheus HTTP API root address.
	APIRootPath *string `json:"apiRootPath,omitempty" tf:"api_root_path,omitempty"`

	// Data retention time(in days). Value range: 15, 30, 45, 90, 180, 360, 720.
	// Data retention time(in days). Value range: 15, 30, 45, 90, 180, 360, 720.
	DataRetentionTime *float64 `json:"dataRetentionTime,omitempty" tf:"data_retention_time,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Instance IPv4 address.
	// Instance IPv4 address.
	IPv4Address *string `json:"ipv4Address,omitempty" tf:"ipv4_address,omitempty"`

	// Instance name.
	// Instance name.
	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name,omitempty"`

	// Proxy address.
	// Proxy address.
	ProxyAddress *string `json:"proxyAddress,omitempty" tf:"proxy_address,omitempty"`

	// Prometheus remote write address.
	// Prometheus remote write address.
	RemoteWrite *string `json:"remoteWrite,omitempty" tf:"remote_write,omitempty"`

	// Subnet Id.
	// Subnet Id.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Tag description list.
	// Tag description list.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Vpc Id.
	// Vpc Id.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Available zone.
	// Available zone.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type TmpInstanceParameters struct {

	// Data retention time(in days). Value range: 15, 30, 45, 90, 180, 360, 720.
	// Data retention time(in days). Value range: 15, 30, 45, 90, 180, 360, 720.
	// +kubebuilder:validation:Optional
	DataRetentionTime *float64 `json:"dataRetentionTime,omitempty" tf:"data_retention_time,omitempty"`

	// Instance name.
	// Instance name.
	// +kubebuilder:validation:Optional
	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name,omitempty"`

	// Subnet Id.
	// Subnet Id.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// Tag description list.
	// Tag description list.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Vpc Id.
	// Vpc Id.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`

	// Available zone.
	// Available zone.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// TmpInstanceSpec defines the desired state of TmpInstance
type TmpInstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TmpInstanceParameters `json:"forProvider"`
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
	InitProvider TmpInstanceInitParameters `json:"initProvider,omitempty"`
}

// TmpInstanceStatus defines the observed state of TmpInstance.
type TmpInstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TmpInstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// TmpInstance is the Schema for the TmpInstances API. Provides a resource to create a monitor tmpInstance
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type TmpInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.dataRetentionTime) || (has(self.initProvider) && has(self.initProvider.dataRetentionTime))",message="spec.forProvider.dataRetentionTime is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.instanceName) || (has(self.initProvider) && has(self.initProvider.instanceName))",message="spec.forProvider.instanceName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.zone) || (has(self.initProvider) && has(self.initProvider.zone))",message="spec.forProvider.zone is a required parameter"
	Spec   TmpInstanceSpec   `json:"spec"`
	Status TmpInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TmpInstanceList contains a list of TmpInstances
type TmpInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TmpInstance `json:"items"`
}

// Repository type metadata.
var (
	TmpInstance_Kind             = "TmpInstance"
	TmpInstance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TmpInstance_Kind}.String()
	TmpInstance_KindAPIVersion   = TmpInstance_Kind + "." + CRDGroupVersion.String()
	TmpInstance_GroupVersionKind = CRDGroupVersion.WithKind(TmpInstance_Kind)
)

func init() {
	SchemeBuilder.Register(&TmpInstance{}, &TmpInstanceList{})
}
