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

type LogConfigObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type LogConfigParameters struct {

	// application ID.
	// +crossplane:generate:reference:type=Application
	// +kubebuilder:validation:Optional
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// +kubebuilder:validation:Optional
	ApplicationIDRef *v1.Reference `json:"applicationIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ApplicationIDSelector *v1.Selector `json:"applicationIdSelector,omitempty" tf:"-"`

	// regex pattern.
	// +kubebuilder:validation:Optional
	BeginningRegex *string `json:"beginningRegex,omitempty" tf:"beginning_regex,omitempty"`

	// environment ID.
	// +crossplane:generate:reference:type=Environment
	// +kubebuilder:validation:Optional
	EnvironmentID *string `json:"environmentId,omitempty" tf:"environment_id,omitempty"`

	// +kubebuilder:validation:Optional
	EnvironmentIDRef *v1.Reference `json:"environmentIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	EnvironmentIDSelector *v1.Selector `json:"environmentIdSelector,omitempty" tf:"-"`

	// file name pattern if container_file.
	// +kubebuilder:validation:Optional
	FilePattern *string `json:"filePattern,omitempty" tf:"file_pattern,omitempty"`

	// container_stdout or container_file.
	// +kubebuilder:validation:Required
	InputType *string `json:"inputType" tf:"input_type,omitempty"`

	// directory if container_file.
	// +kubebuilder:validation:Optional
	LogPath *string `json:"logPath,omitempty" tf:"log_path,omitempty"`

	// minimalist_log or multiline_log.
	// +kubebuilder:validation:Required
	LogType *string `json:"logType" tf:"log_type,omitempty"`

	// logset.
	// +kubebuilder:validation:Required
	LogsetID *string `json:"logsetId" tf:"logset_id,omitempty"`

	// appConfig name.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// topic.
	// +kubebuilder:validation:Required
	TopicID *string `json:"topicId" tf:"topic_id,omitempty"`

	// application ID, which is combined by environment ID and application ID, like `en-o5edaepv#app-3j29aa2p`.
	// +crossplane:generate:reference:type=Workload
	// +kubebuilder:validation:Optional
	WorkloadID *string `json:"workloadId,omitempty" tf:"workload_id,omitempty"`

	// +kubebuilder:validation:Optional
	WorkloadIDRef *v1.Reference `json:"workloadIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	WorkloadIDSelector *v1.Selector `json:"workloadIdSelector,omitempty" tf:"-"`
}

// LogConfigSpec defines the desired state of LogConfig
type LogConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LogConfigParameters `json:"forProvider"`
}

// LogConfigStatus defines the observed state of LogConfig.
type LogConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LogConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LogConfig is the Schema for the LogConfigs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tencentcloudjet}
type LogConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LogConfigSpec   `json:"spec"`
	Status            LogConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LogConfigList contains a list of LogConfigs
type LogConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LogConfig `json:"items"`
}

// Repository type metadata.
var (
	LogConfig_Kind             = "LogConfig"
	LogConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LogConfig_Kind}.String()
	LogConfig_KindAPIVersion   = LogConfig_Kind + "." + CRDGroupVersion.String()
	LogConfig_GroupVersionKind = CRDGroupVersion.WithKind(LogConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&LogConfig{}, &LogConfigList{})
}
