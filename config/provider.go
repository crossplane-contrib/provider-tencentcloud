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

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	tjconfig "github.com/crossplane/terrajet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/crossplane-contrib/provider-tencentcloud/config/cvm"
	"github.com/crossplane-contrib/provider-tencentcloud/config/vpc"
)

const (
	resourcePrefix = "tencentcloud"
	modulePath     = "github.com/crossplane-contrib/provider-tencentcloud"
	rootGroup      = "tencentcloud.crossplane.io"
)

//go:embed schema.json
var providerSchema string

// IncludedResources include resource
var IncludedResources = []string{
	// VPC
	"tencentcloud_vpc$",
	"tencentcloud_vpc_acl$",
	"tencentcloud_vpc_acl_attachment$",
	"tencentcloud_subnet$",
	"tencentcloud_eip$",
	"tencentcloud_eip_association$",
	"tencentcloud_eni$",
	"tencentcloud_eni_attachment$",
	"tencentcloud_ha_vip$",
	"tencentcloud_security_group$",
	"tencentcloud_security_group_rule$",
	"tencentcloud_security_group_lite_rule$",
	"tencentcloud_address_template$",
	"tencentcloud_address_template_group$",
	"tencentcloud_protocol_template$",
	"tencentcloud_protocol_template_group$",
	"tencentcloud_route_table$",
	"tencentcloud_route_table_entry$",
	"tencentcloud_route_entry$",
	"tencentcloud_nat_gateway$",
	"tencentcloud_nat_gateway_snat$",
	"tencentcloud_dnat$",
	"tencentcloud_vpc_bandwidth_package$",
	"tencentcloud_vpc_bandwidth_package_attachment$",
	// VPN
	"tencentcloud_vpn_gateway$",
	"tencentcloud_vpn_ssl_client$",
	"tencentcloud_vpn_customer_gateway$",
	"tencentcloud_vpn_connection$",
	"tencentcloud_vpn_gateway_route$",
	"tencentcloud_vpn_ssl_server$",

	// cvm
	"tencentcloud_instance$",
	"tencentcloud_instance_set$",
	"tencentcloud_key_pair$",
	"tencentcloud_placement_group$",
	"tencentcloud_reserved_instance$",
	"tencentcloud_image$",
}

// skipList
var skipList []string

// GetProvider returns provider configuration
func GetProvider() *tjconfig.Provider {
	defaultResourceFn := func(name string, terraformResource *schema.Resource, opts ...tjconfig.ResourceOption) *tjconfig.Resource {
		r := tjconfig.DefaultResource(name, terraformResource)
		// Add any provider-specific defaulting here. For example:
		//   r.ExternalName = tjconfig.IdentifierFromProvider
		return r
	}

	pc := tjconfig.NewProviderWithSchema([]byte(providerSchema), resourcePrefix, modulePath,
		tjconfig.WithDefaultResourceFn(defaultResourceFn),
		tjconfig.WithIncludeList(IncludedResources),
		tjconfig.WithSkipList(skipList),
		tjconfig.WithRootGroup(rootGroup),
	)

	for _, configure := range []func(provider *tjconfig.Provider){
		// add custom config functions
		vpc.Configure,
		cvm.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
