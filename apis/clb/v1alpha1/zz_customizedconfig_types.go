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

type CustomizedConfigInitParameters struct {

	// Content of Customized Config.
	// Content of Customized Config.
	ConfigContent *string `json:"configContent,omitempty" tf:"config_content,omitempty"`

	// Name of Customized Config.
	// Name of Customized Config.
	ConfigName *string `json:"configName,omitempty" tf:"config_name,omitempty"`

	// List of LoadBalancer Ids.
	// List of LoadBalancer Ids.
	// +listType=set
	LoadBalancerIds []*string `json:"loadBalancerIds,omitempty" tf:"load_balancer_ids,omitempty"`
}

type CustomizedConfigObservation struct {

	// Content of Customized Config.
	// Content of Customized Config.
	ConfigContent *string `json:"configContent,omitempty" tf:"config_content,omitempty"`

	// Name of Customized Config.
	// Name of Customized Config.
	ConfigName *string `json:"configName,omitempty" tf:"config_name,omitempty"`

	// Create time of Customized Config.
	// Create time of Customized Config.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// ID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// List of LoadBalancer Ids.
	// List of LoadBalancer Ids.
	// +listType=set
	LoadBalancerIds []*string `json:"loadBalancerIds,omitempty" tf:"load_balancer_ids,omitempty"`

	// Update time of Customized Config.
	// Update time of Customized Config.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type CustomizedConfigParameters struct {

	// Content of Customized Config.
	// Content of Customized Config.
	// +kubebuilder:validation:Optional
	ConfigContent *string `json:"configContent,omitempty" tf:"config_content,omitempty"`

	// Name of Customized Config.
	// Name of Customized Config.
	// +kubebuilder:validation:Optional
	ConfigName *string `json:"configName,omitempty" tf:"config_name,omitempty"`

	// List of LoadBalancer Ids.
	// List of LoadBalancer Ids.
	// +kubebuilder:validation:Optional
	// +listType=set
	LoadBalancerIds []*string `json:"loadBalancerIds,omitempty" tf:"load_balancer_ids,omitempty"`
}

// CustomizedConfigSpec defines the desired state of CustomizedConfig
type CustomizedConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CustomizedConfigParameters `json:"forProvider"`
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
	InitProvider CustomizedConfigInitParameters `json:"initProvider,omitempty"`
}

// CustomizedConfigStatus defines the observed state of CustomizedConfig.
type CustomizedConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CustomizedConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// CustomizedConfig is the Schema for the CustomizedConfigs API. Provides a resource to create a CLB customized config which type is
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloud}
type CustomizedConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.configContent) || (has(self.initProvider) && has(self.initProvider.configContent))",message="spec.forProvider.configContent is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.configName) || (has(self.initProvider) && has(self.initProvider.configName))",message="spec.forProvider.configName is a required parameter"
	Spec   CustomizedConfigSpec   `json:"spec"`
	Status CustomizedConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CustomizedConfigList contains a list of CustomizedConfigs
type CustomizedConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CustomizedConfig `json:"items"`
}

// Repository type metadata.
var (
	CustomizedConfig_Kind             = "CustomizedConfig"
	CustomizedConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CustomizedConfig_Kind}.String()
	CustomizedConfig_KindAPIVersion   = CustomizedConfig_Kind + "." + CRDGroupVersion.String()
	CustomizedConfig_GroupVersionKind = CRDGroupVersion.WithKind(CustomizedConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&CustomizedConfig{}, &CustomizedConfigList{})
}
