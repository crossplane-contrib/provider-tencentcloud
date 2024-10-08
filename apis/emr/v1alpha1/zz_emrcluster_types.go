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

type CommonResourceSpecInitParameters struct {
	CPU *float64 `json:"cpu,omitempty" tf:"cpu,omitempty"`

	DiskSize *float64 `json:"diskSize,omitempty" tf:"disk_size,omitempty"`

	DiskType *string `json:"diskType,omitempty" tf:"disk_type,omitempty"`

	MemSize *float64 `json:"memSize,omitempty" tf:"mem_size,omitempty"`

	RootSize *float64 `json:"rootSize,omitempty" tf:"root_size,omitempty"`

	Spec *string `json:"spec,omitempty" tf:"spec,omitempty"`

	StorageType *float64 `json:"storageType,omitempty" tf:"storage_type,omitempty"`
}

type CommonResourceSpecObservation struct {
	CPU *float64 `json:"cpu,omitempty" tf:"cpu,omitempty"`

	DiskSize *float64 `json:"diskSize,omitempty" tf:"disk_size,omitempty"`

	DiskType *string `json:"diskType,omitempty" tf:"disk_type,omitempty"`

	MemSize *float64 `json:"memSize,omitempty" tf:"mem_size,omitempty"`

	RootSize *float64 `json:"rootSize,omitempty" tf:"root_size,omitempty"`

	Spec *string `json:"spec,omitempty" tf:"spec,omitempty"`

	StorageType *float64 `json:"storageType,omitempty" tf:"storage_type,omitempty"`
}

type CommonResourceSpecParameters struct {

	// +kubebuilder:validation:Optional
	CPU *float64 `json:"cpu,omitempty" tf:"cpu,omitempty"`

	// +kubebuilder:validation:Optional
	DiskSize *float64 `json:"diskSize,omitempty" tf:"disk_size,omitempty"`

	// +kubebuilder:validation:Optional
	DiskType *string `json:"diskType,omitempty" tf:"disk_type,omitempty"`

	// +kubebuilder:validation:Optional
	MemSize *float64 `json:"memSize,omitempty" tf:"mem_size,omitempty"`

	// +kubebuilder:validation:Optional
	RootSize *float64 `json:"rootSize,omitempty" tf:"root_size,omitempty"`

	// +kubebuilder:validation:Optional
	Spec *string `json:"spec,omitempty" tf:"spec,omitempty"`

	// +kubebuilder:validation:Optional
	StorageType *float64 `json:"storageType,omitempty" tf:"storage_type,omitempty"`
}

type CoreResourceSpecInitParameters struct {
	CPU *float64 `json:"cpu,omitempty" tf:"cpu,omitempty"`

	DiskSize *float64 `json:"diskSize,omitempty" tf:"disk_size,omitempty"`

	DiskType *string `json:"diskType,omitempty" tf:"disk_type,omitempty"`

	MemSize *float64 `json:"memSize,omitempty" tf:"mem_size,omitempty"`

	RootSize *float64 `json:"rootSize,omitempty" tf:"root_size,omitempty"`

	Spec *string `json:"spec,omitempty" tf:"spec,omitempty"`

	StorageType *float64 `json:"storageType,omitempty" tf:"storage_type,omitempty"`
}

type CoreResourceSpecObservation struct {
	CPU *float64 `json:"cpu,omitempty" tf:"cpu,omitempty"`

	DiskSize *float64 `json:"diskSize,omitempty" tf:"disk_size,omitempty"`

	DiskType *string `json:"diskType,omitempty" tf:"disk_type,omitempty"`

	MemSize *float64 `json:"memSize,omitempty" tf:"mem_size,omitempty"`

	RootSize *float64 `json:"rootSize,omitempty" tf:"root_size,omitempty"`

	Spec *string `json:"spec,omitempty" tf:"spec,omitempty"`

	StorageType *float64 `json:"storageType,omitempty" tf:"storage_type,omitempty"`
}

