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
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Instance) DeepCopyInto(out *Instance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Instance.
func (in *Instance) DeepCopy() *Instance {
	if in == nil {
		return nil
	}
	out := new(Instance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Instance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceList) DeepCopyInto(out *InstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Instance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceList.
func (in *InstanceList) DeepCopy() *InstanceList {
	if in == nil {
		return nil
	}
	out := new(InstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceObservation) DeepCopyInto(out *InstanceObservation) {
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
	if in.StandbyInstanceList != nil {
		in, out := &in.StandbyInstanceList, &out.StandbyInstanceList
		*out = make([]StandbyInstanceListObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(float64)
		**out = **in
	}
	if in.Vip != nil {
		in, out := &in.Vip, &out.Vip
		*out = new(string)
		**out = **in
	}
	if in.Vport != nil {
		in, out := &in.Vport, &out.Vport
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceObservation.
func (in *InstanceObservation) DeepCopy() *InstanceObservation {
	if in == nil {
		return nil
	}
	out := new(InstanceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceParameters) DeepCopyInto(out *InstanceParameters) {
	*out = *in
	if in.AutoRenewFlag != nil {
		in, out := &in.AutoRenewFlag, &out.AutoRenewFlag
		*out = new(float64)
		**out = **in
	}
	if in.AvailabilityZoneList != nil {
		in, out := &in.AvailabilityZoneList, &out.AvailabilityZoneList
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AvailableZone != nil {
		in, out := &in.AvailableZone, &out.AvailableZone
		*out = new(string)
		**out = **in
	}
	if in.ChargeType != nil {
		in, out := &in.ChargeType, &out.ChargeType
		*out = new(string)
		**out = **in
	}
	if in.EngineVersion != nil {
		in, out := &in.EngineVersion, &out.EngineVersion
		*out = new(string)
		**out = **in
	}
	if in.HiddenZone != nil {
		in, out := &in.HiddenZone, &out.HiddenZone
		*out = new(string)
		**out = **in
	}
	if in.InstanceName != nil {
		in, out := &in.InstanceName, &out.InstanceName
		*out = new(string)
		**out = **in
	}
	if in.MachineType != nil {
		in, out := &in.MachineType, &out.MachineType
		*out = new(string)
		**out = **in
	}
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		*out = new(float64)
		**out = **in
	}
	if in.NodeNum != nil {
		in, out := &in.NodeNum, &out.NodeNum
		*out = new(float64)
		**out = **in
	}
	if in.PasswordSecretRef != nil {
		in, out := &in.PasswordSecretRef, &out.PasswordSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.PrepaidPeriod != nil {
		in, out := &in.PrepaidPeriod, &out.PrepaidPeriod
		*out = new(float64)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(float64)
		**out = **in
	}
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
	if in.SubnetIDRef != nil {
		in, out := &in.SubnetIDRef, &out.SubnetIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.SubnetIDSelector != nil {
		in, out := &in.SubnetIDSelector, &out.SubnetIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.VPCID != nil {
		in, out := &in.VPCID, &out.VPCID
		*out = new(string)
		**out = **in
	}
	if in.VPCIDRef != nil {
		in, out := &in.VPCIDRef, &out.VPCIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.VPCIDSelector != nil {
		in, out := &in.VPCIDSelector, &out.VPCIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Volume != nil {
		in, out := &in.Volume, &out.Volume
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceParameters.
func (in *InstanceParameters) DeepCopy() *InstanceParameters {
	if in == nil {
		return nil
	}
	out := new(InstanceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceSpec) DeepCopyInto(out *InstanceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceSpec.
func (in *InstanceSpec) DeepCopy() *InstanceSpec {
	if in == nil {
		return nil
	}
	out := new(InstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceStatus) DeepCopyInto(out *InstanceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceStatus.
func (in *InstanceStatus) DeepCopy() *InstanceStatus {
	if in == nil {
		return nil
	}
	out := new(InstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShardingInstance) DeepCopyInto(out *ShardingInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShardingInstance.
func (in *ShardingInstance) DeepCopy() *ShardingInstance {
	if in == nil {
		return nil
	}
	out := new(ShardingInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ShardingInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShardingInstanceList) DeepCopyInto(out *ShardingInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ShardingInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShardingInstanceList.
func (in *ShardingInstanceList) DeepCopy() *ShardingInstanceList {
	if in == nil {
		return nil
	}
	out := new(ShardingInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ShardingInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShardingInstanceObservation) DeepCopyInto(out *ShardingInstanceObservation) {
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
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(float64)
		**out = **in
	}
	if in.Vip != nil {
		in, out := &in.Vip, &out.Vip
		*out = new(string)
		**out = **in
	}
	if in.Vport != nil {
		in, out := &in.Vport, &out.Vport
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShardingInstanceObservation.
func (in *ShardingInstanceObservation) DeepCopy() *ShardingInstanceObservation {
	if in == nil {
		return nil
	}
	out := new(ShardingInstanceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShardingInstanceParameters) DeepCopyInto(out *ShardingInstanceParameters) {
	*out = *in
	if in.AutoRenewFlag != nil {
		in, out := &in.AutoRenewFlag, &out.AutoRenewFlag
		*out = new(float64)
		**out = **in
	}
	if in.AvailabilityZoneList != nil {
		in, out := &in.AvailabilityZoneList, &out.AvailabilityZoneList
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AvailableZone != nil {
		in, out := &in.AvailableZone, &out.AvailableZone
		*out = new(string)
		**out = **in
	}
	if in.ChargeType != nil {
		in, out := &in.ChargeType, &out.ChargeType
		*out = new(string)
		**out = **in
	}
	if in.EngineVersion != nil {
		in, out := &in.EngineVersion, &out.EngineVersion
		*out = new(string)
		**out = **in
	}
	if in.HiddenZone != nil {
		in, out := &in.HiddenZone, &out.HiddenZone
		*out = new(string)
		**out = **in
	}
	if in.InstanceName != nil {
		in, out := &in.InstanceName, &out.InstanceName
		*out = new(string)
		**out = **in
	}
	if in.MachineType != nil {
		in, out := &in.MachineType, &out.MachineType
		*out = new(string)
		**out = **in
	}
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		*out = new(float64)
		**out = **in
	}
	if in.MongosCPU != nil {
		in, out := &in.MongosCPU, &out.MongosCPU
		*out = new(float64)
		**out = **in
	}
	if in.MongosMemory != nil {
		in, out := &in.MongosMemory, &out.MongosMemory
		*out = new(float64)
		**out = **in
	}
	if in.MongosNodeNum != nil {
		in, out := &in.MongosNodeNum, &out.MongosNodeNum
		*out = new(float64)
		**out = **in
	}
	if in.NodesPerShard != nil {
		in, out := &in.NodesPerShard, &out.NodesPerShard
		*out = new(float64)
		**out = **in
	}
	if in.PasswordSecretRef != nil {
		in, out := &in.PasswordSecretRef, &out.PasswordSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.PrepaidPeriod != nil {
		in, out := &in.PrepaidPeriod, &out.PrepaidPeriod
		*out = new(float64)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(float64)
		**out = **in
	}
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ShardQuantity != nil {
		in, out := &in.ShardQuantity, &out.ShardQuantity
		*out = new(float64)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
	if in.SubnetIDRef != nil {
		in, out := &in.SubnetIDRef, &out.SubnetIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.SubnetIDSelector != nil {
		in, out := &in.SubnetIDSelector, &out.SubnetIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.VPCID != nil {
		in, out := &in.VPCID, &out.VPCID
		*out = new(string)
		**out = **in
	}
	if in.VPCIDRef != nil {
		in, out := &in.VPCIDRef, &out.VPCIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.VPCIDSelector != nil {
		in, out := &in.VPCIDSelector, &out.VPCIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Volume != nil {
		in, out := &in.Volume, &out.Volume
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShardingInstanceParameters.
func (in *ShardingInstanceParameters) DeepCopy() *ShardingInstanceParameters {
	if in == nil {
		return nil
	}
	out := new(ShardingInstanceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShardingInstanceSpec) DeepCopyInto(out *ShardingInstanceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShardingInstanceSpec.
func (in *ShardingInstanceSpec) DeepCopy() *ShardingInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(ShardingInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShardingInstanceStatus) DeepCopyInto(out *ShardingInstanceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShardingInstanceStatus.
func (in *ShardingInstanceStatus) DeepCopy() *ShardingInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(ShardingInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StandbyInstance) DeepCopyInto(out *StandbyInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StandbyInstance.
func (in *StandbyInstance) DeepCopy() *StandbyInstance {
	if in == nil {
		return nil
	}
	out := new(StandbyInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StandbyInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StandbyInstanceList) DeepCopyInto(out *StandbyInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StandbyInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StandbyInstanceList.
func (in *StandbyInstanceList) DeepCopy() *StandbyInstanceList {
	if in == nil {
		return nil
	}
	out := new(StandbyInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StandbyInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StandbyInstanceListObservation) DeepCopyInto(out *StandbyInstanceListObservation) {
	*out = *in
	if in.StandbyInstanceID != nil {
		in, out := &in.StandbyInstanceID, &out.StandbyInstanceID
		*out = new(string)
		**out = **in
	}
	if in.StandbyInstanceRegion != nil {
		in, out := &in.StandbyInstanceRegion, &out.StandbyInstanceRegion
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StandbyInstanceListObservation.
func (in *StandbyInstanceListObservation) DeepCopy() *StandbyInstanceListObservation {
	if in == nil {
		return nil
	}
	out := new(StandbyInstanceListObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StandbyInstanceListParameters) DeepCopyInto(out *StandbyInstanceListParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StandbyInstanceListParameters.
func (in *StandbyInstanceListParameters) DeepCopy() *StandbyInstanceListParameters {
	if in == nil {
		return nil
	}
	out := new(StandbyInstanceListParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StandbyInstanceObservation) DeepCopyInto(out *StandbyInstanceObservation) {
	*out = *in
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.EngineVersion != nil {
		in, out := &in.EngineVersion, &out.EngineVersion
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.MachineType != nil {
		in, out := &in.MachineType, &out.MachineType
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(float64)
		**out = **in
	}
	if in.Vip != nil {
		in, out := &in.Vip, &out.Vip
		*out = new(string)
		**out = **in
	}
	if in.Vport != nil {
		in, out := &in.Vport, &out.Vport
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StandbyInstanceObservation.
func (in *StandbyInstanceObservation) DeepCopy() *StandbyInstanceObservation {
	if in == nil {
		return nil
	}
	out := new(StandbyInstanceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StandbyInstanceParameters) DeepCopyInto(out *StandbyInstanceParameters) {
	*out = *in
	if in.AutoRenewFlag != nil {
		in, out := &in.AutoRenewFlag, &out.AutoRenewFlag
		*out = new(float64)
		**out = **in
	}
	if in.AvailableZone != nil {
		in, out := &in.AvailableZone, &out.AvailableZone
		*out = new(string)
		**out = **in
	}
	if in.ChargeType != nil {
		in, out := &in.ChargeType, &out.ChargeType
		*out = new(string)
		**out = **in
	}
	if in.FatherInstanceID != nil {
		in, out := &in.FatherInstanceID, &out.FatherInstanceID
		*out = new(string)
		**out = **in
	}
	if in.FatherInstanceIDRef != nil {
		in, out := &in.FatherInstanceIDRef, &out.FatherInstanceIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.FatherInstanceIDSelector != nil {
		in, out := &in.FatherInstanceIDSelector, &out.FatherInstanceIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.FatherInstanceRegion != nil {
		in, out := &in.FatherInstanceRegion, &out.FatherInstanceRegion
		*out = new(string)
		**out = **in
	}
	if in.InstanceName != nil {
		in, out := &in.InstanceName, &out.InstanceName
		*out = new(string)
		**out = **in
	}
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		*out = new(float64)
		**out = **in
	}
	if in.PrepaidPeriod != nil {
		in, out := &in.PrepaidPeriod, &out.PrepaidPeriod
		*out = new(float64)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(float64)
		**out = **in
	}
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.VPCID != nil {
		in, out := &in.VPCID, &out.VPCID
		*out = new(string)
		**out = **in
	}
	if in.Volume != nil {
		in, out := &in.Volume, &out.Volume
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StandbyInstanceParameters.
func (in *StandbyInstanceParameters) DeepCopy() *StandbyInstanceParameters {
	if in == nil {
		return nil
	}
	out := new(StandbyInstanceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StandbyInstanceSpec) DeepCopyInto(out *StandbyInstanceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StandbyInstanceSpec.
func (in *StandbyInstanceSpec) DeepCopy() *StandbyInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(StandbyInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StandbyInstanceStatus) DeepCopyInto(out *StandbyInstanceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StandbyInstanceStatus.
func (in *StandbyInstanceStatus) DeepCopy() *StandbyInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(StandbyInstanceStatus)
	in.DeepCopyInto(out)
	return out
}
