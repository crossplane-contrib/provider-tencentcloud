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

	"github.com/crossplane-contrib/provider-tencentcloud/config/as"
	"github.com/crossplane-contrib/provider-tencentcloud/config/audit"
	"github.com/crossplane-contrib/provider-tencentcloud/config/cam"
	"github.com/crossplane-contrib/provider-tencentcloud/config/cbs"
	"github.com/crossplane-contrib/provider-tencentcloud/config/ccn"
	"github.com/crossplane-contrib/provider-tencentcloud/config/cdh"
	"github.com/crossplane-contrib/provider-tencentcloud/config/cdn"
	"github.com/crossplane-contrib/provider-tencentcloud/config/cfs"
	"github.com/crossplane-contrib/provider-tencentcloud/config/clb"
	"github.com/crossplane-contrib/provider-tencentcloud/config/cos"
	"github.com/crossplane-contrib/provider-tencentcloud/config/cvm"
	"github.com/crossplane-contrib/provider-tencentcloud/config/cynosdb"
	"github.com/crossplane-contrib/provider-tencentcloud/config/privatedns"
	"github.com/crossplane-contrib/provider-tencentcloud/config/ssl"
	"github.com/crossplane-contrib/provider-tencentcloud/config/tcr"
	"github.com/crossplane-contrib/provider-tencentcloud/config/vod"
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

	//audit
	"tencentcloud_audit$",

	//as
	"tencentcloud_as_scaling_config$",
	"tencentcloud_as_scaling_group$",
	"tencentcloud_as_attachment$",
	"tencentcloud_as_scaling_policy$",
	"tencentcloud_as_schedule$",
	"tencentcloud_as_lifecycle_hook$",
	"tencentcloud_as_notification$",

	//cdn
	"tencentcloud_cdn_domain$",
	"tencentcloud_cdn_url_purge$",
	"tencentcloud_cdn_url_push$",

	//cam
	"tencentcloud_cam_role$",
	"tencentcloud_cam_role_sso$",
	"tencentcloud_cam_policy$",
	"tencentcloud_cam_role_policy_attachment$",
	"tencentcloud_cam_oidc_sso$",
	"tencentcloud_cam_group$",
	"tencentcloud_cam_user$",
	"tencentcloud_cam_group_membership$",
	"tencentcloud_cam_group_policy_attachment$",
	"tencentcloud_cam_user_policy_attachment$",
	"tencentcloud_cam_saml_provider$",

	//cbs
	"tencentcloud_cbs_storage$",
	"tencentcloud_cbs_storage_set$",
	"tencentcloud_cbs_storage_attachment$",
	"tencentcloud_cbs_snapshot$",
	"tencentcloud_cbs_snapshot_policy$",
	"tencentcloud_cbs_snapshot_policy_attachment$",

	//ccn
	"tencentcloud_ccn$",
	"tencentcloud_ccn_attachment$",
	"tencentcloud_ccn_bandwidth_limit$",

	//cdh
	"tencentcloud_cdh_instance$",

	//cfs
	"tencentcloud_cfs_file_system$",
	"tencentcloud_cfs_access_group$",
	"tencentcloud_cfs_access_rule$",

	//container_cluster
	"tencentcloud_container_cluster$",
	"tencentcloud_container_cluster_instance$",

	//clb
	"tencentcloud_clb_instance$",
	"tencentcloud_clb_listener$",
	"tencentcloud_clb_listener_rule$",
	"tencentcloud_clb_attachment$",
	"tencentcloud_clb_customized_config$",
	"tencentcloud_clb_log_set$",
	"tencentcloud_clb_log_topic$",
	"tencentcloud_clb_redirection$",
	"tencentcloud_clb_snat_ip$",
	"tencentcloud_clb_target_group",
	"tencentcloud_clb_target_group_attachment$",
	"tencentcloud_clb_target_group_instance_attachment$",
	"tencentcloud_lb$",
	"tencentcloud_alb_server_attachment$",

	//cos
	"tencentcloud_cos_bucket$",
	"tencentcloud_cos_bucket_policy$",
	"tencentcloud_cos_bucket_object$",
	"tencentcloud_cos_bucket_domain_certificate_attachment$",

	//cynosdb
	"tencentcloud_cynosdb_cluster$",
	"tencentcloud_cynosdb_readonly_instance$",

	//ssl
	"tencentcloud_ssl_certificate$",
	"tencentcloud_ssl_free_certificate$",
	"tencentcloud_ssl_pay_certificate$",

	//tcr
	"tencentcloud_tcr_instance$",
	"tencentcloud_tcr_namespace$",
	"tencentcloud_tcr_repository$",
	"tencentcloud_tcr_token$",
	"tencentcloud_tcr_vpc_attachment$",

	//vod
	"tencentcloud_vod_adaptive_dynamic_streaming_template$",
	"tencentcloud_vod_procedure_template$",
	"tencentcloud_vod_snapshot_by_time_offset_template$",
	"tencentcloud_vod_image_sprite_template$",
	"tencentcloud_vod_super_player_config$",
	"tencentcloud_vod_sub_application$",

	//private dns
	"tencentcloud_private_dns_record$",
	"tencentcloud_private_dns_zone$",
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
		audit.Configure,
		as.Configure,
		cdn.Configure,
		cam.Configure,
		cbs.Configure,
		ccn.Configure,
		cdh.Configure,
		cfs.Configure,
		clb.Configure,
		cos.Configure,
		cynosdb.Configure,
		ssl.Configure,
		tcr.Configure,
		vod.Configure,
		privatedns.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