type CoreResourceSpecParameters struct {

	// +kubebuilder:validation:Optional
	CPU *float64 `json:"cpu,omitempty" tf:"cpu,omitempty"`

	// +kubebuilder:validation:Optional
	DiskSize *float64 `json:"diskSize,omitempty" tf:"disk_size,omitempty"`

	// +kubebuilder:validation:Optional
	DiskType *string `json:"diskType,omitempty" tf:"disk_type,omitempty"`

	// +kubebuilder:validation:Optional
	MemSize *float64 `json:"memSize,omitempty" tf:"mem_size,omitempty"`

	// +kubebuilder:validation:Optional
	RootSize *float64 `json:"rootSize,omitempty" tf:"root_size,omitempty"`

	// +kubebuilder:validation:Optional
	Spec *string `json:"spec,omitempty" tf:"spec,omitempty"`

	// +kubebuilder:validation:Optional
	StorageType *float64 `json:"storageType,omitempty" tf:"storage_type,omitempty"`
}

type EmrClusterInitParameters struct {

	// 0 means turn off automatic renewal, 1 means turn on automatic renewal. Default is 0.
	// 0 means turn off automatic renewal, 1 means turn on automatic renewal. Default is 0.
	AutoRenew *float64 `json:"autoRenew,omitempty" tf:"auto_renew,omitempty"`

	// It will be deprecated in later versions. Display strategy of EMR instance.
	// Display strategy of EMR instance.
	DisplayStrategy *string `json:"displayStrategy,omitempty" tf:"display_strategy,omitempty"`

	// Access the external file system.
	// Access the external file system.
	ExtendFsField *string `json:"extendFsField,omitempty" tf:"extend_fs_field,omitempty"`

	// Name of the instance, which can contain 6 to 36 English letters, Chinese characters, digits, dashes(-), or underscores(_).
	// Name of the instance, which can contain 6 to 36 English letters, Chinese characters, digits, dashes(-), or underscores(_).
	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name,omitempty"`

	// Whether to enable the cluster Master node public network. Value range:
	// - NEED_MASTER_WAN: Indicates that the cluster Master node public network is enabled.
	// - NOT_NEED_MASTER_WAN: Indicates that it is not turned on.
	// By default, the cluster Master node internet is enabled.
	// Whether to enable the cluster Master node public network. Value range:
	// - NEED_MASTER_WAN: Indicates that the cluster Master node public network is enabled.
	// - NOT_NEED_MASTER_WAN: Indicates that it is not turned on.
	// By default, the cluster Master node internet is enabled.
	NeedMasterWan *string `json:"needMasterWan,omitempty" tf:"need_master_wan,omitempty"`

	// The pay mode of instance. 0 represent POSTPAID_BY_HOUR, 1 represent PREPAID.
	// The pay mode of instance. 0 represent POSTPAID_BY_HOUR, 1 represent PREPAID.
	PayMode *float64 `json:"payMode,omitempty" tf:"pay_mode,omitempty"`

	// It will be deprecated in later versions. Use placement_info instead. The location of the instance.
	// The location of the instance.
	// +mapType=granular
	Placement map[string]*string `json:"placement,omitempty" tf:"placement,omitempty"`

	// The location of the instance.
	// The location of the instance.
	PlacementInfo []PlacementInfoInitParameters `json:"placementInfo,omitempty" tf:"placement_info,omitempty"`

	// Product ID. Different products ID represents different EMR product versions. Value range:
	// Product ID. Different products ID represents different EMR product versions. Value range:
	// - 16: represents EMR-V2.3.0
	// - 20: indicates EMR-V2.5.0
	// - 25: represents EMR-V3.1.0
	// - 27: represents KAFKA-V1.0.0
	// - 30: indicates EMR-V2.6.0
	// - 33: represents EMR-V3.2.1
	// - 34: stands for EMR-V3.3.0
	// - 36: represents STARROCKS-V1.0.0
	// - 37: indicates EMR-V3.4.0
	// - 38: represents EMR-V2.7.0
	// - 39: stands for STARROCKS-V1.1.0
	// - 41: represents DRUID-V1.1.0.
	ProductID *float64 `json:"productId,omitempty" tf:"product_id,omitempty"`

	// Resource specification of EMR instance.
	// Resource specification of EMR instance.
	ResourceSpec []ResourceSpecInitParameters `json:"resourceSpec,omitempty" tf:"resource_spec,omitempty"`

	// The ID of the security group to which the instance belongs, in the form of sg-xxxxxxxx.
	// The ID of the security group to which the instance belongs, in the form of sg-xxxxxxxx.
	SgID *string `json:"sgId,omitempty" tf:"sg_id,omitempty"`

	// The softwares of a EMR instance.
	// The softwares of a EMR instance.
	// +listType=set
	Softwares []*string `json:"softwares,omitempty" tf:"softwares,omitempty"`

	// The flag whether the instance support high availability.(0=>not support, 1=>support).
	// The flag whether the instance support high availability.(0=>not support, 1=>support).
	SupportHa *float64 `json:"supportHa,omitempty" tf:"support_ha,omitempty"`

	// Tag description list.
	// Tag description list.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The length of time the instance was purchased. Use with TimeUnit.When TimeUnit is s, the parameter can only be filled in at 3600, representing a metered instance.
	// When TimeUnit is m, the number filled in by this parameter indicates the length of purchase of the monthly instance of the package year, such as 1 for one month of purchase.
	// The length of time the instance was purchased. Use with TimeUnit.When TimeUnit is s, the parameter can only be filled in at 3600, representing a metered instance.
	// When TimeUnit is m, the number filled in by this parameter indicates the length of purchase of the monthly instance of the package year, such as 1 for one month of purchase.
	TimeSpan *float64 `json:"timeSpan,omitempty" tf:"time_span,omitempty"`

	// The unit of time in which the instance was purchased. When PayMode is 0, TimeUnit can only take values of s(second). When PayMode is 1, TimeUnit can only take the value m(month).
	// The unit of time in which the instance was purchased. When PayMode is 0, TimeUnit can only take values of s(second). When PayMode is 1, TimeUnit can only take the value m(month).
	TimeUnit *string `json:"timeUnit,omitempty" tf:"time_unit,omitempty"`

	// The private net config of EMR instance.
	// The private net config of EMR instance.
	// +mapType=granular
	VPCSettings map[string]*string `json:"vpcSettings,omitempty" tf:"vpc_settings,omitempty"`
}

