//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Dcx) DeepCopyInto(out *Dcx) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Dcx.
func (in *Dcx) DeepCopy() *Dcx {
	if in == nil {
		return nil
	}
	out := new(Dcx)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Dcx) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DcxList) DeepCopyInto(out *DcxList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Dcx, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DcxList.
func (in *DcxList) DeepCopy() *DcxList {
	if in == nil {
		return nil
	}
	out := new(DcxList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DcxList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DcxObservation) DeepCopyInto(out *DcxObservation) {
	*out = *in
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DcxObservation.
func (in *DcxObservation) DeepCopy() *DcxObservation {
	if in == nil {
		return nil
	}
	out := new(DcxObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DcxParameters) DeepCopyInto(out *DcxParameters) {
	*out = *in
	if in.BGPAsn != nil {
		in, out := &in.BGPAsn, &out.BGPAsn
		*out = new(float64)
		**out = **in
	}
	if in.BGPAuthKey != nil {
		in, out := &in.BGPAuthKey, &out.BGPAuthKey
		*out = new(string)
		**out = **in
	}
	if in.Bandwidth != nil {
		in, out := &in.Bandwidth, &out.Bandwidth
		*out = new(float64)
		**out = **in
	}
	if in.CustomerAddress != nil {
		in, out := &in.CustomerAddress, &out.CustomerAddress
		*out = new(string)
		**out = **in
	}
	if in.DcID != nil {
		in, out := &in.DcID, &out.DcID
		*out = new(string)
		**out = **in
	}
	if in.DcgID != nil {
		in, out := &in.DcgID, &out.DcgID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NetworkType != nil {
		in, out := &in.NetworkType, &out.NetworkType
		*out = new(string)
		**out = **in
	}
	if in.RouteFilterPrefixes != nil {
		in, out := &in.RouteFilterPrefixes, &out.RouteFilterPrefixes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.RouteType != nil {
		in, out := &in.RouteType, &out.RouteType
		*out = new(string)
		**out = **in
	}
	if in.TencentAddress != nil {
		in, out := &in.TencentAddress, &out.TencentAddress
		*out = new(string)
		**out = **in
	}
	if in.VPCID != nil {
		in, out := &in.VPCID, &out.VPCID
		*out = new(string)
		**out = **in
	}
	if in.Vlan != nil {
		in, out := &in.Vlan, &out.Vlan
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DcxParameters.
func (in *DcxParameters) DeepCopy() *DcxParameters {
	if in == nil {
		return nil
	}
	out := new(DcxParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DcxSpec) DeepCopyInto(out *DcxSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DcxSpec.
func (in *DcxSpec) DeepCopy() *DcxSpec {
	if in == nil {
		return nil
	}
	out := new(DcxSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DcxStatus) DeepCopyInto(out *DcxStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DcxStatus.
func (in *DcxStatus) DeepCopy() *DcxStatus {
	if in == nil {
		return nil
	}
	out := new(DcxStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Gateway) DeepCopyInto(out *Gateway) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Gateway.
func (in *Gateway) DeepCopy() *Gateway {
	if in == nil {
		return nil
	}
	out := new(Gateway)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Gateway) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayCcnRoute) DeepCopyInto(out *GatewayCcnRoute) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayCcnRoute.
func (in *GatewayCcnRoute) DeepCopy() *GatewayCcnRoute {
	if in == nil {
		return nil
	}
	out := new(GatewayCcnRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GatewayCcnRoute) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayCcnRouteList) DeepCopyInto(out *GatewayCcnRouteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GatewayCcnRoute, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayCcnRouteList.
func (in *GatewayCcnRouteList) DeepCopy() *GatewayCcnRouteList {
	if in == nil {
		return nil
	}
	out := new(GatewayCcnRouteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GatewayCcnRouteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayCcnRouteObservation) DeepCopyInto(out *GatewayCcnRouteObservation) {
	*out = *in
	if in.AsPath != nil {
		in, out := &in.AsPath, &out.AsPath
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayCcnRouteObservation.
func (in *GatewayCcnRouteObservation) DeepCopy() *GatewayCcnRouteObservation {
	if in == nil {
		return nil
	}
	out := new(GatewayCcnRouteObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayCcnRouteParameters) DeepCopyInto(out *GatewayCcnRouteParameters) {
	*out = *in
	if in.CidrBlock != nil {
		in, out := &in.CidrBlock, &out.CidrBlock
		*out = new(string)
		**out = **in
	}
	if in.DcgID != nil {
		in, out := &in.DcgID, &out.DcgID
		*out = new(string)
		**out = **in
	}
	if in.DcgIDRef != nil {
		in, out := &in.DcgIDRef, &out.DcgIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.DcgIDSelector != nil {
		in, out := &in.DcgIDSelector, &out.DcgIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayCcnRouteParameters.
func (in *GatewayCcnRouteParameters) DeepCopy() *GatewayCcnRouteParameters {
	if in == nil {
		return nil
	}
	out := new(GatewayCcnRouteParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayCcnRouteSpec) DeepCopyInto(out *GatewayCcnRouteSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayCcnRouteSpec.
func (in *GatewayCcnRouteSpec) DeepCopy() *GatewayCcnRouteSpec {
	if in == nil {
		return nil
	}
	out := new(GatewayCcnRouteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayCcnRouteStatus) DeepCopyInto(out *GatewayCcnRouteStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayCcnRouteStatus.
func (in *GatewayCcnRouteStatus) DeepCopy() *GatewayCcnRouteStatus {
	if in == nil {
		return nil
	}
	out := new(GatewayCcnRouteStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayList) DeepCopyInto(out *GatewayList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Gateway, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayList.
func (in *GatewayList) DeepCopy() *GatewayList {
	if in == nil {
		return nil
	}
	out := new(GatewayList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GatewayList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayObservation) DeepCopyInto(out *GatewayObservation) {
	*out = *in
	if in.CnnRouteType != nil {
		in, out := &in.CnnRouteType, &out.CnnRouteType
		*out = new(string)
		**out = **in
	}
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.EnableBGP != nil {
		in, out := &in.EnableBGP, &out.EnableBGP
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayObservation.
func (in *GatewayObservation) DeepCopy() *GatewayObservation {
	if in == nil {
		return nil
	}
	out := new(GatewayObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayParameters) DeepCopyInto(out *GatewayParameters) {
	*out = *in
	if in.GatewayType != nil {
		in, out := &in.GatewayType, &out.GatewayType
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NetworkInstanceID != nil {
		in, out := &in.NetworkInstanceID, &out.NetworkInstanceID
		*out = new(string)
		**out = **in
	}
	if in.NetworkType != nil {
		in, out := &in.NetworkType, &out.NetworkType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayParameters.
func (in *GatewayParameters) DeepCopy() *GatewayParameters {
	if in == nil {
		return nil
	}
	out := new(GatewayParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewaySpec) DeepCopyInto(out *GatewaySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewaySpec.
func (in *GatewaySpec) DeepCopy() *GatewaySpec {
	if in == nil {
		return nil
	}
	out := new(GatewaySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayStatus) DeepCopyInto(out *GatewayStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayStatus.
func (in *GatewayStatus) DeepCopy() *GatewayStatus {
	if in == nil {
		return nil
	}
	out := new(GatewayStatus)
	in.DeepCopyInto(out)
	return out
}