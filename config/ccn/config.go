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

package ccn

import "github.com/crossplane/upjet/pkg/config"

const shortGroupCcn = "ccn"

// Configure configures the ccn group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("tencentcloud_ccn", func(r *config.Resource) {
		r.ShortGroup = shortGroupCcn
		r.Kind = "CCN"
	})

	p.AddResourceConfigurator("tencentcloud_ccn_attachment", func(r *config.Resource) {
		r.ShortGroup = shortGroupCcn
		r.Kind = "Attachment"
		r.References["ccn_id"] = config.Reference{
			Type: "CCN",
		}
	})

	p.AddResourceConfigurator("tencentcloud_ccn_bandwidth_limit", func(r *config.Resource) {
		r.ShortGroup = shortGroupCcn
		r.Kind = "BandwidthLimit"
		r.References["ccn_id"] = config.Reference{
			Type: "CCN",
		}
	})

	p.AddResourceConfigurator("tencentcloud_ccn_route_table", func(r *config.Resource) {
		r.ShortGroup = shortGroupCcn
		r.Kind = "RouteTable"
		r.References["ccn_id"] = config.Reference{
			Type: "CCN",
		}
	})

	p.AddResourceConfigurator("tencentcloud_ccn_route_table_associate_instance_config", func(r *config.Resource) {
		r.ShortGroup = shortGroupCcn
		r.Kind = "RouteTableAssociateInstanceConfig"
		r.References["ccn_id"] = config.Reference{
			Type: "CCN",
		}
		r.References["route_table_id"] = config.Reference{
			Type: "RouteTable",
		}
	})

	p.AddResourceConfigurator("tencentcloud_ccn_route_table_broadcast_policies", func(r *config.Resource) {
		r.ShortGroup = shortGroupCcn
		r.Kind = "RouteTableBroadcastPolicies"
		r.References["ccn_id"] = config.Reference{
			Type: "CCN",
		}
		r.References["route_table_id"] = config.Reference{
			Type: "RouteTable",
		}
	})

	p.AddResourceConfigurator("tencentcloud_ccn_route_table_input_policies", func(r *config.Resource) {
		r.ShortGroup = shortGroupCcn
		r.Kind = "RouteTableInputPolicies"
		r.References["ccn_id"] = config.Reference{
			Type: "CCN",
		}
		r.References["route_table_id"] = config.Reference{
			Type: "RouteTable",
		}
	})

	p.AddResourceConfigurator("tencentcloud_ccn_route_table_selection_policies", func(r *config.Resource) {
		r.ShortGroup = shortGroupCcn
		r.Kind = "RouteTableSelectionPolicies"
		r.References["ccn_id"] = config.Reference{
			Type: "CCN",
		}
	})

	p.AddResourceConfigurator("tencentcloud_ccn_routes", func(r *config.Resource) {
		r.ShortGroup = shortGroupCcn
		r.Kind = "Routes"
		r.References["ccn_id"] = config.Reference{
			Type: "CCN",
		}
	})
}