type EmrClusterObservation struct {

	// 0 means turn off automatic renewal, 1 means turn on automatic renewal. Default is 0.
	// 0 means turn off automatic renewal, 1 means turn on automatic renewal. Default is 0.
	AutoRenew *float64 `json:"autoRenew,omitempty" tf:"auto_renew,omitempty"`

	// It will be deprecated in later versions. Display strategy of EMR instance.
	// Display strategy of EMR instance.
	DisplayStrategy *string `json:"displayStrategy,omitempty" tf:"display_strategy,omitempty"`

	// Access the external file system.
	// Access the external file system.
	ExtendFsField *string `json:"extendFsField,omitempty" tf:"extend_fs_field,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Created EMR instance id.
	// Created EMR instance id.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Name of the instance, which can contain 6 to 36 English letters, Chinese characters, digits, dashes(-), or underscores(_).
	// Name of the instance, which can contain 6 to 36 English letters, Chinese characters, digits, dashes(-), or underscores(_).
	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name,omitempty"`

	// Whether to enable the cluster Master node public network. Value range:
	// - NEED_MASTER_WAN: Indicates that the cluster Master node public network is enabled.
	// - NOT_NEED_MASTER_WAN: Indicates that it is not turned on.
	// By default, the cluster Master node internet is enabled.
	// Whether to enable the cluster Master node public network. Value range:
	// - NEED_MASTER_WAN: Indicates that the cluster Master node public network is enabled.
	// - NOT_NEED_MASTER_WAN: Indicates that it is not turned on.
	// By default, the cluster Master node internet is enabled.
	NeedMasterWan *string `json:"needMasterWan,omitempty" tf:"need_master_wan,omitempty"`

	// The pay mode of instance. 0 represent POSTPAID_BY_HOUR, 1 represent PREPAID.
	// The pay mode of instance. 0 represent POSTPAID_BY_HOUR, 1 represent PREPAID.
	PayMode *float64 `json:"payMode,omitempty" tf:"pay_mode,omitempty"`

	// It will be deprecated in later versions. Use placement_info instead. The location of the instance.
	// The location of the instance.
	// +mapType=granular
	Placement map[string]*string `json:"placement,omitempty" tf:"placement,omitempty"`

	// The location of the instance.
	// The location of the instance.
	PlacementInfo []PlacementInfoObservation `json:"placementInfo,omitempty" tf:"placement_info,omitempty"`

	// Product ID. Different products ID represents different EMR product versions. Value range:
	// Product ID. Different products ID represents different EMR product versions. Value range:
	// - 16: represents EMR-V2.3.0
	// - 20: indicates EMR-V2.5.0
	// - 25: represents EMR-V3.1.0
	// - 27: represents KAFKA-V1.0.0
	// - 30: indicates EMR-V2.6.0
	// - 33: represents EMR-V3.2.1
	// - 34: stands for EMR-V3.3.0
	// - 36: represents STARROCKS-V1.0.0
	// - 37: indicates EMR-V3.4.0
	// - 38: represents EMR-V2.7.0
	// - 39: stands for STARROCKS-V1.1.0
	// - 41: represents DRUID-V1.1.0.
	ProductID *float64 `json:"productId,omitempty" tf:"product_id,omitempty"`

	// Resource specification of EMR instance.
	// Resource specification of EMR instance.
	ResourceSpec []ResourceSpecObservation `json:"resourceSpec,omitempty" tf:"resource_spec,omitempty"`

	// The ID of the security group to which the instance belongs, in the form of sg-xxxxxxxx.
	// The ID of the security group to which the instance belongs, in the form of sg-xxxxxxxx.
	SgID *string `json:"sgId,omitempty" tf:"sg_id,omitempty"`

	// The softwares of a EMR instance.
	// The softwares of a EMR instance.
	// +listType=set
	Softwares []*string `json:"softwares,omitempty" tf:"softwares,omitempty"`

	// The flag whether the instance support high availability.(0=>not support, 1=>support).
	// The flag whether the instance support high availability.(0=>not support, 1=>support).
	SupportHa *float64 `json:"supportHa,omitempty" tf:"support_ha,omitempty"`

	// Tag description list.
	// Tag description list.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The length of time the instance was purchased. Use with TimeUnit.When TimeUnit is s, the parameter can only be filled in at 3600, representing a metered instance.
	// When TimeUnit is m, the number filled in by this parameter indicates the length of purchase of the monthly instance of the package year, such as 1 for one month of purchase.
	// The length of time the instance was purchased. Use with TimeUnit.When TimeUnit is s, the parameter can only be filled in at 3600, representing a metered instance.
	// When TimeUnit is m, the number filled in by this parameter indicates the length of purchase of the monthly instance of the package year, such as 1 for one month of purchase.
	TimeSpan *float64 `json:"timeSpan,omitempty" tf:"time_span,omitempty"`

	// The unit of time in which the instance was purchased. When PayMode is 0, TimeUnit can only take values of s(second). When PayMode is 1, TimeUnit can only take the value m(month).
	// The unit of time in which the instance was purchased. When PayMode is 0, TimeUnit can only take values of s(second). When PayMode is 1, TimeUnit can only take the value m(month).
	TimeUnit *string `json:"timeUnit,omitempty" tf:"time_unit,omitempty"`

	// The private net config of EMR instance.
	// The private net config of EMR instance.
	// +mapType=granular
	VPCSettings map[string]*string `json:"vpcSettings,omitempty" tf:"vpc_settings,omitempty"`
}

