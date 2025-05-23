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

type NatGatewayInitParameters struct {

	// EIP IP address set bound to the gateway. The value of at least 1 and at most 10 if do not apply for a whitelist.
	// EIP IP address set bound to the gateway. The value of at least 1 and at most 10 if do not apply for a whitelist.
	// +listType=set
	AssignedEIPSet []*string `json:"assignedEipSet,omitempty" tf:"assigned_eip_set,omitempty"`

	// The maximum public network output bandwidth of NAT gateway (unit: Mbps). Valid values: 20, 50, 100, 200, 500, 1000, 2000, 5000. Default is 100. When the value of parameter nat_product_version is 2, which is the standard NAT type, this parameter does not need to be filled in and defaults to 5000.
	// The maximum public network output bandwidth of NAT gateway (unit: Mbps). Valid values: `20`, `50`, `100`, `200`, `500`, `1000`, `2000`, `5000`. Default is `100`. When the value of parameter `nat_product_version` is 2, which is the standard NAT type, this parameter does not need to be filled in and defaults to `5000`.
	Bandwidth *float64 `json:"bandwidth,omitempty" tf:"bandwidth,omitempty"`

	// The upper limit of concurrent connection of NAT gateway. Valid values: 1000000, 3000000, 10000000. Default is 1000000. When the value of parameter nat_product_version is 2, which is the standard NAT type, this parameter does not need to be filled in and defaults to 2000000.
	// The upper limit of concurrent connection of NAT gateway. Valid values: `1000000`, `3000000`, `10000000`. Default is `1000000`. When the value of parameter `nat_product_version` is 2, which is the standard NAT type, this parameter does not need to be filled in and defaults to `2000000`.
	MaxConcurrent *float64 `json:"maxConcurrent,omitempty" tf:"max_concurrent,omitempty"`

	// 1: traditional NAT, 2: standard NAT, default value is 1.
	// 1: traditional NAT, 2: standard NAT, default value is 1.
	NATProductVersion *float64 `json:"natProductVersion,omitempty" tf:"nat_product_version,omitempty"`

	// Name of the NAT gateway.
	// Name of the NAT gateway.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The elastic public IP bandwidth value (unit: Mbps) for binding NAT gateway. When this parameter is not filled in, it defaults to the bandwidth value of the elastic public IP, and for some users, it defaults to the bandwidth limit of the elastic public IP of that user type.
	// The elastic public IP bandwidth value (unit: Mbps) for binding NAT gateway. When this parameter is not filled in, it defaults to the bandwidth value of the elastic public IP, and for some users, it defaults to the bandwidth limit of the elastic public IP of that user type.
	StockPublicIPAddressesBandwidthOut *float64 `json:"stockPublicIpAddressesBandwidthOut,omitempty" tf:"stock_public_ip_addresses_bandwidth_out,omitempty"`

	// Subnet of NAT.
	// Subnet of NAT.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// The available tags within this NAT gateway.
	// The available tags within this NAT gateway.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// ID of the vpc.
	// ID of the vpc.
	// +crossplane:generate:reference:type=VPC
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`

	// The availability zone, such as ap-guangzhou-3.
	// The availability zone, such as `ap-guangzhou-3`.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type NatGatewayObservation struct {

	// EIP IP address set bound to the gateway. The value of at least 1 and at most 10 if do not apply for a whitelist.
	// EIP IP address set bound to the gateway. The value of at least 1 and at most 10 if do not apply for a whitelist.
	// +listType=set
	AssignedEIPSet []*string `json:"assignedEipSet,omitempty" tf:"assigned_eip_set,omitempty"`

	// The maximum public network output bandwidth of NAT gateway (unit: Mbps). Valid values: 20, 50, 100, 200, 500, 1000, 2000, 5000. Default is 100. When the value of parameter nat_product_version is 2, which is the standard NAT type, this parameter does not need to be filled in and defaults to 5000.
	// The maximum public network output bandwidth of NAT gateway (unit: Mbps). Valid values: `20`, `50`, `100`, `200`, `500`, `1000`, `2000`, `5000`. Default is `100`. When the value of parameter `nat_product_version` is 2, which is the standard NAT type, this parameter does not need to be filled in and defaults to `5000`.
	Bandwidth *float64 `json:"bandwidth,omitempty" tf:"bandwidth,omitempty"`

	// Create time of the NAT gateway.
	// Create time of the NAT gateway.
	CreatedTime *string `json:"createdTime,omitempty" tf:"created_time,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The upper limit of concurrent connection of NAT gateway. Valid values: 1000000, 3000000, 10000000. Default is 1000000. When the value of parameter nat_product_version is 2, which is the standard NAT type, this parameter does not need to be filled in and defaults to 2000000.
	// The upper limit of concurrent connection of NAT gateway. Valid values: `1000000`, `3000000`, `10000000`. Default is `1000000`. When the value of parameter `nat_product_version` is 2, which is the standard NAT type, this parameter does not need to be filled in and defaults to `2000000`.
	MaxConcurrent *float64 `json:"maxConcurrent,omitempty" tf:"max_concurrent,omitempty"`

	// 1: traditional NAT, 2: standard NAT, default value is 1.
	// 1: traditional NAT, 2: standard NAT, default value is 1.
	NATProductVersion *float64 `json:"natProductVersion,omitempty" tf:"nat_product_version,omitempty"`

	// Name of the NAT gateway.
	// Name of the NAT gateway.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The elastic public IP bandwidth value (unit: Mbps) for binding NAT gateway. When this parameter is not filled in, it defaults to the bandwidth value of the elastic public IP, and for some users, it defaults to the bandwidth limit of the elastic public IP of that user type.
	// The elastic public IP bandwidth value (unit: Mbps) for binding NAT gateway. When this parameter is not filled in, it defaults to the bandwidth value of the elastic public IP, and for some users, it defaults to the bandwidth limit of the elastic public IP of that user type.
	StockPublicIPAddressesBandwidthOut *float64 `json:"stockPublicIpAddressesBandwidthOut,omitempty" tf:"stock_public_ip_addresses_bandwidth_out,omitempty"`

	// Subnet of NAT.
	// Subnet of NAT.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// The available tags within this NAT gateway.
	// The available tags within this NAT gateway.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// ID of the vpc.
	// ID of the vpc.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// The availability zone, such as ap-guangzhou-3.
	// The availability zone, such as `ap-guangzhou-3`.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type NatGatewayParameters struct {

	// EIP IP address set bound to the gateway. The value of at least 1 and at most 10 if do not apply for a whitelist.
	// EIP IP address set bound to the gateway. The value of at least 1 and at most 10 if do not apply for a whitelist.
	// +kubebuilder:validation:Optional
	// +listType=set
	AssignedEIPSet []*string `json:"assignedEipSet,omitempty" tf:"assigned_eip_set,omitempty"`

	// The maximum public network output bandwidth of NAT gateway (unit: Mbps). Valid values: 20, 50, 100, 200, 500, 1000, 2000, 5000. Default is 100. When the value of parameter nat_product_version is 2, which is the standard NAT type, this parameter does not need to be filled in and defaults to 5000.
	// The maximum public network output bandwidth of NAT gateway (unit: Mbps). Valid values: `20`, `50`, `100`, `200`, `500`, `1000`, `2000`, `5000`. Default is `100`. When the value of parameter `nat_product_version` is 2, which is the standard NAT type, this parameter does not need to be filled in and defaults to `5000`.
	// +kubebuilder:validation:Optional
	Bandwidth *float64 `json:"bandwidth,omitempty" tf:"bandwidth,omitempty"`

	// The upper limit of concurrent connection of NAT gateway. Valid values: 1000000, 3000000, 10000000. Default is 1000000. When the value of parameter nat_product_version is 2, which is the standard NAT type, this parameter does not need to be filled in and defaults to 2000000.
	// The upper limit of concurrent connection of NAT gateway. Valid values: `1000000`, `3000000`, `10000000`. Default is `1000000`. When the value of parameter `nat_product_version` is 2, which is the standard NAT type, this parameter does not need to be filled in and defaults to `2000000`.
	// +kubebuilder:validation:Optional
	MaxConcurrent *float64 `json:"maxConcurrent,omitempty" tf:"max_concurrent,omitempty"`

	// 1: traditional NAT, 2: standard NAT, default value is 1.
	// 1: traditional NAT, 2: standard NAT, default value is 1.
	// +kubebuilder:validation:Optional
	NATProductVersion *float64 `json:"natProductVersion,omitempty" tf:"nat_product_version,omitempty"`

	// Name of the NAT gateway.
	// Name of the NAT gateway.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The elastic public IP bandwidth value (unit: Mbps) for binding NAT gateway. When this parameter is not filled in, it defaults to the bandwidth value of the elastic public IP, and for some users, it defaults to the bandwidth limit of the elastic public IP of that user type.
	// The elastic public IP bandwidth value (unit: Mbps) for binding NAT gateway. When this parameter is not filled in, it defaults to the bandwidth value of the elastic public IP, and for some users, it defaults to the bandwidth limit of the elastic public IP of that user type.
	// +kubebuilder:validation:Optional
	StockPublicIPAddressesBandwidthOut *float64 `json:"stockPublicIpAddressesBandwidthOut,omitempty" tf:"stock_public_ip_addresses_bandwidth_out,omitempty"`

	// Subnet of NAT.
	// Subnet of NAT.
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// The available tags within this NAT gateway.
	// The available tags within this NAT gateway.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// ID of the vpc.
	// ID of the vpc.
	// +crossplane:generate:reference:type=VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`

	// The availability zone, such as ap-guangzhou-3.
	// The availability zone, such as `ap-guangzhou-3`.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// NatGatewaySpec defines the desired state of NatGateway
type NatGatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NatGatewayParameters `json:"forProvider"`
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
	InitProvider NatGatewayInitParameters `json:"initProvider,omitempty"`
}

// NatGatewayStatus defines the observed state of NatGateway.
type NatGatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NatGatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// NatGateway is the Schema for the NatGateways API. Provides a resource to create a NAT gateway.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type NatGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.assignedEipSet) || (has(self.initProvider) && has(self.initProvider.assignedEipSet))",message="spec.forProvider.assignedEipSet is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   NatGatewaySpec   `json:"spec"`
	Status NatGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NatGatewayList contains a list of NatGateways
type NatGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NatGateway `json:"items"`
}

// Repository type metadata.
var (
	NatGateway_Kind             = "NatGateway"
	NatGateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NatGateway_Kind}.String()
	NatGateway_KindAPIVersion   = NatGateway_Kind + "." + CRDGroupVersion.String()
	NatGateway_GroupVersionKind = CRDGroupVersion.WithKind(NatGateway_Kind)
)

func init() {
	SchemeBuilder.Register(&NatGateway{}, &NatGatewayList{})
}
