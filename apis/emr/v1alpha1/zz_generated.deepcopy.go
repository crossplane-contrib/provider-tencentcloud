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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonResourceSpecObservation) DeepCopyInto(out *CommonResourceSpecObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonResourceSpecObservation.
func (in *CommonResourceSpecObservation) DeepCopy() *CommonResourceSpecObservation {
	if in == nil {
		return nil
	}
	out := new(CommonResourceSpecObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonResourceSpecParameters) DeepCopyInto(out *CommonResourceSpecParameters) {
	*out = *in
	if in.CPU != nil {
		in, out := &in.CPU, &out.CPU
		*out = new(float64)
		**out = **in
	}
	if in.DiskSize != nil {
		in, out := &in.DiskSize, &out.DiskSize
		*out = new(float64)
		**out = **in
	}
	if in.DiskType != nil {
		in, out := &in.DiskType, &out.DiskType
		*out = new(string)
		**out = **in
	}
	if in.MemSize != nil {
		in, out := &in.MemSize, &out.MemSize
		*out = new(float64)
		**out = **in
	}
	if in.RootSize != nil {
		in, out := &in.RootSize, &out.RootSize
		*out = new(float64)
		**out = **in
	}
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(string)
		**out = **in
	}
	if in.StorageType != nil {
		in, out := &in.StorageType, &out.StorageType
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonResourceSpecParameters.
func (in *CommonResourceSpecParameters) DeepCopy() *CommonResourceSpecParameters {
	if in == nil {
		return nil
	}
	out := new(CommonResourceSpecParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoreResourceSpecObservation) DeepCopyInto(out *CoreResourceSpecObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoreResourceSpecObservation.
func (in *CoreResourceSpecObservation) DeepCopy() *CoreResourceSpecObservation {
	if in == nil {
		return nil
	}
	out := new(CoreResourceSpecObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoreResourceSpecParameters) DeepCopyInto(out *CoreResourceSpecParameters) {
	*out = *in
	if in.CPU != nil {
		in, out := &in.CPU, &out.CPU
		*out = new(float64)
		**out = **in
	}
	if in.DiskSize != nil {
		in, out := &in.DiskSize, &out.DiskSize
		*out = new(float64)
		**out = **in
	}
	if in.DiskType != nil {
		in, out := &in.DiskType, &out.DiskType
		*out = new(string)
		**out = **in
	}
	if in.MemSize != nil {
		in, out := &in.MemSize, &out.MemSize
		*out = new(float64)
		**out = **in
	}
	if in.RootSize != nil {
		in, out := &in.RootSize, &out.RootSize
		*out = new(float64)
		**out = **in
	}
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(string)
		**out = **in
	}
	if in.StorageType != nil {
		in, out := &in.StorageType, &out.StorageType
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoreResourceSpecParameters.
func (in *CoreResourceSpecParameters) DeepCopy() *CoreResourceSpecParameters {
	if in == nil {
		return nil
	}
	out := new(CoreResourceSpecParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmrCluster) DeepCopyInto(out *EmrCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmrCluster.
func (in *EmrCluster) DeepCopy() *EmrCluster {
	if in == nil {
		return nil
	}
	out := new(EmrCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EmrCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmrClusterList) DeepCopyInto(out *EmrClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EmrCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmrClusterList.
func (in *EmrClusterList) DeepCopy() *EmrClusterList {
	if in == nil {
		return nil
	}
	out := new(EmrClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EmrClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmrClusterObservation) DeepCopyInto(out *EmrClusterObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.InstanceID != nil {
		in, out := &in.InstanceID, &out.InstanceID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmrClusterObservation.
func (in *EmrClusterObservation) DeepCopy() *EmrClusterObservation {
	if in == nil {
		return nil
	}
	out := new(EmrClusterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmrClusterParameters) DeepCopyInto(out *EmrClusterParameters) {
	*out = *in
	if in.DisplayStrategy != nil {
		in, out := &in.DisplayStrategy, &out.DisplayStrategy
		*out = new(string)
		**out = **in
	}
	if in.ExtendFsField != nil {
		in, out := &in.ExtendFsField, &out.ExtendFsField
		*out = new(string)
		**out = **in
	}
	if in.InstanceName != nil {
		in, out := &in.InstanceName, &out.InstanceName
		*out = new(string)
		**out = **in
	}
	if in.LoginSettings != nil {
		in, out := &in.LoginSettings, &out.LoginSettings
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
	if in.NeedMasterWan != nil {
		in, out := &in.NeedMasterWan, &out.NeedMasterWan
		*out = new(string)
		**out = **in
	}
	if in.PayMode != nil {
		in, out := &in.PayMode, &out.PayMode
		*out = new(float64)
		**out = **in
	}
	if in.Placement != nil {
		in, out := &in.Placement, &out.Placement
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
	if in.ProductID != nil {
		in, out := &in.ProductID, &out.ProductID
		*out = new(float64)
		**out = **in
	}
	if in.ResourceSpec != nil {
		in, out := &in.ResourceSpec, &out.ResourceSpec
		*out = make([]ResourceSpecParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SgID != nil {
		in, out := &in.SgID, &out.SgID
		*out = new(string)
		**out = **in
	}
	if in.Softwares != nil {
		in, out := &in.Softwares, &out.Softwares
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SupportHa != nil {
		in, out := &in.SupportHa, &out.SupportHa
		*out = new(float64)
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
	if in.TimeSpan != nil {
		in, out := &in.TimeSpan, &out.TimeSpan
		*out = new(float64)
		**out = **in
	}
	if in.TimeUnit != nil {
		in, out := &in.TimeUnit, &out.TimeUnit
		*out = new(string)
		**out = **in
	}
	if in.VPCSettings != nil {
		in, out := &in.VPCSettings, &out.VPCSettings
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmrClusterParameters.
func (in *EmrClusterParameters) DeepCopy() *EmrClusterParameters {
	if in == nil {
		return nil
	}
	out := new(EmrClusterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmrClusterSpec) DeepCopyInto(out *EmrClusterSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmrClusterSpec.
func (in *EmrClusterSpec) DeepCopy() *EmrClusterSpec {
	if in == nil {
		return nil
	}
	out := new(EmrClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmrClusterStatus) DeepCopyInto(out *EmrClusterStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmrClusterStatus.
func (in *EmrClusterStatus) DeepCopy() *EmrClusterStatus {
	if in == nil {
		return nil
	}
	out := new(EmrClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MasterResourceSpecObservation) DeepCopyInto(out *MasterResourceSpecObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MasterResourceSpecObservation.
func (in *MasterResourceSpecObservation) DeepCopy() *MasterResourceSpecObservation {
	if in == nil {
		return nil
	}
	out := new(MasterResourceSpecObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MasterResourceSpecParameters) DeepCopyInto(out *MasterResourceSpecParameters) {
	*out = *in
	if in.CPU != nil {
		in, out := &in.CPU, &out.CPU
		*out = new(float64)
		**out = **in
	}
	if in.DiskSize != nil {
		in, out := &in.DiskSize, &out.DiskSize
		*out = new(float64)
		**out = **in
	}
	if in.DiskType != nil {
		in, out := &in.DiskType, &out.DiskType
		*out = new(string)
		**out = **in
	}
	if in.MemSize != nil {
		in, out := &in.MemSize, &out.MemSize
		*out = new(float64)
		**out = **in
	}
	if in.RootSize != nil {
		in, out := &in.RootSize, &out.RootSize
		*out = new(float64)
		**out = **in
	}
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(string)
		**out = **in
	}
	if in.StorageType != nil {
		in, out := &in.StorageType, &out.StorageType
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MasterResourceSpecParameters.
func (in *MasterResourceSpecParameters) DeepCopy() *MasterResourceSpecParameters {
	if in == nil {
		return nil
	}
	out := new(MasterResourceSpecParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSpecObservation) DeepCopyInto(out *ResourceSpecObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSpecObservation.
func (in *ResourceSpecObservation) DeepCopy() *ResourceSpecObservation {
	if in == nil {
		return nil
	}
	out := new(ResourceSpecObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSpecParameters) DeepCopyInto(out *ResourceSpecParameters) {
	*out = *in
	if in.CommonCount != nil {
		in, out := &in.CommonCount, &out.CommonCount
		*out = new(float64)
		**out = **in
	}
	if in.CommonResourceSpec != nil {
		in, out := &in.CommonResourceSpec, &out.CommonResourceSpec
		*out = make([]CommonResourceSpecParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CoreCount != nil {
		in, out := &in.CoreCount, &out.CoreCount
		*out = new(float64)
		**out = **in
	}
	if in.CoreResourceSpec != nil {
		in, out := &in.CoreResourceSpec, &out.CoreResourceSpec
		*out = make([]CoreResourceSpecParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MasterCount != nil {
		in, out := &in.MasterCount, &out.MasterCount
		*out = new(float64)
		**out = **in
	}
	if in.MasterResourceSpec != nil {
		in, out := &in.MasterResourceSpec, &out.MasterResourceSpec
		*out = make([]MasterResourceSpecParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TaskCount != nil {
		in, out := &in.TaskCount, &out.TaskCount
		*out = new(float64)
		**out = **in
	}
	if in.TaskResourceSpec != nil {
		in, out := &in.TaskResourceSpec, &out.TaskResourceSpec
		*out = make([]TaskResourceSpecParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSpecParameters.
func (in *ResourceSpecParameters) DeepCopy() *ResourceSpecParameters {
	if in == nil {
		return nil
	}
	out := new(ResourceSpecParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskResourceSpecObservation) DeepCopyInto(out *TaskResourceSpecObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskResourceSpecObservation.
func (in *TaskResourceSpecObservation) DeepCopy() *TaskResourceSpecObservation {
	if in == nil {
		return nil
	}
	out := new(TaskResourceSpecObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskResourceSpecParameters) DeepCopyInto(out *TaskResourceSpecParameters) {
	*out = *in
	if in.CPU != nil {
		in, out := &in.CPU, &out.CPU
		*out = new(float64)
		**out = **in
	}
	if in.DiskSize != nil {
		in, out := &in.DiskSize, &out.DiskSize
		*out = new(float64)
		**out = **in
	}
	if in.DiskType != nil {
		in, out := &in.DiskType, &out.DiskType
		*out = new(string)
		**out = **in
	}
	if in.MemSize != nil {
		in, out := &in.MemSize, &out.MemSize
		*out = new(float64)
		**out = **in
	}
	if in.RootSize != nil {
		in, out := &in.RootSize, &out.RootSize
		*out = new(float64)
		**out = **in
	}
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(string)
		**out = **in
	}
	if in.StorageType != nil {
		in, out := &in.StorageType, &out.StorageType
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskResourceSpecParameters.
func (in *TaskResourceSpecParameters) DeepCopy() *TaskResourceSpecParameters {
	if in == nil {
		return nil
	}
	out := new(TaskResourceSpecParameters)
	in.DeepCopyInto(out)
	return out
}