type EmrClusterParameters struct {

	// 0 means turn off automatic renewal, 1 means turn on automatic renewal. Default is 0.
	// 0 means turn off automatic renewal, 1 means turn on automatic renewal. Default is 0.
	// +kubebuilder:validation:Optional
	AutoRenew *float64 `json:"autoRenew,omitempty" tf:"auto_renew,omitempty"`

	// It will be deprecated in later versions. Display strategy of EMR instance.
	// Display strategy of EMR instance.
	// +kubebuilder:validation:Optional
	DisplayStrategy *string `json:"displayStrategy,omitempty" tf:"display_strategy,omitempty"`

	// Access the external file system.
	// Access the external file system.
	// +kubebuilder:validation:Optional
	ExtendFsField *string `json:"extendFsField,omitempty" tf:"extend_fs_field,omitempty"`

	// Name of the instance, which can contain 6 to 36 English letters, Chinese characters, digits, dashes(-), or underscores(_).
	// Name of the instance, which can contain 6 to 36 English letters, Chinese characters, digits, dashes(-), or underscores(_).
	// +kubebuilder:validation:Optional
	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name,omitempty"`

	// Instance login settings. There are two optional fields:- password: Instance login password: 8-16 characters, including uppercase letters, lowercase letters, numbers and special characters. Special symbols only support! @% ^ *. The first bit of the password cannot be a special character;- public_key_id: Public key id. After the key is associated, the instance can be accessed through the corresponding private key.
	// Instance login settings. There are two optional fields:- password: Instance login password: 8-16 characters, including uppercase letters, lowercase letters, numbers and special characters. Special symbols only support! @% ^ *. The first bit of the password cannot be a special character;- public_key_id: Public key id. After the key is associated, the instance can be accessed through the corresponding private key.
	// +kubebuilder:validation:Optional
	LoginSettingsSecretRef *v1.SecretReference `json:"loginSettingsSecretRef,omitempty" tf:"-"`

	// Whether to enable the cluster Master node public network. Value range:
	// - NEED_MASTER_WAN: Indicates that the cluster Master node public network is enabled.
	// - NOT_NEED_MASTER_WAN: Indicates that it is not turned on.
	// By default, the cluster Master node internet is enabled.
	// Whether to enable the cluster Master node public network. Value range:
	// - NEED_MASTER_WAN: Indicates that the cluster Master node public network is enabled.
	// - NOT_NEED_MASTER_WAN: Indicates that it is not turned on.
	// By default, the cluster Master node internet is enabled.
	// +kubebuilder:validation:Optional
	NeedMasterWan *string `json:"needMasterWan,omitempty" tf:"need_master_wan,omitempty"`

	// The pay mode of instance. 0 represent POSTPAID_BY_HOUR, 1 represent PREPAID.
	// The pay mode of instance. 0 represent POSTPAID_BY_HOUR, 1 represent PREPAID.
	// +kubebuilder:validation:Optional
	PayMode *float64 `json:"payMode,omitempty" tf:"pay_mode,omitempty"`

	// It will be deprecated in later versions. Use placement_info instead. The location of the instance.
	// The location of the instance.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Placement map[string]*string `json:"placement,omitempty" tf:"placement,omitempty"`

	// The location of the instance.
	// The location of the instance.
	// +kubebuilder:validation:Optional
	PlacementInfo []PlacementInfoParameters `json:"placementInfo,omitempty" tf:"placement_info,omitempty"`

	// Product ID. Different products ID represents different EMR product versions. Value range:
	// Product ID. Different products ID represents different EMR product versions. Value range:
	// - 16: represents EMR-V2.3.0
	// - 20: indicates EMR-V2.5.0
	// - 25: represents EMR-V3.1.0
	// - 27: represents KAFKA-V1.0.0
	// - 30: indicates EMR-V2.6.0
	// - 33: represents EMR-V3.2.1
	// - 34: stands for EMR-V3.3.0
	// - 36: represents STARROCKS-V1.0.0
	// - 37: indicates EMR-V3.4.0
	// - 38: represents EMR-V2.7.0
	// - 39: stands for STARROCKS-V1.1.0
	// - 41: represents DRUID-V1.1.0.
	// +kubebuilder:validation:Optional
	ProductID *float64 `json:"productId,omitempty" tf:"product_id,omitempty"`

	// Resource specification of EMR instance.
	// Resource specification of EMR instance.
	// +kubebuilder:validation:Optional
	ResourceSpec []ResourceSpecParameters `json:"resourceSpec,omitempty" tf:"resource_spec,omitempty"`

	// The ID of the security group to which the instance belongs, in the form of sg-xxxxxxxx.
	// The ID of the security group to which the instance belongs, in the form of sg-xxxxxxxx.
	// +kubebuilder:validation:Optional
	SgID *string `json:"sgId,omitempty" tf:"sg_id,omitempty"`

	// The softwares of a EMR instance.
	// The softwares of a EMR instance.
	// +kubebuilder:validation:Optional
	// +listType=set
	Softwares []*string `json:"softwares,omitempty" tf:"softwares,omitempty"`

	// The flag whether the instance support high availability.(0=>not support, 1=>support).
	// The flag whether the instance support high availability.(0=>not support, 1=>support).
	// +kubebuilder:validation:Optional
	SupportHa *float64 `json:"supportHa,omitempty" tf:"support_ha,omitempty"`

	// Tag description list.
	// Tag description list.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The length of time the instance was purchased. Use with TimeUnit.When TimeUnit is s, the parameter can only be filled in at 3600, representing a metered instance.
	// When TimeUnit is m, the number filled in by this parameter indicates the length of purchase of the monthly instance of the package year, such as 1 for one month of purchase.
	// The length of time the instance was purchased. Use with TimeUnit.When TimeUnit is s, the parameter can only be filled in at 3600, representing a metered instance.
	// When TimeUnit is m, the number filled in by this parameter indicates the length of purchase of the monthly instance of the package year, such as 1 for one month of purchase.
	// +kubebuilder:validation:Optional
	TimeSpan *float64 `json:"timeSpan,omitempty" tf:"time_span,omitempty"`

	// The unit of time in which the instance was purchased. When PayMode is 0, TimeUnit can only take values of s(second). When PayMode is 1, TimeUnit can only take the value m(month).
	// The unit of time in which the instance was purchased. When PayMode is 0, TimeUnit can only take values of s(second). When PayMode is 1, TimeUnit can only take the value m(month).
	// +kubebuilder:validation:Optional
	TimeUnit *string `json:"timeUnit,omitempty" tf:"time_unit,omitempty"`

	// The private net config of EMR instance.
	// The private net config of EMR instance.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	VPCSettings map[string]*string `json:"vpcSettings,omitempty" tf:"vpc_settings,omitempty"`
}

