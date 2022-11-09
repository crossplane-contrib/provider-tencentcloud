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

package vod

import (
	tjconfig "github.com/crossplane/terrajet/pkg/config"
)

const shortGroupVod = "vod"

// Configure configures the vod group
func Configure(p *tjconfig.Provider) {
	p.AddResourceConfigurator("tencentcloud_vod_adaptive_dynamic_streaming_template", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupVod
		r.Kind = "AdaptiveDynamicStreamingTemplate"
	})
	p.AddResourceConfigurator("tencentcloud_vod_procedure_template", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupVod
		r.Kind = "ProcedureTemplate"
	})
	p.AddResourceConfigurator("tencentcloud_vod_snapshot_by_time_offset_template", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupVod
		r.Kind = "SnapshotByTimeOffsetTemplate"
	})
	p.AddResourceConfigurator("tencentcloud_vod_image_sprite_template", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupVod
		r.Kind = "ImageSpriteTemplate"
	})
	p.AddResourceConfigurator("tencentcloud_vod_super_player_config", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupVod
		r.Kind = "SuperPlayerConfig"
	})
	p.AddResourceConfigurator("tencentcloud_vod_sub_application", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupVod
		r.Kind = "SubApplication"
	})

}
