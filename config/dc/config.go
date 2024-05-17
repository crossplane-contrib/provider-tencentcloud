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

package dc

import "github.com/crossplane/upjet/pkg/config"

const shortGroupDc = "dc"

// Configure configures the dc group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("tencentcloud_dcx", func(r *config.Resource) {
		r.ShortGroup = shortGroupDc
		r.Kind = "Dcx"
	})
	p.AddResourceConfigurator("tencentcloud_dc_gateway", func(r *config.Resource) {
		r.ShortGroup = shortGroupDc
		r.Kind = "Gateway"
		r.References["vpc_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.VPC",
		}
	})

	p.AddResourceConfigurator("tencentcloud_dc_gateway_ccn_route", func(r *config.Resource) {
		r.ShortGroup = shortGroupDc
		r.Kind = "GatewayCcnRoute"
		r.References["dcg_id"] = config.Reference{
			Type: "Gateway",
		}
	})
}
