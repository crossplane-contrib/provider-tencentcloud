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

type DataDisksObservation struct {
}

type DataDisksParameters struct {

	// Data disk ID used to initialize the data disk. When data disk type is `LOCAL_BASIC` and `LOCAL_SSD`, disk id is not supported.
	// +kubebuilder:validation:Optional
	DataDiskID *string `json:"dataDiskId,omitempty" tf:"data_disk_id,omitempty"`

	// Size of the data disk, and unit is GB.
	// +kubebuilder:validation:Required
	DataDiskSize *float64 `json:"dataDiskSize" tf:"data_disk_size,omitempty"`

	// Snapshot ID of the data disk. The selected data disk snapshot size must be smaller than the data disk size.
	// +kubebuilder:validation:Optional
	DataDiskSnapshotID *string `json:"dataDiskSnapshotId,omitempty" tf:"data_disk_snapshot_id,omitempty"`

	// Data disk type. For more information about limits on different data disk types, see [Storage Overview](https://intl.cloud.tencent.com/document/product/213/4952). Valid values: LOCAL_BASIC: local disk, LOCAL_SSD: local SSD disk, LOCAL_NVME: local NVME disk, specified in the InstanceType, LOCAL_PRO: local HDD disk, specified in the InstanceType, CLOUD_BASIC: HDD cloud disk, CLOUD_PREMIUM: Premium Cloud Storage, CLOUD_SSD: SSD, CLOUD_HSSD: Enhanced SSD, CLOUD_TSSD: Tremendous SSD, CLOUD_BSSD: Balanced SSD.
	// +kubebuilder:validation:Required
	DataDiskType *string `json:"dataDiskType" tf:"data_disk_type,omitempty"`

	// Decides whether the disk is deleted with instance(only applied to `CLOUD_BASIC`, `CLOUD_SSD` and `CLOUD_PREMIUM` disk with `POSTPAID_BY_HOUR` instance), default is true.
	// +kubebuilder:validation:Optional
	DeleteWithInstance *bool `json:"deleteWithInstance,omitempty" tf:"delete_with_instance,omitempty"`

	// Decides whether the disk is encrypted. Default is `false`.
	// +kubebuilder:validation:Optional
	Encrypt *bool `json:"encrypt,omitempty" tf:"encrypt,omitempty"`

	// Add extra performance to the data disk. Only works when disk type is `CLOUD_TSSD` or `CLOUD_HSSD`.
	// +kubebuilder:validation:Optional
	ThroughputPerformance *float64 `json:"throughputPerformance,omitempty" tf:"throughput_performance,omitempty"`
}

