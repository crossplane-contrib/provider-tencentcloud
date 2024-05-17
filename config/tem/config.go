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

package tem

import "github.com/crossplane/upjet/pkg/config"

const shortGroupTem = "tem"

// Configure configures the tem group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("tencentcloud_tem_environment", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroupTem
		r.Kind = "Environment"
	})

	p.AddResourceConfigurator("tencentcloud_tem_app_config", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroupTem
		r.Kind = "AppConfig"
		r.References["environment_id"] = config.Reference{
			Type: "Environment",
		}
	})

	p.AddResourceConfigurator("tencentcloud_tem_application", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroupTem
		r.Kind = "Application"
	})

	p.AddResourceConfigurator("tencentcloud_tem_gateway", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroupTem
		r.Kind = "Gateway"
	})

	p.AddResourceConfigurator("tencentcloud_tem_workload", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroupTem
		r.Kind = "Workload"
		r.References["application_id"] = config.Reference{
			Type: "Application",
		}
		r.References["environment_id"] = config.Reference{
			Type: "Environment",
		}
	})

	p.AddResourceConfigurator("tencentcloud_tem_scale_rule", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroupTem
		r.Kind = "ScaleRule"
		r.References["application_id"] = config.Reference{
			Type: "Application",
		}
		r.References["environment_id"] = config.Reference{
			Type: "Environment",
		}
		r.References["workload_id"] = config.Reference{
			Type: "Workload",
		}
	})

	p.AddResourceConfigurator("tencentcloud_tem_log_config", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroupTem
		r.Kind = "LogConfig"
		r.References["application_id"] = config.Reference{
			Type: "Application",
		}
		r.References["environment_id"] = config.Reference{
			Type: "Environment",
		}
		r.References["workload_id"] = config.Reference{
			Type: "Workload",
		}
	})
}
