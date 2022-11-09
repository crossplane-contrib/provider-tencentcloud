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

package as

import (
	tjconfig "github.com/crossplane/terrajet/pkg/config"
)

const shortGroupAs = "as"

// Configure configures the as group
func Configure(p *tjconfig.Provider) {
	p.AddResourceConfigurator("tencentcloud_as_scaling_config", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupAs
		r.Kind = "ScalingConfig"
	})

	p.AddResourceConfigurator("tencentcloud_as_scaling_group", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupAs
		r.Kind = "ScalingGroup"
		r.References["configuration_id"] = tjconfig.Reference{
			Type: "ScalingConfig",
		}
	})

	p.AddResourceConfigurator("tencentcloud_as_attachment", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupAs
		r.Kind = "Attachment"
		r.References["scaling_group_id"] = tjconfig.Reference{
			Type: "ScalingGroup",
		}
	})

	p.AddResourceConfigurator("tencentcloud_as_scaling_policy", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupAs
		r.Kind = "ScalingPolicy"
		r.References["scaling_group_id"] = tjconfig.Reference{
			Type: "ScalingGroup",
		}
	})

	p.AddResourceConfigurator("tencentcloud_as_schedule", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupAs
		r.Kind = "Schedule"
		r.References["scaling_group_id"] = tjconfig.Reference{
			Type: "ScalingGroup",
		}
	})

	p.AddResourceConfigurator("tencentcloud_as_lifecycle_hook", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupAs
		r.Kind = "LifecycleHook"
		r.References["scaling_group_id"] = tjconfig.Reference{
			Type: "ScalingGroup",
		}
	})

	p.AddResourceConfigurator("tencentcloud_as_notification", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupAs
		r.Kind = "Notification"
		r.References["scaling_group_id"] = tjconfig.Reference{
			Type: "ScalingGroup",
		}
	})
}
