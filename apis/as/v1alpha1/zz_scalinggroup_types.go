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

type ForwardBalancerIdsInitParameters struct {

	// Listener ID for application load balancers.
	// Listener ID for application load balancers.
	ListenerID *string `json:"listenerId,omitempty" tf:"listener_id,omitempty"`

	// ID of available load balancers.
	// ID of available load balancers.
	LoadBalancerID *string `json:"loadBalancerId,omitempty" tf:"load_balancer_id,omitempty"`

	// ID of forwarding rules.
	// ID of forwarding rules.
	RuleID *string `json:"ruleId,omitempty" tf:"rule_id,omitempty"`

	// Attribute list of target rules.
	// Attribute list of target rules.
	TargetAttribute []TargetAttributeInitParameters `json:"targetAttribute,omitempty" tf:"target_attribute,omitempty"`
}

type ForwardBalancerIdsObservation struct {

	// Listener ID for application load balancers.
	// Listener ID for application load balancers.
	ListenerID *string `json:"listenerId,omitempty" tf:"listener_id,omitempty"`

	// ID of available load balancers.
	// ID of available load balancers.
	LoadBalancerID *string `json:"loadBalancerId,omitempty" tf:"load_balancer_id,omitempty"`

	// ID of forwarding rules.
	// ID of forwarding rules.
	RuleID *string `json:"ruleId,omitempty" tf:"rule_id,omitempty"`

	// Attribute list of target rules.
	// Attribute list of target rules.
	TargetAttribute []TargetAttributeObservation `json:"targetAttribute,omitempty" tf:"target_attribute,omitempty"`
}

type ForwardBalancerIdsParameters struct {

	// Listener ID for application load balancers.
	// Listener ID for application load balancers.
	// +kubebuilder:validation:Optional
	ListenerID *string `json:"listenerId" tf:"listener_id,omitempty"`

	// ID of available load balancers.
	// ID of available load balancers.
	// +kubebuilder:validation:Optional
	LoadBalancerID *string `json:"loadBalancerId" tf:"load_balancer_id,omitempty"`

	// ID of forwarding rules.
	// ID of forwarding rules.
	// +kubebuilder:validation:Optional
	RuleID *string `json:"ruleId,omitempty" tf:"rule_id,omitempty"`

	// Attribute list of target rules.
	// Attribute list of target rules.
	// +kubebuilder:validation:Optional
	TargetAttribute []TargetAttributeParameters `json:"targetAttribute" tf:"target_attribute,omitempty"`
}

