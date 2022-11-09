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

package clb

import (
	tjconfig "github.com/crossplane/terrajet/pkg/config"
)

// Configure configures the clb group
func Configure(p *tjconfig.Provider) {
	p.AddResourceConfigurator("tencentcloud_clb_instance", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = "clb"
		r.Kind = "Instance"
		r.References["vpc_id"] = tjconfig.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.VPC",
		}
		r.References["subnet_id"] = tjconfig.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.Subnet",
		}
	})

	p.AddResourceConfigurator("tencentcloud_clb_listener", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = "clb"
		r.Kind = "Listener"
		r.References["clb_id"] = tjconfig.Reference{
			Type: "Instance",
		}
	})

	p.AddResourceConfigurator("tencentcloud_clb_listener_rule", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = "clb"
		r.Kind = "ListenerRule"
		r.References["clb_id"] = tjconfig.Reference{
			Type: "Instance",
		}
		r.References["listener_id"] = tjconfig.Reference{
			Type: "Listener",
		}
	})

	p.AddResourceConfigurator("tencentcloud_clb_attachment", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = "clb"
		r.Kind = "Attachment"
		r.References["clb_id"] = tjconfig.Reference{
			Type: "Instance",
		}
		r.References["listener_id"] = tjconfig.Reference{
			Type: "Listener",
		}
		r.References["rule_id"] = tjconfig.Reference{
			Type: "ListenerRule",
		}
	})

	p.AddResourceConfigurator("tencentcloud_clb_customized_config", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = "clb"
		r.Kind = "CustomizedConfig"
	})

	p.AddResourceConfigurator("tencentcloud_clb_log_set", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = "clb"
		r.Kind = "LogSet"
	})

	p.AddResourceConfigurator("tencentcloud_clb_log_topic", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = "clb"
		r.Kind = "LogTopic"
		r.References["log_set_id"] = tjconfig.Reference{
			Type: "LogSet",
		}
	})

	p.AddResourceConfigurator("tencentcloud_clb_redirection", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = "clb"
		r.Kind = "Redirection"
		r.References["clb_id"] = tjconfig.Reference{
			Type: "Instance",
		}
		r.References["source_listener_id"] = tjconfig.Reference{
			Type: "Listener",
		}
		r.References["target_listener_id"] = tjconfig.Reference{
			Type: "Listener",
		}
		r.References["source_rule_id"] = tjconfig.Reference{
			Type: "ListenerRule",
		}
		r.References["target_rule_id"] = tjconfig.Reference{
			Type: "ListenerRule",
		}
	})

	p.AddResourceConfigurator("tencentcloud_clb_snat_ip", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = "clb"
		r.Kind = "SnatIp"
		r.References["clb_id"] = tjconfig.Reference{
			Type: "Instance",
		}
	})

	p.AddResourceConfigurator("tencentcloud_clb_target_group", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = "clb"
		r.Kind = "TargetGroup"
	})

	p.AddResourceConfigurator("tencentcloud_clb_target_group_attachment", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = "clb"
		r.Kind = "TargetGroupAttachment"
		r.References["clb_id"] = tjconfig.Reference{
			Type: "Instance",
		}
		r.References["listener_id"] = tjconfig.Reference{
			Type: "Listener",
		}
		r.References["rule_id"] = tjconfig.Reference{
			Type: "ListenerRule",
		}
		r.References["target_group_id"] = tjconfig.Reference{
			Type: "TargetGroup",
		}
	})

	p.AddResourceConfigurator("tencentcloud_clb_target_group_instance_attachment", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = "clb"
		r.Kind = "TargetGroupInstanceAttachment"
		r.References["target_group_id"] = tjconfig.Reference{
			Type: "TargetGroup",
		}
	})

	p.AddResourceConfigurator("tencentcloud_lb", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = "clb"
		r.Kind = "LB"
	})

	p.AddResourceConfigurator("tencentcloud_alb_server_attachment", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = "clb"
		r.Kind = "AlbServerAttachment"
	})
}
