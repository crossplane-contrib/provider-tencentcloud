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

package apigateway

import (
	tjconfig "github.com/crossplane/terrajet/pkg/config"
)

const shortGroupAPIGateway = "apigateway"

// Configure configures the apigateway group
func Configure(p *tjconfig.Provider) {
	p.AddResourceConfigurator("tencentcloud_api_gateway_service", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupAPIGateway
		r.Kind = "Service"
	})

	p.AddResourceConfigurator("tencentcloud_api_gateway_api", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupAPIGateway
		r.Kind = "Api"
		r.References["service_id"] = tjconfig.Reference{
			Type: "Service",
		}
	})

	p.AddResourceConfigurator("tencentcloud_api_gateway_api_key", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupAPIGateway
		r.Kind = "ApiKey"
	})

	p.AddResourceConfigurator("tencentcloud_api_gateway_usage_plan", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupAPIGateway
		r.Kind = "UsagePlan"
	})

	p.AddResourceConfigurator("tencentcloud_api_gateway_usage_plan_attachment", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupAPIGateway
		r.Kind = "UsagePlanAttachment"
		r.References["service_id"] = tjconfig.Reference{
			Type: "Service",
		}
		r.References["usage_plan_id"] = tjconfig.Reference{
			Type: "UsagePlan",
		}
		r.References["api_id"] = tjconfig.Reference{
			Type: "Api",
		}
	})

	p.AddResourceConfigurator("tencentcloud_api_gateway_ip_strategy", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupAPIGateway
		r.Kind = "IpStrategy"
		r.References["service_id"] = tjconfig.Reference{
			Type: "Service",
		}
	})

	p.AddResourceConfigurator("tencentcloud_api_gateway_custom_domain", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupAPIGateway
		r.Kind = "CustomDomain"
		r.References["service_id"] = tjconfig.Reference{
			Type: "Service",
		}
	})

	p.AddResourceConfigurator("tencentcloud_api_gateway_api_key_attachment", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupAPIGateway
		r.Kind = "ApiKeyAttachment"
		r.References["api_key_id"] = tjconfig.Reference{
			Type: "ApiKey",
		}
		r.References["usage_plan_id"] = tjconfig.Reference{
			Type: "UsagePlan",
		}
	})

	p.AddResourceConfigurator("tencentcloud_api_gateway_service_release", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupAPIGateway
		r.Kind = "ServiceRelease"
		r.References["service_id"] = tjconfig.Reference{
			Type: "Service",
		}
	})

	p.AddResourceConfigurator("tencentcloud_api_gateway_strategy_attachment", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupAPIGateway
		r.Kind = "StrategyAttachment"
		r.References["service_id"] = tjconfig.Reference{
			Type: "Service",
		}
		r.References["strategy_id"] = tjconfig.Reference{
			Type: "IpStrategy",
		}
		r.References["bind_api_id"] = tjconfig.Reference{
			Type: "Api",
		}
	})
}