type ScalingGroupInitParameters struct {

	// An available ID for a launch configuration.
	// An available ID for a launch configuration.
	// +crossplane:generate:reference:type=ScalingConfig
	ConfigurationID *string `json:"configurationId,omitempty" tf:"configuration_id,omitempty"`

	// Reference to a ScalingConfig to populate configurationId.
	// +kubebuilder:validation:Optional
	ConfigurationIDRef *v1.Reference `json:"configurationIdRef,omitempty" tf:"-"`

	// Selector for a ScalingConfig to populate configurationId.
	// +kubebuilder:validation:Optional
	ConfigurationIDSelector *v1.Selector `json:"configurationIdSelector,omitempty" tf:"-"`

	// Default cooldown time in second, and default value is 300.
	// Default cooldown time in second, and default value is `300`.
	DefaultCooldown *float64 `json:"defaultCooldown,omitempty" tf:"default_cooldown,omitempty"`

	// Desired volume of CVM instances, which is between max_size and min_size.
	// Desired volume of CVM instances, which is between `max_size` and `min_size`.
	DesiredCapacity *float64 `json:"desiredCapacity,omitempty" tf:"desired_capacity,omitempty"`

	// The expected number of instances is synchronized with the maximum and minimum values. The default value is False. This parameter is effective only in the scenario where the expected number is not passed in when modifying the scaling group interface. True: When modifying the maximum or minimum value, if there is a conflict with the current expected number, the expected number is adjusted synchronously. For example, when modifying, if the minimum value 2 is passed in and the current expected number is 1, the expected number is adjusted synchronously to 2; False: When modifying the maximum or minimum value, if there is a conflict with the current expected number, an error message is displayed indicating that the modification is not allowed.
	// The expected number of instances is synchronized with the maximum and minimum values. The default value is `False`. This parameter is effective only in the scenario where the expected number is not passed in when modifying the scaling group interface. True: When modifying the maximum or minimum value, if there is a conflict with the current expected number, the expected number is adjusted synchronously. For example, when modifying, if the minimum value 2 is passed in and the current expected number is 1, the expected number is adjusted synchronously to 2; False: When modifying the maximum or minimum value, if there is a conflict with the current expected number, an error message is displayed indicating that the modification is not allowed.
	DesiredCapacitySyncWithMaxMinSize *bool `json:"desiredCapacitySyncWithMaxMinSize,omitempty" tf:"desired_capacity_sync_with_max_min_size,omitempty"`

	// List of application load balancers, which can't be specified with load_balancer_ids together.
	// List of application load balancers, which can't be specified with `load_balancer_ids` together.
	ForwardBalancerIds []ForwardBalancerIdsInitParameters `json:"forwardBalancerIds,omitempty" tf:"forward_balancer_ids,omitempty"`

	// Health check type of instances in a scaling group.CVM: confirm whether an instance is healthy based on the network status. If the pinged instance is unreachable, the instance will be considered unhealthy. For more information, see Instance Health CheckCLB: confirm whether an instance is healthy based on the CLB health check status. For more information, see Health Check Overview.If the parameter is set to CLB, the scaling group will check both the network status and the CLB health check status. If the network check indicates unhealthy, the HealthStatus field will return UNHEALTHY. If the CLB health check indicates unhealthy, the HealthStatus field will return CLB_UNHEALTHY. If both checks indicate unhealthy, the HealthStatus field will return UNHEALTHY|CLB_UNHEALTHY. Default value: CLB.
	// Health check type of instances in a scaling group.<br><li>CVM: confirm whether an instance is healthy based on the network status. If the pinged instance is unreachable, the instance will be considered unhealthy. For more information, see [Instance Health Check](https://intl.cloud.tencent.com/document/product/377/8553?from_cn_redirect=1)<br><li>CLB: confirm whether an instance is healthy based on the CLB health check status. For more information, see [Health Check Overview](https://intl.cloud.tencent.com/document/product/214/6097?from_cn_redirect=1).<br>If the parameter is set to `CLB`, the scaling group will check both the network status and the CLB health check status. If the network check indicates unhealthy, the `HealthStatus` field will return `UNHEALTHY`. If the CLB health check indicates unhealthy, the `HealthStatus` field will return `CLB_UNHEALTHY`. If both checks indicate unhealthy, the `HealthStatus` field will return `UNHEALTHY|CLB_UNHEALTHY`. Default value: `CLB`.
	HealthCheckType *string `json:"healthCheckType,omitempty" tf:"health_check_type,omitempty"`

	// Grace period of the CLB health check during which the IN_SERVICE instances added will not be marked as CLB_UNHEALTHY.Valid range: 0-7200, in seconds. Default value: 0.
	// Grace period of the CLB health check during which the `IN_SERVICE` instances added will not be marked as `CLB_UNHEALTHY`.<br>Valid range: 0-7200, in seconds. Default value: `0`.
	LBHealthCheckGracePeriod *float64 `json:"lbHealthCheckGracePeriod,omitempty" tf:"lb_health_check_grace_period,omitempty"`

	// ID list of traditional load balancers.
	// ID list of traditional load balancers.
	LoadBalancerIds []*string `json:"loadBalancerIds,omitempty" tf:"load_balancer_ids,omitempty"`

	// Maximum number of CVM instances. Valid value ranges: (0~2000).
	// Maximum number of CVM instances. Valid value ranges: (0~2000).
	MaxSize *float64 `json:"maxSize,omitempty" tf:"max_size,omitempty"`

	// Minimum number of CVM instances. Valid value ranges: (0~2000).
	// Minimum number of CVM instances. Valid value ranges: (0~2000).
	MinSize *float64 `json:"minSize,omitempty" tf:"min_size,omitempty"`

	// Multi zone or subnet strategy, Valid values: PRIORITY and EQUALITY.
	// Multi zone or subnet strategy, Valid values: PRIORITY and EQUALITY.
	MultiZoneSubnetPolicy *string `json:"multiZoneSubnetPolicy,omitempty" tf:"multi_zone_subnet_policy,omitempty"`

	// Specifies to which project the scaling group belongs.
	// Specifies to which project the scaling group belongs.
	ProjectID *float64 `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Enable unhealthy instance replacement. If set to true, AS will replace instances that are found unhealthy in the CLB health check.
	// Enable unhealthy instance replacement. If set to `true`, AS will replace instances that are found unhealthy in the CLB health check.
	ReplaceLoadBalancerUnhealthy *bool `json:"replaceLoadBalancerUnhealthy,omitempty" tf:"replace_load_balancer_unhealthy,omitempty"`

	// Replace mode of unhealthy replacement service. Valid values: RECREATE: Rebuild an instance to replace the original unhealthy instance. RESET: Performing a system reinstallation on unhealthy instances to keep information such as data disks, private IP addresses, and instance IDs unchanged. The instance login settings, HostName, enhanced services, and UserData will remain consistent with the current launch configuration. Default value: RECREATE. Note: This field may return null, indicating that no valid values can be obtained.
	// Replace mode of unhealthy replacement service. Valid values: RECREATE: Rebuild an instance to replace the original unhealthy instance. RESET: Performing a system reinstallation on unhealthy instances to keep information such as data disks, private IP addresses, and instance IDs unchanged. The instance login settings, HostName, enhanced services, and UserData will remain consistent with the current launch configuration. Default value: RECREATE. Note: This field may return null, indicating that no valid values can be obtained.
	ReplaceMode *string `json:"replaceMode,omitempty" tf:"replace_mode,omitempty"`

	// Enables unhealthy instance replacement. If set to true, AS will replace instances that are flagged as unhealthy by Cloud Monitor.
	// Enables unhealthy instance replacement. If set to `true`, AS will replace instances that are flagged as unhealthy by Cloud Monitor.
	ReplaceMonitorUnhealthy *bool `json:"replaceMonitorUnhealthy,omitempty" tf:"replace_monitor_unhealthy,omitempty"`

	// Available values for retry policies. Valid values: IMMEDIATE_RETRY and INCREMENTAL_INTERVALS.
	// Available values for retry policies. Valid values: IMMEDIATE_RETRY and INCREMENTAL_INTERVALS.
	RetryPolicy *string `json:"retryPolicy,omitempty" tf:"retry_policy,omitempty"`

	// Name of a scaling group.
	// Name of a scaling group.
	ScalingGroupName *string `json:"scalingGroupName,omitempty" tf:"scaling_group_name,omitempty"`

	// Indicates scaling mode which creates and terminates instances (classic method), or method first tries to start stopped instances (wake up stopped) to perform scaling operations. Available values: CLASSIC_SCALING, WAKE_UP_STOPPED_SCALING. Default: CLASSIC_SCALING.
	// Indicates scaling mode which creates and terminates instances (classic method), or method first tries to start stopped instances (wake up stopped) to perform scaling operations. Available values: `CLASSIC_SCALING`, `WAKE_UP_STOPPED_SCALING`. Default: `CLASSIC_SCALING`.
	ScalingMode *string `json:"scalingMode,omitempty" tf:"scaling_mode,omitempty"`

	// ID list of subnet, and for VPC it is required.
	// ID list of subnet, and for VPC it is required.
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// Tags of a scaling group.
	// Tags of a scaling group.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Available values for termination policies. Valid values: OLDEST_INSTANCE and NEWEST_INSTANCE.
	// Available values for termination policies. Valid values: OLDEST_INSTANCE and NEWEST_INSTANCE.
	TerminationPolicies []*string `json:"terminationPolicies,omitempty" tf:"termination_policies,omitempty"`

	// ID of VPC network.
	// ID of VPC network.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// List of available zones, for Basic network it is required.
	// List of available zones, for Basic network it is required.
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type ScalingGroupObservation struct {

	// An available ID for a launch configuration.
	// An available ID for a launch configuration.
	ConfigurationID *string `json:"configurationId,omitempty" tf:"configuration_id,omitempty"`

	// The time when the AS group was created.
	// The time when the AS group was created.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Default cooldown time in second, and default value is 300.
	// Default cooldown time in second, and default value is `300`.
	DefaultCooldown *float64 `json:"defaultCooldown,omitempty" tf:"default_cooldown,omitempty"`

	// Desired volume of CVM instances, which is between max_size and min_size.
	// Desired volume of CVM instances, which is between `max_size` and `min_size`.
	DesiredCapacity *float64 `json:"desiredCapacity,omitempty" tf:"desired_capacity,omitempty"`

	// The expected number of instances is synchronized with the maximum and minimum values. The default value is False. This parameter is effective only in the scenario where the expected number is not passed in when modifying the scaling group interface. True: When modifying the maximum or minimum value, if there is a conflict with the current expected number, the expected number is adjusted synchronously. For example, when modifying, if the minimum value 2 is passed in and the current expected number is 1, the expected number is adjusted synchronously to 2; False: When modifying the maximum or minimum value, if there is a conflict with the current expected number, an error message is displayed indicating that the modification is not allowed.
	// The expected number of instances is synchronized with the maximum and minimum values. The default value is `False`. This parameter is effective only in the scenario where the expected number is not passed in when modifying the scaling group interface. True: When modifying the maximum or minimum value, if there is a conflict with the current expected number, the expected number is adjusted synchronously. For example, when modifying, if the minimum value 2 is passed in and the current expected number is 1, the expected number is adjusted synchronously to 2; False: When modifying the maximum or minimum value, if there is a conflict with the current expected number, an error message is displayed indicating that the modification is not allowed.
	DesiredCapacitySyncWithMaxMinSize *bool `json:"desiredCapacitySyncWithMaxMinSize,omitempty" tf:"desired_capacity_sync_with_max_min_size,omitempty"`

	// List of application load balancers, which can't be specified with load_balancer_ids together.
	// List of application load balancers, which can't be specified with `load_balancer_ids` together.
	ForwardBalancerIds []ForwardBalancerIdsObservation `json:"forwardBalancerIds,omitempty" tf:"forward_balancer_ids,omitempty"`

	// Health check type of instances in a scaling group.CVM: confirm whether an instance is healthy based on the network status. If the pinged instance is unreachable, the instance will be considered unhealthy. For more information, see Instance Health CheckCLB: confirm whether an instance is healthy based on the CLB health check status. For more information, see Health Check Overview.If the parameter is set to CLB, the scaling group will check both the network status and the CLB health check status. If the network check indicates unhealthy, the HealthStatus field will return UNHEALTHY. If the CLB health check indicates unhealthy, the HealthStatus field will return CLB_UNHEALTHY. If both checks indicate unhealthy, the HealthStatus field will return UNHEALTHY|CLB_UNHEALTHY. Default value: CLB.
	// Health check type of instances in a scaling group.<br><li>CVM: confirm whether an instance is healthy based on the network status. If the pinged instance is unreachable, the instance will be considered unhealthy. For more information, see [Instance Health Check](https://intl.cloud.tencent.com/document/product/377/8553?from_cn_redirect=1)<br><li>CLB: confirm whether an instance is healthy based on the CLB health check status. For more information, see [Health Check Overview](https://intl.cloud.tencent.com/document/product/214/6097?from_cn_redirect=1).<br>If the parameter is set to `CLB`, the scaling group will check both the network status and the CLB health check status. If the network check indicates unhealthy, the `HealthStatus` field will return `UNHEALTHY`. If the CLB health check indicates unhealthy, the `HealthStatus` field will return `CLB_UNHEALTHY`. If both checks indicate unhealthy, the `HealthStatus` field will return `UNHEALTHY|CLB_UNHEALTHY`. Default value: `CLB`.
	HealthCheckType *string `json:"healthCheckType,omitempty" tf:"health_check_type,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Instance number of a scaling group.
	// Instance number of a scaling group.
	InstanceCount *float64 `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`

	// Grace period of the CLB health check during which the IN_SERVICE instances added will not be marked as CLB_UNHEALTHY.Valid range: 0-7200, in seconds. Default value: 0.
	// Grace period of the CLB health check during which the `IN_SERVICE` instances added will not be marked as `CLB_UNHEALTHY`.<br>Valid range: 0-7200, in seconds. Default value: `0`.
	LBHealthCheckGracePeriod *float64 `json:"lbHealthCheckGracePeriod,omitempty" tf:"lb_health_check_grace_period,omitempty"`

	// ID list of traditional load balancers.
	// ID list of traditional load balancers.
	LoadBalancerIds []*string `json:"loadBalancerIds,omitempty" tf:"load_balancer_ids,omitempty"`

	// Maximum number of CVM instances. Valid value ranges: (0~2000).
	// Maximum number of CVM instances. Valid value ranges: (0~2000).
	MaxSize *float64 `json:"maxSize,omitempty" tf:"max_size,omitempty"`

	// Minimum number of CVM instances. Valid value ranges: (0~2000).
	// Minimum number of CVM instances. Valid value ranges: (0~2000).
	MinSize *float64 `json:"minSize,omitempty" tf:"min_size,omitempty"`

	// Multi zone or subnet strategy, Valid values: PRIORITY and EQUALITY.
	// Multi zone or subnet strategy, Valid values: PRIORITY and EQUALITY.
	MultiZoneSubnetPolicy *string `json:"multiZoneSubnetPolicy,omitempty" tf:"multi_zone_subnet_policy,omitempty"`

	// Specifies to which project the scaling group belongs.
	// Specifies to which project the scaling group belongs.
	ProjectID *float64 `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Enable unhealthy instance replacement. If set to true, AS will replace instances that are found unhealthy in the CLB health check.
	// Enable unhealthy instance replacement. If set to `true`, AS will replace instances that are found unhealthy in the CLB health check.
	ReplaceLoadBalancerUnhealthy *bool `json:"replaceLoadBalancerUnhealthy,omitempty" tf:"replace_load_balancer_unhealthy,omitempty"`

	// Replace mode of unhealthy replacement service. Valid values: RECREATE: Rebuild an instance to replace the original unhealthy instance. RESET: Performing a system reinstallation on unhealthy instances to keep information such as data disks, private IP addresses, and instance IDs unchanged. The instance login settings, HostName, enhanced services, and UserData will remain consistent with the current launch configuration. Default value: RECREATE. Note: This field may return null, indicating that no valid values can be obtained.
	// Replace mode of unhealthy replacement service. Valid values: RECREATE: Rebuild an instance to replace the original unhealthy instance. RESET: Performing a system reinstallation on unhealthy instances to keep information such as data disks, private IP addresses, and instance IDs unchanged. The instance login settings, HostName, enhanced services, and UserData will remain consistent with the current launch configuration. Default value: RECREATE. Note: This field may return null, indicating that no valid values can be obtained.
	ReplaceMode *string `json:"replaceMode,omitempty" tf:"replace_mode,omitempty"`

	// Enables unhealthy instance replacement. If set to true, AS will replace instances that are flagged as unhealthy by Cloud Monitor.
	// Enables unhealthy instance replacement. If set to `true`, AS will replace instances that are flagged as unhealthy by Cloud Monitor.
	ReplaceMonitorUnhealthy *bool `json:"replaceMonitorUnhealthy,omitempty" tf:"replace_monitor_unhealthy,omitempty"`

	// Available values for retry policies. Valid values: IMMEDIATE_RETRY and INCREMENTAL_INTERVALS.
	// Available values for retry policies. Valid values: IMMEDIATE_RETRY and INCREMENTAL_INTERVALS.
	RetryPolicy *string `json:"retryPolicy,omitempty" tf:"retry_policy,omitempty"`

	// Name of a scaling group.
	// Name of a scaling group.
	ScalingGroupName *string `json:"scalingGroupName,omitempty" tf:"scaling_group_name,omitempty"`

	// Indicates scaling mode which creates and terminates instances (classic method), or method first tries to start stopped instances (wake up stopped) to perform scaling operations. Available values: CLASSIC_SCALING, WAKE_UP_STOPPED_SCALING. Default: CLASSIC_SCALING.
	// Indicates scaling mode which creates and terminates instances (classic method), or method first tries to start stopped instances (wake up stopped) to perform scaling operations. Available values: `CLASSIC_SCALING`, `WAKE_UP_STOPPED_SCALING`. Default: `CLASSIC_SCALING`.
	ScalingMode *string `json:"scalingMode,omitempty" tf:"scaling_mode,omitempty"`

	// Current status of a scaling group.
	// Current status of a scaling group.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// ID list of subnet, and for VPC it is required.
	// ID list of subnet, and for VPC it is required.
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// Tags of a scaling group.
	// Tags of a scaling group.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Available values for termination policies. Valid values: OLDEST_INSTANCE and NEWEST_INSTANCE.
	// Available values for termination policies. Valid values: OLDEST_INSTANCE and NEWEST_INSTANCE.
	TerminationPolicies []*string `json:"terminationPolicies,omitempty" tf:"termination_policies,omitempty"`

	// ID of VPC network.
	// ID of VPC network.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// List of available zones, for Basic network it is required.
	// List of available zones, for Basic network it is required.
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type ScalingGroupParameters struct {

	// An available ID for a launch configuration.
	// An available ID for a launch configuration.
	// +crossplane:generate:reference:type=ScalingConfig
	// +kubebuilder:validation:Optional
	ConfigurationID *string `json:"configurationId,omitempty" tf:"configuration_id,omitempty"`

	// Reference to a ScalingConfig to populate configurationId.
	// +kubebuilder:validation:Optional
	ConfigurationIDRef *v1.Reference `json:"configurationIdRef,omitempty" tf:"-"`

	// Selector for a ScalingConfig to populate configurationId.
	// +kubebuilder:validation:Optional
	ConfigurationIDSelector *v1.Selector `json:"configurationIdSelector,omitempty" tf:"-"`

	// Default cooldown time in second, and default value is 300.
	// Default cooldown time in second, and default value is `300`.
	// +kubebuilder:validation:Optional
	DefaultCooldown *float64 `json:"defaultCooldown,omitempty" tf:"default_cooldown,omitempty"`

	// Desired volume of CVM instances, which is between max_size and min_size.
	// Desired volume of CVM instances, which is between `max_size` and `min_size`.
	// +kubebuilder:validation:Optional
	DesiredCapacity *float64 `json:"desiredCapacity,omitempty" tf:"desired_capacity,omitempty"`

	// The expected number of instances is synchronized with the maximum and minimum values. The default value is False. This parameter is effective only in the scenario where the expected number is not passed in when modifying the scaling group interface. True: When modifying the maximum or minimum value, if there is a conflict with the current expected number, the expected number is adjusted synchronously. For example, when modifying, if the minimum value 2 is passed in and the current expected number is 1, the expected number is adjusted synchronously to 2; False: When modifying the maximum or minimum value, if there is a conflict with the current expected number, an error message is displayed indicating that the modification is not allowed.
	// The expected number of instances is synchronized with the maximum and minimum values. The default value is `False`. This parameter is effective only in the scenario where the expected number is not passed in when modifying the scaling group interface. True: When modifying the maximum or minimum value, if there is a conflict with the current expected number, the expected number is adjusted synchronously. For example, when modifying, if the minimum value 2 is passed in and the current expected number is 1, the expected number is adjusted synchronously to 2; False: When modifying the maximum or minimum value, if there is a conflict with the current expected number, an error message is displayed indicating that the modification is not allowed.
	// +kubebuilder:validation:Optional
	DesiredCapacitySyncWithMaxMinSize *bool `json:"desiredCapacitySyncWithMaxMinSize,omitempty" tf:"desired_capacity_sync_with_max_min_size,omitempty"`

	// List of application load balancers, which can't be specified with load_balancer_ids together.
	// List of application load balancers, which can't be specified with `load_balancer_ids` together.
	// +kubebuilder:validation:Optional
	ForwardBalancerIds []ForwardBalancerIdsParameters `json:"forwardBalancerIds,omitempty" tf:"forward_balancer_ids,omitempty"`

	// Health check type of instances in a scaling group.CVM: confirm whether an instance is healthy based on the network status. If the pinged instance is unreachable, the instance will be considered unhealthy. For more information, see Instance Health CheckCLB: confirm whether an instance is healthy based on the CLB health check status. For more information, see Health Check Overview.If the parameter is set to CLB, the scaling group will check both the network status and the CLB health check status. If the network check indicates unhealthy, the HealthStatus field will return UNHEALTHY. If the CLB health check indicates unhealthy, the HealthStatus field will return CLB_UNHEALTHY. If both checks indicate unhealthy, the HealthStatus field will return UNHEALTHY|CLB_UNHEALTHY. Default value: CLB.
	// Health check type of instances in a scaling group.<br><li>CVM: confirm whether an instance is healthy based on the network status. If the pinged instance is unreachable, the instance will be considered unhealthy. For more information, see [Instance Health Check](https://intl.cloud.tencent.com/document/product/377/8553?from_cn_redirect=1)<br><li>CLB: confirm whether an instance is healthy based on the CLB health check status. For more information, see [Health Check Overview](https://intl.cloud.tencent.com/document/product/214/6097?from_cn_redirect=1).<br>If the parameter is set to `CLB`, the scaling group will check both the network status and the CLB health check status. If the network check indicates unhealthy, the `HealthStatus` field will return `UNHEALTHY`. If the CLB health check indicates unhealthy, the `HealthStatus` field will return `CLB_UNHEALTHY`. If both checks indicate unhealthy, the `HealthStatus` field will return `UNHEALTHY|CLB_UNHEALTHY`. Default value: `CLB`.
	// +kubebuilder:validation:Optional
	HealthCheckType *string `json:"healthCheckType,omitempty" tf:"health_check_type,omitempty"`

	// Grace period of the CLB health check during which the IN_SERVICE instances added will not be marked as CLB_UNHEALTHY.Valid range: 0-7200, in seconds. Default value: 0.
	// Grace period of the CLB health check during which the `IN_SERVICE` instances added will not be marked as `CLB_UNHEALTHY`.<br>Valid range: 0-7200, in seconds. Default value: `0`.
	// +kubebuilder:validation:Optional
	LBHealthCheckGracePeriod *float64 `json:"lbHealthCheckGracePeriod,omitempty" tf:"lb_health_check_grace_period,omitempty"`

	// ID list of traditional load balancers.
	// ID list of traditional load balancers.
	// +kubebuilder:validation:Optional
	LoadBalancerIds []*string `json:"loadBalancerIds,omitempty" tf:"load_balancer_ids,omitempty"`

	// Maximum number of CVM instances. Valid value ranges: (0~2000).
	// Maximum number of CVM instances. Valid value ranges: (0~2000).
	// +kubebuilder:validation:Optional
	MaxSize *float64 `json:"maxSize,omitempty" tf:"max_size,omitempty"`

	// Minimum number of CVM instances. Valid value ranges: (0~2000).
	// Minimum number of CVM instances. Valid value ranges: (0~2000).
	// +kubebuilder:validation:Optional
	MinSize *float64 `json:"minSize,omitempty" tf:"min_size,omitempty"`

	// Multi zone or subnet strategy, Valid values: PRIORITY and EQUALITY.
	// Multi zone or subnet strategy, Valid values: PRIORITY and EQUALITY.
	// +kubebuilder:validation:Optional
	MultiZoneSubnetPolicy *string `json:"multiZoneSubnetPolicy,omitempty" tf:"multi_zone_subnet_policy,omitempty"`

	// Specifies to which project the scaling group belongs.
	// Specifies to which project the scaling group belongs.
	// +kubebuilder:validation:Optional
	ProjectID *float64 `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Enable unhealthy instance replacement. If set to true, AS will replace instances that are found unhealthy in the CLB health check.
	// Enable unhealthy instance replacement. If set to `true`, AS will replace instances that are found unhealthy in the CLB health check.
	// +kubebuilder:validation:Optional
	ReplaceLoadBalancerUnhealthy *bool `json:"replaceLoadBalancerUnhealthy,omitempty" tf:"replace_load_balancer_unhealthy,omitempty"`

	// Replace mode of unhealthy replacement service. Valid values: RECREATE: Rebuild an instance to replace the original unhealthy instance. RESET: Performing a system reinstallation on unhealthy instances to keep information such as data disks, private IP addresses, and instance IDs unchanged. The instance login settings, HostName, enhanced services, and UserData will remain consistent with the current launch configuration. Default value: RECREATE. Note: This field may return null, indicating that no valid values can be obtained.
	// Replace mode of unhealthy replacement service. Valid values: RECREATE: Rebuild an instance to replace the original unhealthy instance. RESET: Performing a system reinstallation on unhealthy instances to keep information such as data disks, private IP addresses, and instance IDs unchanged. The instance login settings, HostName, enhanced services, and UserData will remain consistent with the current launch configuration. Default value: RECREATE. Note: This field may return null, indicating that no valid values can be obtained.
	// +kubebuilder:validation:Optional
	ReplaceMode *string `json:"replaceMode,omitempty" tf:"replace_mode,omitempty"`

	// Enables unhealthy instance replacement. If set to true, AS will replace instances that are flagged as unhealthy by Cloud Monitor.
	// Enables unhealthy instance replacement. If set to `true`, AS will replace instances that are flagged as unhealthy by Cloud Monitor.
	// +kubebuilder:validation:Optional
	ReplaceMonitorUnhealthy *bool `json:"replaceMonitorUnhealthy,omitempty" tf:"replace_monitor_unhealthy,omitempty"`

	// Available values for retry policies. Valid values: IMMEDIATE_RETRY and INCREMENTAL_INTERVALS.
	// Available values for retry policies. Valid values: IMMEDIATE_RETRY and INCREMENTAL_INTERVALS.
	// +kubebuilder:validation:Optional
	RetryPolicy *string `json:"retryPolicy,omitempty" tf:"retry_policy,omitempty"`

	// Name of a scaling group.
	// Name of a scaling group.
	// +kubebuilder:validation:Optional
	ScalingGroupName *string `json:"scalingGroupName,omitempty" tf:"scaling_group_name,omitempty"`

	// Indicates scaling mode which creates and terminates instances (classic method), or method first tries to start stopped instances (wake up stopped) to perform scaling operations. Available values: CLASSIC_SCALING, WAKE_UP_STOPPED_SCALING. Default: CLASSIC_SCALING.
	// Indicates scaling mode which creates and terminates instances (classic method), or method first tries to start stopped instances (wake up stopped) to perform scaling operations. Available values: `CLASSIC_SCALING`, `WAKE_UP_STOPPED_SCALING`. Default: `CLASSIC_SCALING`.
	// +kubebuilder:validation:Optional
	ScalingMode *string `json:"scalingMode,omitempty" tf:"scaling_mode,omitempty"`

	// ID list of subnet, and for VPC it is required.
	// ID list of subnet, and for VPC it is required.
	// +kubebuilder:validation:Optional
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// Tags of a scaling group.
	// Tags of a scaling group.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Available values for termination policies. Valid values: OLDEST_INSTANCE and NEWEST_INSTANCE.
	// Available values for termination policies. Valid values: OLDEST_INSTANCE and NEWEST_INSTANCE.
	// +kubebuilder:validation:Optional
	TerminationPolicies []*string `json:"terminationPolicies,omitempty" tf:"termination_policies,omitempty"`

	// ID of VPC network.
	// ID of VPC network.
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// List of available zones, for Basic network it is required.
	// List of available zones, for Basic network it is required.
	// +kubebuilder:validation:Optional
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type TargetAttributeInitParameters struct {

	// Port number.
	// Port number.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Weight.
	// Weight.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type TargetAttributeObservation struct {

	// Port number.
	// Port number.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Weight.
	// Weight.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type TargetAttributeParameters struct {

	// Port number.
	// Port number.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port" tf:"port,omitempty"`

	// Weight.
	// Weight.
	// +kubebuilder:validation:Optional
	Weight *float64 `json:"weight" tf:"weight,omitempty"`
}

// ScalingGroupSpec defines the desired state of ScalingGroup
type ScalingGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ScalingGroupParameters `json:"forProvider"`
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
	InitProvider ScalingGroupInitParameters `json:"initProvider,omitempty"`
}

