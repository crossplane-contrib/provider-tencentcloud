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

type GatewayObservation struct {
	CnnRouteType *string `json:"cnnRouteType,omitempty" tf:"cnn_route_type,omitempty"`

	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	EnableBGP *bool `json:"enableBgp,omitempty" tf:"enable_bgp,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type GatewayParameters struct {

	// Type of the gateway. Valid value: `NORMAL` and `NAT`. Default is `NORMAL`. NOTES: CCN only supports `NORMAL` and a VPC can create two DCGs, the one is NAT type and the other is non-NAT type.
	// +kubebuilder:validation:Optional
	GatewayType *string `json:"gatewayType,omitempty" tf:"gateway_type,omitempty"`

	// Name of the DCG.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// If the `network_type` value is `VPC`, the available value is VPC ID. But when the `network_type` value is `CCN`, the available value is CCN instance ID.
	// +kubebuilder:validation:Required
	NetworkInstanceID *string `json:"networkInstanceId" tf:"network_instance_id,omitempty"`

	// Type of associated network. Valid value: `VPC` and `CCN`.
	// +kubebuilder:validation:Required
	NetworkType *string `json:"networkType" tf:"network_type,omitempty"`
}

// GatewaySpec defines the desired state of Gateway
type GatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GatewayParameters `json:"forProvider"`
}

// GatewayStatus defines the observed state of Gateway.
type GatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Gateway is the Schema for the Gateways API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type Gateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GatewaySpec   `json:"spec"`
	Status            GatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayList contains a list of Gateways
type GatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Gateway `json:"items"`
}

// Repository type metadata.
var (
	Gateway_Kind             = "Gateway"
	Gateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Gateway_Kind}.String()
	Gateway_KindAPIVersion   = Gateway_Kind + "." + CRDGroupVersion.String()
	Gateway_GroupVersionKind = CRDGroupVersion.WithKind(Gateway_Kind)
)

func init() {
	SchemeBuilder.Register(&Gateway{}, &GatewayList{})
}