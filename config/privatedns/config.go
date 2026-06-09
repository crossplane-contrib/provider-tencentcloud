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

package privatedns

import "github.com/crossplane/upjet/pkg/config"

const shortGroupPrivateDNS = "privatedns"

// Configure configures the privatedns group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("tencentcloud_private_dns_end_point", func(r *config.Resource) {
		r.ShortGroup = shortGroupPrivateDNS
		r.Kind = "EndPoint"
	})
	p.AddResourceConfigurator("tencentcloud_private_dns_extend_end_point", func(r *config.Resource) {
		r.ShortGroup = shortGroupPrivateDNS
		r.Kind = "ExtendEndPoint"
	})
	p.AddResourceConfigurator("tencentcloud_private_dns_forward_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroupPrivateDNS
		r.Kind = "ForwardRule"
		r.References["zone_id"] = config.Reference{
			Type: "Zone",
		}
		r.References["end_point_id"] = config.Reference{
			Type: "EndPoint",
		}
	})
	p.AddResourceConfigurator("tencentcloud_private_dns_record", func(r *config.Resource) {
		r.ShortGroup = shortGroupPrivateDNS
		r.Kind = "Record"
		r.References["zone_id"] = config.Reference{
			Type: "Zone",
		}
	})
	p.AddResourceConfigurator("tencentcloud_private_dns_zone", func(r *config.Resource) {
		r.ShortGroup = shortGroupPrivateDNS
		r.Kind = "Zone"
	})
	p.AddResourceConfigurator("tencentcloud_private_dns_zone_vpc_attachment", func(r *config.Resource) {
		r.ShortGroup = shortGroupPrivateDNS
		r.Kind = "ZoneVPCAttachment"
		r.References["zone_id"] = config.Reference{
			Type: "Zone",
		}
	})
}