// ScalingGroupStatus defines the observed state of ScalingGroup.
type ScalingGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ScalingGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ScalingGroup is the Schema for the ScalingGroups API. Provides a resource to create a group of AS (Auto scaling) instances.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type ScalingGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.maxSize) || (has(self.initProvider) && has(self.initProvider.maxSize))",message="spec.forProvider.maxSize is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.minSize) || (has(self.initProvider) && has(self.initProvider.minSize))",message="spec.forProvider.minSize is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.scalingGroupName) || (has(self.initProvider) && has(self.initProvider.scalingGroupName))",message="spec.forProvider.scalingGroupName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.vpcId) || (has(self.initProvider) && has(self.initProvider.vpcId))",message="spec.forProvider.vpcId is a required parameter"
	Spec   ScalingGroupSpec   `json:"spec"`
	Status ScalingGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ScalingGroupList contains a list of ScalingGroups
type ScalingGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ScalingGroup `json:"items"`
}

// Repository type metadata.
var (
	ScalingGroup_Kind             = "ScalingGroup"
	ScalingGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ScalingGroup_Kind}.String()
	ScalingGroup_KindAPIVersion   = ScalingGroup_Kind + "." + CRDGroupVersion.String()
	ScalingGroup_GroupVersionKind = CRDGroupVersion.WithKind(ScalingGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&ScalingGroup{}, &ScalingGroupList{})
}
