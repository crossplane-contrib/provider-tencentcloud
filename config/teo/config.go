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

package teo

import (
	tjconfig "github.com/crossplane/terrajet/pkg/config"
)

const shortGroupTeo = "teo"

// Configure configures the teo group
func Configure(p *tjconfig.Provider) {
	p.AddResourceConfigurator("tencentcloud_teo_zone", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupTeo
		r.Kind = "Zone"
	})

	p.AddResourceConfigurator("tencentcloud_teo_zone_setting", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupTeo
		r.Kind = "ZoneSetting"
		r.References["zone_id"] = tjconfig.Reference{
			Type: "Zone",
		}
	})

	p.AddResourceConfigurator("tencentcloud_teo_security_policy", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupTeo
		r.Kind = "SecurityPolicy"
		r.References["zone_id"] = tjconfig.Reference{
			Type: "Zone",
		}
	})

	p.AddResourceConfigurator("tencentcloud_teo_rule_engine", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupTeo
		r.Kind = "RuleEngine"
		r.References["zone_id"] = tjconfig.Reference{
			Type: "Zone",
		}
	})

	p.AddResourceConfigurator("tencentcloud_teo_rule_engine_priority", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupTeo
		r.Kind = "RuleEnginePriority"
		r.References["zone_id"] = tjconfig.Reference{
			Type: "Zone",
		}
	})

	p.AddResourceConfigurator("tencentcloud_teo_origin_group", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupTeo
		r.Kind = "OriginGroup"
		r.References["zone_id"] = tjconfig.Reference{
			Type: "Zone",
		}
	})

	p.AddResourceConfigurator("tencentcloud_teo_load_balancing", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupTeo
		r.Kind = "OriginGrouprity"
		r.References["origin_group_id"] = tjconfig.Reference{
			Type: "OriginGroup",
		}
	})

	p.AddResourceConfigurator("tencentcloud_teo_dns_sec", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupTeo
		r.Kind = "DnsSec"
		r.References["zone_id"] = tjconfig.Reference{
			Type: "Zone",
		}
	})

	p.AddResourceConfigurator("tencentcloud_teo_dns_record", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupTeo
		r.Kind = "DnsRecord"
		r.References["zone_id"] = tjconfig.Reference{
			Type: "Zone",
		}
	})

	p.AddResourceConfigurator("tencentcloud_teo_ddos_policy", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupTeo
		r.Kind = "DDosPolicy"
		r.References["zone_id"] = tjconfig.Reference{
			Type: "Zone",
		}
	})

	p.AddResourceConfigurator("tencentcloud_teo_application_proxy", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupTeo
		r.Kind = "ApplicationProxy"
		r.References["zone_id"] = tjconfig.Reference{
			Type: "Zone",
		}
	})

	p.AddResourceConfigurator("tencentcloud_teo_application_proxy_rule", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupTeo
		r.Kind = "ApplicationProxyRule"
		r.References["zone_id"] = tjconfig.Reference{
			Type: "Zone",
		}
	})
}
