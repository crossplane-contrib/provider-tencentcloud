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

package vpc

import (
	tjconfig "github.com/crossplane/terrajet/pkg/config"
)

const shortGroupVpc = "VPC"

// Configure configures the vpc group
func Configure(p *tjconfig.Provider) {
	p.AddResourceConfigurator("tencentcloud_vpc", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupVpc
	})

	p.AddResourceConfigurator("tencentcloud_subnet", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupVpc
		r.References["vpc_id"] = tjconfig.Reference{
			Type: "VPC",
		}
	})

	p.AddResourceConfigurator("tencentcloud_eip", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupVpc
		r.Kind = "Eip"
	})

	p.AddResourceConfigurator("tencentcloud_ha_vip", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupVpc
		r.Kind = "HaVip"
		r.References["vpc_id"] = tjconfig.Reference{
			Type: "VPC",
		}
		r.References["subnet_id"] = tjconfig.Reference{
			Type: "Subnet",
		}
	})

	p.AddResourceConfigurator("tencentcloud_route_table", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupVpc
		r.Kind = "RouteTable"
		r.References["vpc_id"] = tjconfig.Reference{
			Type: "VPC",
		}
	})

	p.AddResourceConfigurator("tencentcloud_route_table_entry", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupVpc
		r.Kind = "RouteTableEntry"
		r.References["route_table_id"] = tjconfig.Reference{
			Type: "RouteTable",
		}
	})

	p.AddResourceConfigurator("tencentcloud_route_entry", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupVpc
		r.Kind = "RouteEntry"
		r.References["route_table_id"] = tjconfig.Reference{
			Type: "RouteTable",
		}
		r.References["vpc_id"] = tjconfig.Reference{
			Type: "VPC",
		}
	})

	p.AddResourceConfigurator("tencentcloud_nat_gateway", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupVpc
		r.Kind = "NatGateway"
		r.References["vpc_id"] = tjconfig.Reference{
			Type: "VPC",
		}
	})

	p.AddResourceConfigurator("tencentcloud_nat_gateway_snat", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupVpc
		r.Kind = "NatGatewaySnat"
		r.References["subnet_id"] = tjconfig.Reference{
			Type: "Subnet",
		}
	})

	p.AddResourceConfigurator("tencentcloud_dnat", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupVpc
		r.Kind = "Dnat"
		r.References["vpc_id"] = tjconfig.Reference{
			Type: "VPC",
		}
		r.References["nat_id"] = tjconfig.Reference{
			Type: "NatGateway",
		}
	})

	p.AddResourceConfigurator("tencentcloud_security_group", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupVpc
		r.Kind = "SecurityGroup"
	})

	p.AddResourceConfigurator("tencentcloud_security_group_rule", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupVpc
		r.Kind = "SecurityGroupRule"
		r.References["security_group_id"] = tjconfig.Reference{
			Type: "SecurityGroup",
		}
	})

	p.AddResourceConfigurator("tencentcloud_security_group_lite_rule", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupVpc
		r.Kind = "SecurityGroupLiteRule"
		r.References["security_group_id"] = tjconfig.Reference{
			Type: "SecurityGroup",
		}
	})

	p.AddResourceConfigurator("tencentcloud_vpc_acl", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupVpc
		r.Kind = "VPCAcl"
		r.References["vpc_id"] = tjconfig.Reference{
			Type: "VPC",
		}
	})

	p.AddResourceConfigurator("tencentcloud_vpc_acl_attachment", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupVpc
		r.Kind = "VPVAclAttachment"
		r.References["acl_id"] = tjconfig.Reference{
			Type: "VPCAcl",
		}
		r.References["subnet_id"] = tjconfig.Reference{
			Type: "Subnet",
		}
	})

	p.AddResourceConfigurator("tencentcloud_vpc_bandwidth_package", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupVpc
		r.Kind = "VPCBandwidthPackage"
	})

	p.AddResourceConfigurator("tencentcloud_vpc_bandwidth_package_attachment", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupVpc
		r.Kind = "VPCBandwidthPackageAttachment"
		r.References["bandwidth_package_id"] = tjconfig.Reference{
			Type: "VPCBandwidthPackage",
		}

	})

	p.AddResourceConfigurator("tencentcloud_vpn_gateway", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupVpc
		r.Kind = "VPNGateway"
		r.References["vpc_id"] = tjconfig.Reference{
			Type: "VPC",
		}
	})

	p.AddResourceConfigurator("tencentcloud_vpn_connection", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupVpc
		r.Kind = "VPNConnection"
		r.References["vpc_id"] = tjconfig.Reference{
			Type: "VPC",
		}
		r.References["vpn_gateway_id"] = tjconfig.Reference{
			Type: "VPNGateway",
		}
		r.References["customer_gateway_id"] = tjconfig.Reference{
			Type: "VPNCustomerGateway",
		}
	})

	p.AddResourceConfigurator("tencentcloud_vpn_gateway_route", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupVpc
		r.Kind = "VPNGatewayRoute"
		r.References["vpn_gateway_id"] = tjconfig.Reference{
			Type: "VPNGateway",
		}
	})

	p.AddResourceConfigurator("tencentcloud_vpn_ssl_server", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupVpc
		r.Kind = "VPNSSLServer"
		r.References["vpn_gateway_id"] = tjconfig.Reference{
			Type: "VPNGateway",
		}
	})

	p.AddResourceConfigurator("tencentcloud_vpn_ssl_client", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupVpc
		r.Kind = "VPNSSLClient"
	})

	p.AddResourceConfigurator("tencentcloud_vpn_customer_gateway", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupVpc
		r.Kind = "VPNCustomerGateway"
	})

	p.AddResourceConfigurator("tencentcloud_address_template", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupVpc
		r.Kind = "AddressTemplate"
	})

	p.AddResourceConfigurator("tencentcloud_protocol_template", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupVpc
		r.Kind = "ProtocolTemplate"
	})

	p.AddResourceConfigurator("tencentcloud_address_template_group", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupVpc
		r.Kind = "AddressTemplateGroup"
		r.References["template_ids"] = tjconfig.Reference{
			Type: "AddressTemplate",
		}
	})

	p.AddResourceConfigurator("tencentcloud_protocol_template_group", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupVpc
		r.Kind = "ProtocolTemplateGroup"
		r.References["template_ids"] = tjconfig.Reference{
			Type: "ProtocolTemplate",
		}
	})

	p.AddResourceConfigurator("tencentcloud_eip_association", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupVpc
		r.Kind = "EipAssociation"
		r.References["eip_id"] = tjconfig.Reference{
			Type: "Eip",
		}
	})
}
