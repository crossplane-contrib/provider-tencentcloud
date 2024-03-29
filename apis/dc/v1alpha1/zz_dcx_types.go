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

type DcxObservation struct {
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type DcxParameters struct {

	// BGP ASN of the user. A required field within BGP.
	// +kubebuilder:validation:Optional
	BGPAsn *float64 `json:"bgpAsn,omitempty" tf:"bgp_asn,omitempty"`

	// BGP key of the user.
	// +kubebuilder:validation:Optional
	BGPAuthKey *string `json:"bgpAuthKey,omitempty" tf:"bgp_auth_key,omitempty"`

	// Bandwidth of the DC.
	// +kubebuilder:validation:Optional
	Bandwidth *float64 `json:"bandwidth,omitempty" tf:"bandwidth,omitempty"`

	// Interconnect IP of the DC within client.
	// +kubebuilder:validation:Optional
	CustomerAddress *string `json:"customerAddress,omitempty" tf:"customer_address,omitempty"`

	// ID of the DC to be queried, application deployment offline.
	// +kubebuilder:validation:Required
	DcID *string `json:"dcId" tf:"dc_id,omitempty"`

	// Connection owner, who is the current customer by default. The developer account ID should be entered for shared connections.
	// +kubebuilder:validation:Optional
	DcOwnerAccount *string `json:"dcOwnerAccount,omitempty" tf:"dc_owner_account,omitempty"`

	// ID of the DC Gateway. Currently only new in the console.
	// +kubebuilder:validation:Required
	DcgID *string `json:"dcgId" tf:"dcg_id,omitempty"`

	// Name of the dedicated tunnel.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Type of the network. Valid value: `VPC`, `BMVPC` and `CCN`. The default value is `VPC`.
	// +kubebuilder:validation:Optional
	NetworkType *string `json:"networkType,omitempty" tf:"network_type,omitempty"`

	// Static route, the network address of the user IDC. It can be modified after setting but cannot be deleted. AN unable field within BGP.
	// +kubebuilder:validation:Optional
	RouteFilterPrefixes []*string `json:"routeFilterPrefixes,omitempty" tf:"route_filter_prefixes,omitempty"`

	// Type of the route, and available values include BGP and STATIC. The default value is `BGP`.
	// +kubebuilder:validation:Optional
	RouteType *string `json:"routeType,omitempty" tf:"route_type,omitempty"`

	// Interconnect IP of the DC within Tencent.
	// +kubebuilder:validation:Optional
	TencentAddress *string `json:"tencentAddress,omitempty" tf:"tencent_address,omitempty"`

	// ID of the VPC or BMVPC.
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Vlan of the dedicated tunnels. Valid value ranges: (0~3000). `0` means that only one tunnel can be created for the physical connect.
	// +kubebuilder:validation:Optional
	Vlan *float64 `json:"vlan,omitempty" tf:"vlan,omitempty"`
}

// DcxSpec defines the desired state of Dcx
type DcxSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DcxParameters `json:"forProvider"`
}

// DcxStatus defines the observed state of Dcx.
type DcxStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DcxObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Dcx is the Schema for the Dcxs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type Dcx struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DcxSpec   `json:"spec"`
	Status            DcxStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DcxList contains a list of Dcxs
type DcxList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Dcx `json:"items"`
}

// Repository type metadata.
var (
	Dcx_Kind             = "Dcx"
	Dcx_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Dcx_Kind}.String()
	Dcx_KindAPIVersion   = Dcx_Kind + "." + CRDGroupVersion.String()
	Dcx_GroupVersionKind = CRDGroupVersion.WithKind(Dcx_Kind)
)

func init() {
	SchemeBuilder.Register(&Dcx{}, &DcxList{})
}