type MasterResourceSpecInitParameters struct {
	CPU *float64 `json:"cpu,omitempty" tf:"cpu,omitempty"`

	DiskSize *float64 `json:"diskSize,omitempty" tf:"disk_size,omitempty"`

	DiskType *string `json:"diskType,omitempty" tf:"disk_type,omitempty"`

	MemSize *float64 `json:"memSize,omitempty" tf:"mem_size,omitempty"`

	RootSize *float64 `json:"rootSize,omitempty" tf:"root_size,omitempty"`

	Spec *string `json:"spec,omitempty" tf:"spec,omitempty"`

	StorageType *float64 `json:"storageType,omitempty" tf:"storage_type,omitempty"`
}

type MasterResourceSpecObservation struct {
	CPU *float64 `json:"cpu,omitempty" tf:"cpu,omitempty"`

	DiskSize *float64 `json:"diskSize,omitempty" tf:"disk_size,omitempty"`

	DiskType *string `json:"diskType,omitempty" tf:"disk_type,omitempty"`

	MemSize *float64 `json:"memSize,omitempty" tf:"mem_size,omitempty"`

	RootSize *float64 `json:"rootSize,omitempty" tf:"root_size,omitempty"`

	Spec *string `json:"spec,omitempty" tf:"spec,omitempty"`

	StorageType *float64 `json:"storageType,omitempty" tf:"storage_type,omitempty"`
}

