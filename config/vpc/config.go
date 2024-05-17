package vpc

import "github.com/crossplane/upjet/pkg/config"

const shortGroupVpc = "VPC"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("tencentcloud_vpc", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "vpc"
		r.ShortGroup = shortGroupVpc
	})

	p.AddResourceConfigurator("tencentcloud_subnet", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = shortGroupVpc
		r.References["vpc_id"] = config.Reference{
			Type: "VPC",
		}
	})

	p.AddResourceConfigurator("tencentcloud_eip", func(r *config.Resource) {
		r.ShortGroup = shortGroupVpc
		r.Kind = "Eip"
	})

	p.AddResourceConfigurator("tencentcloud_ha_vip", func(r *config.Resource) {
		r.ShortGroup = shortGroupVpc
		r.Kind = "HaVip"
		r.References["vpc_id"] = config.Reference{
			Type: "VPC",
		}
		r.References["subnet_id"] = config.Reference{
			Type: "Subnet",
		}
	})

	p.AddResourceConfigurator("tencentcloud_route_table", func(r *config.Resource) {
		r.ShortGroup = shortGroupVpc
		r.Kind = "RouteTable"
		r.References["vpc_id"] = config.Reference{
			Type: "VPC",
		}
	})

	p.AddResourceConfigurator("tencentcloud_route_table_entry", func(r *config.Resource) {
		r.ShortGroup = shortGroupVpc
		r.Kind = "RouteTableEntry"
		r.References["route_table_id"] = config.Reference{
			Type: "RouteTable",
		}
	})

	p.AddResourceConfigurator("tencentcloud_route_entry", func(r *config.Resource) {
		r.ShortGroup = shortGroupVpc
		r.Kind = "RouteEntry"
		r.References["route_table_id"] = config.Reference{
			Type: "RouteTable",
		}
		r.References["vpc_id"] = config.Reference{
			Type: "VPC",
		}
	})

	p.AddResourceConfigurator("tencentcloud_nat_gateway", func(r *config.Resource) {
		r.ShortGroup = shortGroupVpc
		r.Kind = "NatGateway"
		r.References["vpc_id"] = config.Reference{
			Type: "VPC",
		}
	})

	p.AddResourceConfigurator("tencentcloud_nat_gateway_snat", func(r *config.Resource) {
		r.ShortGroup = shortGroupVpc
		r.Kind = "NatGatewaySnat"
		r.References["subnet_id"] = config.Reference{
			Type: "Subnet",
		}
	})

	p.AddResourceConfigurator("tencentcloud_dnat", func(r *config.Resource) {
		r.ShortGroup = shortGroupVpc
		r.Kind = "Dnat"
		r.References["vpc_id"] = config.Reference{
			Type: "VPC",
		}
		r.References["nat_id"] = config.Reference{
			Type: "NatGateway",
		}
	})

	p.AddResourceConfigurator("tencentcloud_security_group", func(r *config.Resource) {
		r.ShortGroup = shortGroupVpc
		r.Kind = "SecurityGroup"
	})

	p.AddResourceConfigurator("tencentcloud_security_group_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroupVpc
		r.Kind = "SecurityGroupRule"
		r.References["security_group_id"] = config.Reference{
			Type: "SecurityGroup",
		}
	})

	p.AddResourceConfigurator("tencentcloud_security_group_lite_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroupVpc
		r.Kind = "SecurityGroupLiteRule"
		r.References["security_group_id"] = config.Reference{
			Type: "SecurityGroup",
		}
	})

	p.AddResourceConfigurator("tencentcloud_vpc_acl", func(r *config.Resource) {
		r.ShortGroup = shortGroupVpc
		r.Kind = "VPCAcl"
		r.References["vpc_id"] = config.Reference{
			Type: "VPC",
		}
	})

	p.AddResourceConfigurator("tencentcloud_vpc_acl_attachment", func(r *config.Resource) {
		r.ShortGroup = shortGroupVpc
		r.Kind = "VPVAclAttachment"
		r.References["acl_id"] = config.Reference{
			Type: "VPCAcl",
		}
		r.References["subnet_id"] = config.Reference{
			Type: "Subnet",
		}
	})

	p.AddResourceConfigurator("tencentcloud_vpc_bandwidth_package", func(r *config.Resource) {
		r.ShortGroup = shortGroupVpc
		r.Kind = "VPCBandwidthPackage"
	})

	p.AddResourceConfigurator("tencentcloud_vpc_bandwidth_package_attachment", func(r *config.Resource) {
		r.ShortGroup = shortGroupVpc
		r.Kind = "VPCBandwidthPackageAttachment"
		r.References["bandwidth_package_id"] = config.Reference{
			Type: "VPCBandwidthPackage",
		}

	})

	p.AddResourceConfigurator("tencentcloud_vpn_gateway", func(r *config.Resource) {
		r.ShortGroup = shortGroupVpc
		r.Kind = "VPNGateway"
		r.References["vpc_id"] = config.Reference{
			Type: "VPC",
		}
	})

	p.AddResourceConfigurator("tencentcloud_vpn_connection", func(r *config.Resource) {
		r.ShortGroup = shortGroupVpc
		r.Kind = "VPNConnection"
		r.References["vpc_id"] = config.Reference{
			Type: "VPC",
		}
		r.References["vpn_gateway_id"] = config.Reference{
			Type: "VPNGateway",
		}
		r.References["customer_gateway_id"] = config.Reference{
			Type: "VPNCustomerGateway",
		}
	})

	p.AddResourceConfigurator("tencentcloud_vpn_gateway_route", func(r *config.Resource) {
		r.ShortGroup = shortGroupVpc
		r.Kind = "VPNGatewayRoute"
		r.References["vpn_gateway_id"] = config.Reference{
			Type: "VPNGateway",
		}
	})

	p.AddResourceConfigurator("tencentcloud_vpn_ssl_server", func(r *config.Resource) {
		r.ShortGroup = shortGroupVpc
		r.Kind = "VPNSSLServer"
		r.References["vpn_gateway_id"] = config.Reference{
			Type: "VPNGateway",
		}
	})

	p.AddResourceConfigurator("tencentcloud_vpn_ssl_client", func(r *config.Resource) {
		r.ShortGroup = shortGroupVpc
		r.Kind = "VPNSSLClient"
	})

	p.AddResourceConfigurator("tencentcloud_vpn_customer_gateway", func(r *config.Resource) {
		r.ShortGroup = shortGroupVpc
		r.Kind = "VPNCustomerGateway"
	})

	p.AddResourceConfigurator("tencentcloud_address_template", func(r *config.Resource) {
		r.ShortGroup = shortGroupVpc
		r.Kind = "AddressTemplate"
	})

	p.AddResourceConfigurator("tencentcloud_protocol_template", func(r *config.Resource) {
		r.ShortGroup = shortGroupVpc
		r.Kind = "ProtocolTemplate"
	})

	p.AddResourceConfigurator("tencentcloud_address_template_group", func(r *config.Resource) {
		r.ShortGroup = shortGroupVpc
		r.Kind = "AddressTemplateGroup"
		r.References["template_ids"] = config.Reference{
			Type: "AddressTemplate",
		}
	})

	p.AddResourceConfigurator("tencentcloud_protocol_template_group", func(r *config.Resource) {
		r.ShortGroup = shortGroupVpc
		r.Kind = "ProtocolTemplateGroup"
		r.References["template_ids"] = config.Reference{
			Type: "ProtocolTemplate",
		}
	})

	p.AddResourceConfigurator("tencentcloud_eip_association", func(r *config.Resource) {
		r.ShortGroup = shortGroupVpc
		r.Kind = "EipAssociation"
		r.References["eip_id"] = config.Reference{
			Type: "Eip",
		}
	})
}
