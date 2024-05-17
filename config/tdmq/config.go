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

package tdmq

import "github.com/crossplane/upjet/pkg/config"

const shortGroupTdmq = "tdmq"

// Configure configures the tdmq group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("tencentcloud_tdmq_instance", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroupTdmq
		r.Kind = "Instance"
	})

	p.AddResourceConfigurator("tencentcloud_tdmq_namespace", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroupTdmq
		r.Kind = "TdmqNamespace"
		r.References["cluster_id"] = config.Reference{
			Type: "Instance",
		}
	})

	p.AddResourceConfigurator("tencentcloud_tdmq_topic", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroupTdmq
		r.Kind = "Topic"
		r.References["environ_id"] = config.Reference{
			Type: "TdmqNamespace",
		}
		r.References["cluster_id"] = config.Reference{
			Type: "Instance",
		}

		if t, ok := r.TerraformResource.Schema["topic_type"]; ok {
			t.Optional = false
			t.Computed = true
		}
	})

	p.AddResourceConfigurator("tencentcloud_tdmq_role", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroupTdmq
		r.Kind = "Role"
		r.References["cluster_id"] = config.Reference{
			Type: "Instance",
		}
	})

	p.AddResourceConfigurator("tencentcloud_tdmq_namespace_role_attachment", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroupTdmq
		r.Kind = "NamespaceRoleAttachment"
		r.References["cluster_id"] = config.Reference{
			Type: "Instance",
		}
		r.References["environ_id"] = config.Reference{
			Type: "TdmqNamespace",
		}
	})

}