type MasterResourceSpecParameters struct {

	// +kubebuilder:validation:Optional
	CPU *float64 `json:"cpu,omitempty" tf:"cpu,omitempty"`

	// +kubebuilder:validation:Optional
	DiskSize *float64 `json:"diskSize,omitempty" tf:"disk_size,omitempty"`

	// +kubebuilder:validation:Optional
	DiskType *string `json:"diskType,omitempty" tf:"disk_type,omitempty"`

	// +kubebuilder:validation:Optional
	MemSize *float64 `json:"memSize,omitempty" tf:"mem_size,omitempty"`

	// +kubebuilder:validation:Optional
	RootSize *float64 `json:"rootSize,omitempty" tf:"root_size,omitempty"`

	// +kubebuilder:validation:Optional
	Spec *string `json:"spec,omitempty" tf:"spec,omitempty"`

	// +kubebuilder:validation:Optional
	StorageType *float64 `json:"storageType,omitempty" tf:"storage_type,omitempty"`
}

type PlacementInfoInitParameters struct {

	// Project id.
	// Project id.
	ProjectID *float64 `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Zone.
	// Zone.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type PlacementInfoObservation struct {

	// Project id.
	// Project id.
	ProjectID *float64 `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Zone.
	// Zone.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type PlacementInfoParameters struct {

	// Project id.
	// Project id.
	// +kubebuilder:validation:Optional
	ProjectID *float64 `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Zone.
	// Zone.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone" tf:"zone,omitempty"`
}

