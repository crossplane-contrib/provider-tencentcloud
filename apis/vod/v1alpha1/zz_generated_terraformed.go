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
	"github.com/pkg/errors"

	"github.com/crossplane/terrajet/pkg/resource"
	"github.com/crossplane/terrajet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this AdaptiveDynamicStreamingTemplate
func (mg *AdaptiveDynamicStreamingTemplate) GetTerraformResourceType() string {
	return "tencentcloud_vod_adaptive_dynamic_streaming_template"
}

// GetConnectionDetailsMapping for this AdaptiveDynamicStreamingTemplate
func (tr *AdaptiveDynamicStreamingTemplate) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this AdaptiveDynamicStreamingTemplate
func (tr *AdaptiveDynamicStreamingTemplate) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this AdaptiveDynamicStreamingTemplate
func (tr *AdaptiveDynamicStreamingTemplate) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this AdaptiveDynamicStreamingTemplate
func (tr *AdaptiveDynamicStreamingTemplate) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this AdaptiveDynamicStreamingTemplate
func (tr *AdaptiveDynamicStreamingTemplate) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this AdaptiveDynamicStreamingTemplate
func (tr *AdaptiveDynamicStreamingTemplate) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this AdaptiveDynamicStreamingTemplate using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *AdaptiveDynamicStreamingTemplate) LateInitialize(attrs []byte) (bool, error) {
	params := &AdaptiveDynamicStreamingTemplateParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *AdaptiveDynamicStreamingTemplate) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this ImageSpriteTemplate
func (mg *ImageSpriteTemplate) GetTerraformResourceType() string {
	return "tencentcloud_vod_image_sprite_template"
}

// GetConnectionDetailsMapping for this ImageSpriteTemplate
func (tr *ImageSpriteTemplate) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this ImageSpriteTemplate
func (tr *ImageSpriteTemplate) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ImageSpriteTemplate
func (tr *ImageSpriteTemplate) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ImageSpriteTemplate
func (tr *ImageSpriteTemplate) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ImageSpriteTemplate
func (tr *ImageSpriteTemplate) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ImageSpriteTemplate
func (tr *ImageSpriteTemplate) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this ImageSpriteTemplate using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ImageSpriteTemplate) LateInitialize(attrs []byte) (bool, error) {
	params := &ImageSpriteTemplateParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ImageSpriteTemplate) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this ProcedureTemplate
func (mg *ProcedureTemplate) GetTerraformResourceType() string {
	return "tencentcloud_vod_procedure_template"
}

// GetConnectionDetailsMapping for this ProcedureTemplate
func (tr *ProcedureTemplate) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this ProcedureTemplate
func (tr *ProcedureTemplate) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ProcedureTemplate
func (tr *ProcedureTemplate) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ProcedureTemplate
func (tr *ProcedureTemplate) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ProcedureTemplate
func (tr *ProcedureTemplate) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ProcedureTemplate
func (tr *ProcedureTemplate) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this ProcedureTemplate using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ProcedureTemplate) LateInitialize(attrs []byte) (bool, error) {
	params := &ProcedureTemplateParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ProcedureTemplate) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SnapshotByTimeOffsetTemplate
func (mg *SnapshotByTimeOffsetTemplate) GetTerraformResourceType() string {
	return "tencentcloud_vod_snapshot_by_time_offset_template"
}

// GetConnectionDetailsMapping for this SnapshotByTimeOffsetTemplate
func (tr *SnapshotByTimeOffsetTemplate) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SnapshotByTimeOffsetTemplate
func (tr *SnapshotByTimeOffsetTemplate) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SnapshotByTimeOffsetTemplate
func (tr *SnapshotByTimeOffsetTemplate) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SnapshotByTimeOffsetTemplate
func (tr *SnapshotByTimeOffsetTemplate) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SnapshotByTimeOffsetTemplate
func (tr *SnapshotByTimeOffsetTemplate) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SnapshotByTimeOffsetTemplate
func (tr *SnapshotByTimeOffsetTemplate) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SnapshotByTimeOffsetTemplate using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SnapshotByTimeOffsetTemplate) LateInitialize(attrs []byte) (bool, error) {
	params := &SnapshotByTimeOffsetTemplateParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SnapshotByTimeOffsetTemplate) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SubApplication
func (mg *SubApplication) GetTerraformResourceType() string {
	return "tencentcloud_vod_sub_application"
}

// GetConnectionDetailsMapping for this SubApplication
func (tr *SubApplication) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SubApplication
func (tr *SubApplication) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SubApplication
func (tr *SubApplication) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SubApplication
func (tr *SubApplication) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SubApplication
func (tr *SubApplication) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SubApplication
func (tr *SubApplication) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SubApplication using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SubApplication) LateInitialize(attrs []byte) (bool, error) {
	params := &SubApplicationParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SubApplication) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SuperPlayerConfig
func (mg *SuperPlayerConfig) GetTerraformResourceType() string {
	return "tencentcloud_vod_super_player_config"
}

// GetConnectionDetailsMapping for this SuperPlayerConfig
func (tr *SuperPlayerConfig) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SuperPlayerConfig
func (tr *SuperPlayerConfig) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SuperPlayerConfig
func (tr *SuperPlayerConfig) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SuperPlayerConfig
func (tr *SuperPlayerConfig) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SuperPlayerConfig
func (tr *SuperPlayerConfig) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SuperPlayerConfig
func (tr *SuperPlayerConfig) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SuperPlayerConfig using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SuperPlayerConfig) LateInitialize(attrs []byte) (bool, error) {
	params := &SuperPlayerConfigParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SuperPlayerConfig) GetTerraformSchemaVersion() int {
	return 0
}
