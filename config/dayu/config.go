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

package dayu

import (
	tjconfig "github.com/crossplane/terrajet/pkg/config"
)

const shortGroupDayu = "dayu"

// Configure configures the dayu group
func Configure(p *tjconfig.Provider) {
	p.AddResourceConfigurator("tencentcloud_dayu_eip", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupDayu
		r.Kind = "eip"
	})
	p.AddResourceConfigurator("tencentcloud_dayu_l4_rule", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupDayu
		r.Kind = "L4Rule"
	})

	p.AddResourceConfigurator("tencentcloud_dayu_l4_rule_v2", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupDayu
		r.Kind = "L4RuleV2"
	})

	p.AddResourceConfigurator("tencentcloud_dayu_l7_rule", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupDayu
		r.Kind = "L7Rule"
	})

	p.AddResourceConfigurator("tencentcloud_dayu_l7_rule_v2", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupDayu
		r.Kind = "L7RuleV2"
	})

	p.AddResourceConfigurator("tencentcloud_dayu_ddos_policy", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupDayu
		r.Kind = "DdosPolicy"
	})

	p.AddResourceConfigurator("tencentcloud_dayu_ddos_policy_v2", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupDayu
		r.Kind = "DdosPolicyV2"
	})

	p.AddResourceConfigurator("tencentcloud_dayu_ddos_policy_case", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupDayu
		r.Kind = "DdosPolicyCase"
	})

	p.AddResourceConfigurator("tencentcloud_dayu_ddos_policy_attachment", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupDayu
		r.Kind = "DdosPolicyAttachment"
	})

	p.AddResourceConfigurator("tencentcloud_dayu_cc_policy_v2", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupDayu
		r.Kind = "CcPolicyV2"
	})

	p.AddResourceConfigurator("tencentcloud_dayu_cc_https_policy", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupDayu
		r.Kind = "CcHttpsPolicy"
	})

	p.AddResourceConfigurator("tencentcloud_dayu_cc_http_policy", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupDayu
		r.Kind = "CcHttpPolicy"
	})
}