type ResourceSpecInitParameters struct {

	// The number of common node.
	// The number of common node.
	CommonCount *float64 `json:"commonCount,omitempty" tf:"common_count,omitempty"`

	CommonResourceSpec []CommonResourceSpecInitParameters `json:"commonResourceSpec,omitempty" tf:"common_resource_spec,omitempty"`

	// The number of core node.
	// The number of core node.
	CoreCount *float64 `json:"coreCount,omitempty" tf:"core_count,omitempty"`

	CoreResourceSpec []CoreResourceSpecInitParameters `json:"coreResourceSpec,omitempty" tf:"core_resource_spec,omitempty"`

	// The number of master node.
	// The number of master node.
	MasterCount *float64 `json:"masterCount,omitempty" tf:"master_count,omitempty"`

	MasterResourceSpec []MasterResourceSpecInitParameters `json:"masterResourceSpec,omitempty" tf:"master_resource_spec,omitempty"`

	// The number of core node.
	// The number of core node.
	TaskCount *float64 `json:"taskCount,omitempty" tf:"task_count,omitempty"`

	TaskResourceSpec []TaskResourceSpecInitParameters `json:"taskResourceSpec,omitempty" tf:"task_resource_spec,omitempty"`
}

type ResourceSpecObservation struct {

	// The number of common node.
	// The number of common node.
	CommonCount *float64 `json:"commonCount,omitempty" tf:"common_count,omitempty"`

	CommonResourceSpec []CommonResourceSpecObservation `json:"commonResourceSpec,omitempty" tf:"common_resource_spec,omitempty"`

	// The number of core node.
	// The number of core node.
	CoreCount *float64 `json:"coreCount,omitempty" tf:"core_count,omitempty"`

	CoreResourceSpec []CoreResourceSpecObservation `json:"coreResourceSpec,omitempty" tf:"core_resource_spec,omitempty"`

	// The number of master node.
	// The number of master node.
	MasterCount *float64 `json:"masterCount,omitempty" tf:"master_count,omitempty"`

	MasterResourceSpec []MasterResourceSpecObservation `json:"masterResourceSpec,omitempty" tf:"master_resource_spec,omitempty"`

	// The number of core node.
	// The number of core node.
	TaskCount *float64 `json:"taskCount,omitempty" tf:"task_count,omitempty"`

	TaskResourceSpec []TaskResourceSpecObservation `json:"taskResourceSpec,omitempty" tf:"task_resource_spec,omitempty"`
}

type ResourceSpecParameters struct {

	// The number of common node.
	// The number of common node.
	// +kubebuilder:validation:Optional
	CommonCount *float64 `json:"commonCount,omitempty" tf:"common_count,omitempty"`

	// +kubebuilder:validation:Optional
	CommonResourceSpec []CommonResourceSpecParameters `json:"commonResourceSpec,omitempty" tf:"common_resource_spec,omitempty"`

	// The number of core node.
	// The number of core node.
	// +kubebuilder:validation:Optional
	CoreCount *float64 `json:"coreCount,omitempty" tf:"core_count,omitempty"`

	// +kubebuilder:validation:Optional
	CoreResourceSpec []CoreResourceSpecParameters `json:"coreResourceSpec,omitempty" tf:"core_resource_spec,omitempty"`

	// The number of master node.
	// The number of master node.
	// +kubebuilder:validation:Optional
	MasterCount *float64 `json:"masterCount,omitempty" tf:"master_count,omitempty"`

	// +kubebuilder:validation:Optional
	MasterResourceSpec []MasterResourceSpecParameters `json:"masterResourceSpec,omitempty" tf:"master_resource_spec,omitempty"`

	// The number of core node.
	// The number of core node.
	// +kubebuilder:validation:Optional
	TaskCount *float64 `json:"taskCount,omitempty" tf:"task_count,omitempty"`

	// +kubebuilder:validation:Optional
	TaskResourceSpec []TaskResourceSpecParameters `json:"taskResourceSpec,omitempty" tf:"task_resource_spec,omitempty"`
}