type InstanceObservation struct {
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	ExpiredTime *string `json:"expiredTime,omitempty" tf:"expired_time,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	InstanceStatus *string `json:"instanceStatus,omitempty" tf:"instance_status,omitempty"`

	PublicIP *string `json:"publicIp,omitempty" tf:"public_ip,omitempty"`

	SystemDiskID *string `json:"systemDiskId,omitempty" tf:"system_disk_id,omitempty"`
}

type InstanceParameters struct {

	// Associate a public IP address with an instance in a VPC or Classic. Boolean value, Default is false.
	// +kubebuilder:validation:Optional
	AllocatePublicIP *bool `json:"allocatePublicIp,omitempty" tf:"allocate_public_ip,omitempty"`

	// The available zone for the CVM instance.
	// +kubebuilder:validation:Required
	AvailabilityZone *string `json:"availabilityZone" tf:"availability_zone,omitempty"`

	// bandwidth package id. if user is standard user, then the bandwidth_package_id is needed, or default has bandwidth_package_id.
	// +kubebuilder:validation:Optional
	BandwidthPackageID *string `json:"bandwidthPackageId,omitempty" tf:"bandwidth_package_id,omitempty"`

	// CAM role name authorized to access.
	// +kubebuilder:validation:Optional
	CamRoleName *string `json:"camRoleName,omitempty" tf:"cam_role_name,omitempty"`

	// Id of cdh instance. Note: it only works when instance_charge_type is set to `CDHPAID`.
	// +kubebuilder:validation:Optional
	CdhHostID *string `json:"cdhHostId,omitempty" tf:"cdh_host_id,omitempty"`

	// Type of instance created on cdh, the value of this parameter is in the format of CDH_XCXG based on the number of CPU cores and memory capacity. Note: it only works when instance_charge_type is set to `CDHPAID`.
	// +kubebuilder:validation:Optional
	CdhInstanceType *string `json:"cdhInstanceType,omitempty" tf:"cdh_instance_type,omitempty"`

	// Settings for data disks.
	// +kubebuilder:validation:Optional
	DataDisks []DataDisksParameters `json:"dataDisks,omitempty" tf:"data_disks,omitempty"`

	// Whether the termination protection is enabled. Default is `false`. If set true, which means that this instance can not be deleted by an API action.
	// +kubebuilder:validation:Optional
	DisableAPITermination *bool `json:"disableApiTermination,omitempty" tf:"disable_api_termination,omitempty"`

	// Disable enhance service for monitor, it is enabled by default. When this options is set, monitor agent won't be installed. Modifying will cause the instance reset.
	// +kubebuilder:validation:Optional
	DisableMonitorService *bool `json:"disableMonitorService,omitempty" tf:"disable_monitor_service,omitempty"`

	// Disable enhance service for security, it is enabled by default. When this options is set, security agent won't be installed. Modifying will cause the instance reset.
	// +kubebuilder:validation:Optional
	DisableSecurityService *bool `json:"disableSecurityService,omitempty" tf:"disable_security_service,omitempty"`

	// Indicate whether to force delete the instance. Default is `false`. If set true, the instance will be permanently deleted instead of being moved into the recycle bin. Note: only works for `PREPAID` instance.
	// +kubebuilder:validation:Optional
	ForceDelete *bool `json:"forceDelete,omitempty" tf:"force_delete,omitempty"`

	// The hostname of the instance. Windows instance: The name should be a combination of 2 to 15 characters comprised of letters (case insensitive), numbers, and hyphens (-). Period (.) is not supported, and the name cannot be a string of pure numbers. Other types (such as Linux) of instances: The name should be a combination of 2 to 60 characters, supporting multiple periods (.). The piece between two periods is composed of letters (case insensitive), numbers, and hyphens (-). Modifying will cause the instance reset.
	// +kubebuilder:validation:Optional
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// The image to use for the instance. Changing `image_id` will cause the instance reset.
	// +kubebuilder:validation:Required
	ImageID *string `json:"imageId" tf:"image_id,omitempty"`

	// The charge type of instance. Valid values are `PREPAID`, `POSTPAID_BY_HOUR`, `SPOTPAID` and `CDHPAID`. The default is `POSTPAID_BY_HOUR`. Note: TencentCloud International only supports `POSTPAID_BY_HOUR` and `CDHPAID`. `PREPAID` instance may not allow to delete before expired. `SPOTPAID` instance must set `spot_instance_type` and `spot_max_price` at the same time. `CDHPAID` instance must set `cdh_instance_type` and `cdh_host_id`.
	// +kubebuilder:validation:Optional
	InstanceChargeType *string `json:"instanceChargeType,omitempty" tf:"instance_charge_type,omitempty"`

	// The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when instance_charge_type is set to `PREPAID`. Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
	// +kubebuilder:validation:Optional
	InstanceChargeTypePrepaidPeriod *float64 `json:"instanceChargeTypePrepaidPeriod,omitempty" tf:"instance_charge_type_prepaid_period,omitempty"`

	// Auto renewal flag. Valid values: `NOTIFY_AND_AUTO_RENEW`: notify upon expiration and renew automatically, `NOTIFY_AND_MANUAL_RENEW`: notify upon expiration but do not renew automatically, `DISABLE_NOTIFY_AND_MANUAL_RENEW`: neither notify upon expiration nor renew automatically. Default value: `NOTIFY_AND_MANUAL_RENEW`. If this parameter is specified as `NOTIFY_AND_AUTO_RENEW`, the instance will be automatically renewed on a monthly basis if the account balance is sufficient. NOTE: it only works when instance_charge_type is set to `PREPAID`.
	// +kubebuilder:validation:Optional
	InstanceChargeTypePrepaidRenewFlag *string `json:"instanceChargeTypePrepaidRenewFlag,omitempty" tf:"instance_charge_type_prepaid_renew_flag,omitempty"`

	// The number of instances to be purchased. Value range:[1,100]; default value: 1.
	// +kubebuilder:validation:Optional
	InstanceCount *float64 `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`

	// The name of the instance. The max length of instance_name is 60, and default value is `Terraform-CVM-Instance`.
	// +kubebuilder:validation:Optional
	InstanceName *string `json:"instanceName,omitempty" tf:"instance_name,omitempty"`

	// The type of the instance.
	// +kubebuilder:validation:Optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// Internet charge type of the instance, Valid values are `BANDWIDTH_PREPAID`, `TRAFFIC_POSTPAID_BY_HOUR`, `BANDWIDTH_POSTPAID_BY_HOUR` and `BANDWIDTH_PACKAGE`. If not set, internet charge type are consistent with the cvm charge type by default. This value takes NO Effect when changing and does not need to be set when `allocate_public_ip` is false.
	// +kubebuilder:validation:Optional
	InternetChargeType *string `json:"internetChargeType,omitempty" tf:"internet_charge_type,omitempty"`

	// Maximum outgoing bandwidth to the public network, measured in Mbps (Mega bits per second). This value does not need to be set when `allocate_public_ip` is false.
	// +kubebuilder:validation:Optional
	InternetMaxBandwidthOut *float64 `json:"internetMaxBandwidthOut,omitempty" tf:"internet_max_bandwidth_out,omitempty"`

	// Whether to keep image login or not, default is `false`. When the image type is private or shared or imported, this parameter can be set `true`. Modifying will cause the instance reset.
	// +kubebuilder:validation:Optional
	KeepImageLogin *bool `json:"keepImageLogin,omitempty" tf:"keep_image_login,omitempty"`

	// The key pair to use for the instance, it looks like `skey-16jig7tx`. Modifying will cause the instance reset.
	// +kubebuilder:validation:Optional
	KeyIds []*string `json:"keyIds,omitempty" tf:"key_ids,omitempty"`

	// The key pair to use for the instance, it looks like `skey-16jig7tx`. Modifying will cause the instance reset.
	// +kubebuilder:validation:Optional
	KeyName *string `json:"keyName,omitempty" tf:"key_name,omitempty"`

	// A list of orderly security group IDs to associate with.
	// +kubebuilder:validation:Optional
	OrderlySecurityGroups []*string `json:"orderlySecurityGroups,omitempty" tf:"orderly_security_groups,omitempty"`

	// Password for the instance. In order for the new password to take effect, the instance will be restarted after the password change. Modifying will cause the instance reset.
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// The ID of a placement group.
	// +kubebuilder:validation:Optional
	PlacementGroupID *string `json:"placementGroupId,omitempty" tf:"placement_group_id,omitempty"`

	// The private IP to be assigned to this instance, must be in the provided subnet and available.
	// +kubebuilder:validation:Optional
	PrivateIP *string `json:"privateIp,omitempty" tf:"private_ip,omitempty"`

	// The project the instance belongs to, default to 0.
	// +kubebuilder:validation:Optional
	ProjectID *float64 `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Set instance to running or stop. Default value is true, the instance will shutdown when this flag is false.
	// +kubebuilder:validation:Optional
	RunningFlag *bool `json:"runningFlag,omitempty" tf:"running_flag,omitempty"`

	// A list of security group IDs to associate with.
	// +kubebuilder:validation:Optional
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// Type of spot instance, only support `ONE-TIME` now. Note: it only works when instance_charge_type is set to `SPOTPAID`.
	// +kubebuilder:validation:Optional
	SpotInstanceType *string `json:"spotInstanceType,omitempty" tf:"spot_instance_type,omitempty"`

	// Max price of a spot instance, is the format of decimal string, for example "0.50". Note: it only works when instance_charge_type is set to `SPOTPAID`.
	// +kubebuilder:validation:Optional
	SpotMaxPrice *string `json:"spotMaxPrice,omitempty" tf:"spot_max_price,omitempty"`

	// Billing method of a pay-as-you-go instance after shutdown. Available values: `KEEP_CHARGING`,`STOP_CHARGING`. Default `KEEP_CHARGING`.
	// +kubebuilder:validation:Optional
	StoppedMode *string `json:"stoppedMode,omitempty" tf:"stopped_mode,omitempty"`

	// The ID of a VPC subnet. If you want to create instances in a VPC network, this parameter must be set.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// Size of the system disk. unit is GB, Default is 50GB. If modified, the instance may force stop.
	// +kubebuilder:validation:Optional
	SystemDiskSize *float64 `json:"systemDiskSize,omitempty" tf:"system_disk_size,omitempty"`

	// System disk type. For more information on limits of system disk types, see [Storage Overview](https://intl.cloud.tencent.com/document/product/213/4952). Valid values: `LOCAL_BASIC`: local disk, `LOCAL_SSD`: local SSD disk, `CLOUD_BASIC`: cloud disk, `CLOUD_SSD`: cloud SSD disk, `CLOUD_PREMIUM`: Premium Cloud Storage, `CLOUD_BSSD`: Basic SSD, `CLOUD_HSSD`: Enhanced SSD, `CLOUD_TSSD`: Tremendous SSD. NOTE: If modified, the instance may force stop.
	// +kubebuilder:validation:Optional
	SystemDiskType *string `json:"systemDiskType,omitempty" tf:"system_disk_type,omitempty"`

	// A mapping of tags to assign to the resource. For tag limits, please refer to [Use Limits](https://intl.cloud.tencent.com/document/product/651/13354).
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The user data to be injected into this instance. Must be base64 encoded and up to 16 KB.
	// +kubebuilder:validation:Optional
	UserData *string `json:"userData,omitempty" tf:"user_data,omitempty"`

	// The user data to be injected into this instance, in plain text. Conflicts with `user_data`. Up to 16 KB after base64 encoded.
	// +kubebuilder:validation:Optional
	UserDataRaw *string `json:"userDataRaw,omitempty" tf:"user_data_raw,omitempty"`

	// The ID of a VPC network. If you want to create instances in a VPC network, this parameter must be set.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcidRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcidSelector,omitempty" tf:"-"`
}

// InstanceSpec defines the desired state of Instance
type InstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceParameters `json:"forProvider"`
}

// InstanceStatus defines the observed state of Instance.
type InstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Instance is the Schema for the Instances API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceSpec   `json:"spec"`
	Status            InstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceList contains a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Instance `json:"items"`
}

// Repository type metadata.
var (
	Instance_Kind             = "Instance"
	Instance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Instance_Kind}.String()
	Instance_KindAPIVersion   = Instance_Kind + "." + CRDGroupVersion.String()
	Instance_GroupVersionKind = CRDGroupVersion.WithKind(Instance_Kind)
)

func init() {
	SchemeBuilder.Register(&Instance{}, &InstanceList{})
}