type TaskResourceSpecInitParameters struct {
	CPU *float64 `json:"cpu,omitempty" tf:"cpu,omitempty"`

	DiskSize *float64 `json:"diskSize,omitempty" tf:"disk_size,omitempty"`

	DiskType *string `json:"diskType,omitempty" tf:"disk_type,omitempty"`

	MemSize *float64 `json:"memSize,omitempty" tf:"mem_size,omitempty"`

	RootSize *float64 `json:"rootSize,omitempty" tf:"root_size,omitempty"`

	Spec *string `json:"spec,omitempty" tf:"spec,omitempty"`

	StorageType *float64 `json:"storageType,omitempty" tf:"storage_type,omitempty"`
}

type TaskResourceSpecObservation struct {
	CPU *float64 `json:"cpu,omitempty" tf:"cpu,omitempty"`

	DiskSize *float64 `json:"diskSize,omitempty" tf:"disk_size,omitempty"`

	DiskType *string `json:"diskType,omitempty" tf:"disk_type,omitempty"`

	MemSize *float64 `json:"memSize,omitempty" tf:"mem_size,omitempty"`

	RootSize *float64 `json:"rootSize,omitempty" tf:"root_size,omitempty"`

	Spec *string `json:"spec,omitempty" tf:"spec,omitempty"`

	StorageType *float64 `json:"storageType,omitempty" tf:"storage_type,omitempty"`
}

type TaskResourceSpecParameters struct {

	// +kubebuilder:validation:Optional
	CPU *float64 `json:"cpu,omitempty" tf:"cpu,omitempty"`

	// +kubebuilder:validation:Optional
	DiskSize *float64 `json:"diskSize,omitempty" tf:"disk_size,omitempty"`

	// +kubebuilder:validation:Optional
	DiskType *string `json:"diskType,omitempty" tf:"disk_type,omitempty"`

	// +kubebuilder:validation:Optional
	MemSize *float64 `json:"memSize,omitempty" tf:"mem_size,omitempty"`

	// +kubebuilder:validation:Optional
	RootSize *float64 `json:"rootSize,omitempty" tf:"root_size,omitempty"`

	// +kubebuilder:validation:Optional
	Spec *string `json:"spec,omitempty" tf:"spec,omitempty"`

	// +kubebuilder:validation:Optional
	StorageType *float64 `json:"storageType,omitempty" tf:"storage_type,omitempty"`
}

// EmrClusterSpec defines the desired state of EmrCluster
type EmrClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EmrClusterParameters `json:"forProvider"`
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
	InitProvider EmrClusterInitParameters `json:"initProvider,omitempty"`
}

// EmrClusterStatus defines the observed state of EmrCluster.
type EmrClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EmrClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// EmrCluster is the Schema for the EmrClusters API. Provide a resource to create a emr cluster.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type EmrCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.instanceName) || (has(self.initProvider) && has(self.initProvider.instanceName))",message="spec.forProvider.instanceName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.payMode) || (has(self.initProvider) && has(self.initProvider.payMode))",message="spec.forProvider.payMode is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.productId) || (has(self.initProvider) && has(self.initProvider.productId))",message="spec.forProvider.productId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.softwares) || (has(self.initProvider) && has(self.initProvider.softwares))",message="spec.forProvider.softwares is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.supportHa) || (has(self.initProvider) && has(self.initProvider.supportHa))",message="spec.forProvider.supportHa is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.vpcSettings) || (has(self.initProvider) && has(self.initProvider.vpcSettings))",message="spec.forProvider.vpcSettings is a required parameter"
	Spec   EmrClusterSpec   `json:"spec"`
	Status EmrClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EmrClusterList contains a list of EmrClusters
type EmrClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EmrCluster `json:"items"`
}

// Repository type metadata.
var (
	EmrCluster_Kind             = "EmrCluster"
	EmrCluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EmrCluster_Kind}.String()
	EmrCluster_KindAPIVersion   = EmrCluster_Kind + "." + CRDGroupVersion.String()
	EmrCluster_GroupVersionKind = CRDGroupVersion.WithKind(EmrCluster_Kind)
)

func init() {
	SchemeBuilder.Register(&EmrCluster{}, &EmrClusterList{})
}
