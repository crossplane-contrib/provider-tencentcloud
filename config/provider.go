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

	"github.com/crossplane-contrib/provider-tencentcloud/config/apigateway"
	"github.com/crossplane-contrib/provider-tencentcloud/config/as"
	"github.com/crossplane-contrib/provider-tencentcloud/config/audit"
	"github.com/crossplane-contrib/provider-tencentcloud/config/cam"
	"github.com/crossplane-contrib/provider-tencentcloud/config/cat"
	"github.com/crossplane-contrib/provider-tencentcloud/config/cbs"
	"github.com/crossplane-contrib/provider-tencentcloud/config/ccn"
	"github.com/crossplane-contrib/provider-tencentcloud/config/cdh"
	"github.com/crossplane-contrib/provider-tencentcloud/config/cdn"
	"github.com/crossplane-contrib/provider-tencentcloud/config/cfs"
	"github.com/crossplane-contrib/provider-tencentcloud/config/clb"
	"github.com/crossplane-contrib/provider-tencentcloud/config/cls"
	"github.com/crossplane-contrib/provider-tencentcloud/config/containercluster"
	"github.com/crossplane-contrib/provider-tencentcloud/config/cos"
	"github.com/crossplane-contrib/provider-tencentcloud/config/cvm"
	"github.com/crossplane-contrib/provider-tencentcloud/config/cynosdb"
	"github.com/crossplane-contrib/provider-tencentcloud/config/dc"
	"github.com/crossplane-contrib/provider-tencentcloud/config/dcdb"
	"github.com/crossplane-contrib/provider-tencentcloud/config/dnspod"
	"github.com/crossplane-contrib/provider-tencentcloud/config/elasticsearch"
	"github.com/crossplane-contrib/provider-tencentcloud/config/emr"
	"github.com/crossplane-contrib/provider-tencentcloud/config/eni"
	"github.com/crossplane-contrib/provider-tencentcloud/config/gaap"
	"github.com/crossplane-contrib/provider-tencentcloud/config/kafka"
	"github.com/crossplane-contrib/provider-tencentcloud/config/kms"
	"github.com/crossplane-contrib/provider-tencentcloud/config/lighthouse"
	"github.com/crossplane-contrib/provider-tencentcloud/config/mariadb"
	"github.com/crossplane-contrib/provider-tencentcloud/config/mongodb"
	"github.com/crossplane-contrib/provider-tencentcloud/config/monitor"
	"github.com/crossplane-contrib/provider-tencentcloud/config/mysql"
	"github.com/crossplane-contrib/provider-tencentcloud/config/postgresql"
	"github.com/crossplane-contrib/provider-tencentcloud/config/privatedns"
	"github.com/crossplane-contrib/provider-tencentcloud/config/redis"
	"github.com/crossplane-contrib/provider-tencentcloud/config/scf"
	"github.com/crossplane-contrib/provider-tencentcloud/config/ses"
	"github.com/crossplane-contrib/provider-tencentcloud/config/sms"
	"github.com/crossplane-contrib/provider-tencentcloud/config/sqlserver"
	"github.com/crossplane-contrib/provider-tencentcloud/config/ssl"
	"github.com/crossplane-contrib/provider-tencentcloud/config/ssm"
	"github.com/crossplane-contrib/provider-tencentcloud/config/tcaplus"
	"github.com/crossplane-contrib/provider-tencentcloud/config/tcm"
	"github.com/crossplane-contrib/provider-tencentcloud/config/tcr"
	"github.com/crossplane-contrib/provider-tencentcloud/config/tdmq"
	"github.com/crossplane-contrib/provider-tencentcloud/config/tem"
	"github.com/crossplane-contrib/provider-tencentcloud/config/teo"
	"github.com/crossplane-contrib/provider-tencentcloud/config/tke"
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

	// audit
	"tencentcloud_audit$",

	// as
	"tencentcloud_as_scaling_config$",
	"tencentcloud_as_scaling_group$",
	"tencentcloud_as_attachment$",
	"tencentcloud_as_scaling_policy$",
	"tencentcloud_as_schedule$",
	"tencentcloud_as_lifecycle_hook$",
	"tencentcloud_as_notification$",

	// cdn
	"tencentcloud_cdn_domain$",
	"tencentcloud_cdn_url_purge$",
	"tencentcloud_cdn_url_push$",

	// cam
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

	// cbs
	"tencentcloud_cbs_storage$",
	"tencentcloud_cbs_storage_set$",
	"tencentcloud_cbs_storage_attachment$",
	"tencentcloud_cbs_snapshot$",
	"tencentcloud_cbs_snapshot_policy$",
	"tencentcloud_cbs_snapshot_policy_attachment$",

	// ccn
	"tencentcloud_ccn$",
	"tencentcloud_ccn_attachment$",
	"tencentcloud_ccn_bandwidth_limit$",

	// cdh
	"tencentcloud_cdh_instance$",

	// cfs
	"tencentcloud_cfs_file_system$",
	"tencentcloud_cfs_access_group$",
	"tencentcloud_cfs_access_rule$",

	// container_cluster
	"tencentcloud_container_cluster$",
	"tencentcloud_container_cluster_instance$",

	// clb
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

	// cos
	"tencentcloud_cos_bucket$",
	"tencentcloud_cos_bucket_policy$",
	"tencentcloud_cos_bucket_object$",
	"tencentcloud_cos_bucket_domain_certificate_attachment$",

	// cynosdb
	"tencentcloud_cynosdb_cluster$",
	"tencentcloud_cynosdb_readonly_instance$",

	// ssl
	"tencentcloud_ssl_certificate$",
	"tencentcloud_ssl_free_certificate$",
	"tencentcloud_ssl_pay_certificate$",

	// tcr
	"tencentcloud_tcr_instance$",
	"tencentcloud_tcr_namespace$",
	"tencentcloud_tcr_repository$",
	"tencentcloud_tcr_token$",
	"tencentcloud_tcr_vpc_attachment$",

	// vod
	"tencentcloud_vod_adaptive_dynamic_streaming_template$",
	"tencentcloud_vod_procedure_template$",
	"tencentcloud_vod_snapshot_by_time_offset_template$",
	"tencentcloud_vod_image_sprite_template$",
	"tencentcloud_vod_super_player_config$",
	"tencentcloud_vod_sub_application$",

	// private dns
	"tencentcloud_private_dns_record$",
	"tencentcloud_private_dns_zone$",

	// eni
	"tencentcloud_eni$",
	"tencentcloud_eni_attachment$",

	// mariadb
	"tencentcloud_mariadb_hour_db_instance$",
	"tencentcloud_mariadb_dedicatedcluster_db_instance$",
	"tencentcloud_mariadb_security_groups$",
	"tencentcloud_mariadb_parameters$",
	"tencentcloud_mariadb_log_file_retention_period$",
	"tencentcloud_mariadb_account$",

	// cat
	"tencentcloud_cat_task_set$",

	// sms
	"tencentcloud_sms_sign$",
	"tencentcloud_sms_template$",

	// dcdb
	"tencentcloud_dcdb_hourdb_instance$",
	"tencentcloud_dcdb_account$",
	"tencentcloud_dcdb_security_group_attachment$",

	// ses
	"tencentcloud_ses_domain$",
	"tencentcloud_ses_email_address$",
	"tencentcloud_ses_template$",

	// tcm
	"tencentcloud_tcm_mesh$",
	"tencentcloud_tcm_cluster_attachment$",

	// teo
	"tencentcloud_teo_zone$",
	"tencentcloud_teo_zone_setting$",
	"tencentcloud_teo_security_policy$",
	"tencentcloud_teo_rule_engine$",
	"tencentcloud_teo_rule_engine_priority$",
	"tencentcloud_teo_origin_group$",
	"tencentcloud_teo_load_balancing$",
	"tencentcloud_teo_dns_sec$",
	"tencentcloud_teo_dns_record$",
	"tencentcloud_teo_ddos_policy$",
	"tencentcloud_teo_custom_error_page$",
	"tencentcloud_teo_application_proxy$",
	"tencentcloud_teo_application_proxy_rule$",

	// tem
	"tencentcloud_tem_environment$",
	"tencentcloud_tem_app_config$",
	"tencentcloud_tem_application$",
	"tencentcloud_tem_gateway$",
	"tencentcloud_tem_workload$",
	"tencentcloud_tem_scale_rule$",
	"tencentcloud_tem_log_config$",

	// lighthouse
	"tencentcloud_lighthouse_instance$",

	// cls
	"tencentcloud_cls_logset$",
	"tencentcloud_cls_topic$",
	"tencentcloud_cls_index$",
	"tencentcloud_cls_cos_shipper$",
	"tencentcloud_cls_config$",
	"tencentcloud_cls_config_extra$",
	"tencentcloud_cls_machine_group$",
	"tencentcloud_cls_config_attachment$",

	// dnspod
	"tencentcloud_dnspod_record$",
	"tencentcloud_dnspod_domain_instance$",

	// emr
	"tencentcloud_emr_cluster$",

	// tcaplus
	"tencentcloud_tcaplus_cluster$",
	"tencentcloud_tcaplus_tablegroup$",
	"tencentcloud_tcaplus_table$",
	"tencentcloud_tcaplus_idl$",

	// ssm
	"tencentcloud_ssm_secret$",
	"tencentcloud_ssm_secret_version$",

	// sqlserver
	"tencentcloud_sqlserver_basic_instance$",
	"tencentcloud_sqlserver_instance$",
	"tencentcloud_sqlserver_db$",
	"tencentcloud_sqlserver_readonly_instance$",
	"tencentcloud_sqlserver_account$",
	"tencentcloud_sqlserver_account_db_attachment$",
	"tencentcloud_sqlserver_publish_subscribe$",

	// scf
	"tencentcloud_scf_function$",
	"tencentcloud_scf_namespace$",
	"tencentcloud_scf_layer$",

	// redis
	"tencentcloud_redis_instance$",
	"tencentcloud_redis_backup_config$",

	// postgresql
	"tencentcloud_postgresql_instance$",
	"tencentcloud_postgresql_readonly_instance$",
	"tencentcloud_postgresql_readonly_group$",
	"tencentcloud_postgresql_readonly_attachment$",

	// monitor
	"tencentcloud_monitor_alarm_notice$",
	"tencentcloud_monitor_alarm_policy$",
	"tencentcloud_monitor_policy_group$",
	"tencentcloud_monitor_binding_receiver$",
	"tencentcloud_monitor_grafana_instance$",
	"tencentcloud_monitor_grafana_integration$",
	"tencentcloud_monitor_grafana_notification_channel$",
	"tencentcloud_monitor_grafana_plugin$",
	"tencentcloud_monitor_grafana_sso_account$",
	"tencentcloud_monitor_policy_binding_object$",
	"tencentcloud_monitor_policy_group$",
	"tencentcloud_monitor_tmp_instance$",
	"tencentcloud_monitor_tmp_alert_rule$",
	"tencentcloud_monitor_tmp_cvm_agent$",
	"tencentcloud_monitor_tmp_exporter_integration$",
	"tencentcloud_monitor_tmp_recording_rule$",
	"tencentcloud_monitor_tmp_scrape_job$",
	"tencentcloud_monitor_tmp_tke_alert_policy$",
	"tencentcloud_monitor_tmp_tke_cluster_agent$",
	"tencentcloud_monitor_tmp_tke_config$",
	"tencentcloud_monitor_tmp_tke_global_notification$",
	"tencentcloud_monitor_tmp_tke_record_rule_yaml$",
	"tencentcloud_monitor_tmp_tke_template$",
	"tencentcloud_monitor_tmp_tke_template_attachment$",

	// mysql
	"tencentcloud_mysql_instance$",
	"tencentcloud_mysql_account$",
	"tencentcloud_mysql_account_privilege$",
	"tencentcloud_mysql_backup_policy$",
	"tencentcloud_mysql_privilege$",
	"tencentcloud_mysql_readonly_instance$",

	// mongodb
	"tencentcloud_mongodb_instance$",
	"tencentcloud_mongodb_sharding_instance$",
	"tencentcloud_mongodb_standby_instance$",

	// tdmq
	"tencentcloud_tdmq_instance$",
	"tencentcloud_tdmq_namespace$",
	"tencentcloud_tdmq_topic$",
	"tencentcloud_tdmq_role$",
	"tencentcloud_tdmq_namespace_role_attachment$",

	// tke
	"tencentcloud_kubernetes_cluster$",
	"tencentcloud_kubernetes_scale_worker$",
	"tencentcloud_kubernetes_cluster_attachment$",
	"tencentcloud_kubernetes_node_pool$",
	"tencentcloud_kubernetes_addon_attachment$",
	"tencentcloud_kubernetes_auth_attachment$",
	"tencentcloud_kubernetes_cluster_endpoint$",

	// kms
	"tencentcloud_kms_key$",
	"tencentcloud_kms_external_key$",

	// gaap
	"tencentcloud_gaap_proxy$",
	"tencentcloud_gaap_certificate$",
	"tencentcloud_gaap_layer4_listener$",
	"tencentcloud_gaap_layer7_listener$",
	"tencentcloud_gaap_realserver$",
	"tencentcloud_gaap_http_domain$",
	"tencentcloud_gaap_http_rule$",
	"tencentcloud_gaap_certificate$",
	"tencentcloud_gaap_security_policy$",
	"tencentcloud_gaap_security_rule$",

	// elasticsearch
	"tencentcloud_elasticsearch_instance$",

	// dc
	"tencentcloud_dcx$",
	"tencentcloud_dc_gateway$",
	"tencentcloud_dc_gateway_ccn_route$",

	// kafka
	"tencentcloud_ckafka_instance$",
	"tencentcloud_ckafka_user$",
	"tencentcloud_ckafka_topic$",
	"tencentcloud_ckafka_acl$",

	// api gateway
	"tencentcloud_api_gateway_service$",
	"tencentcloud_api_gateway_api$",
	"tencentcloud_api_gateway_api_key$",
	"tencentcloud_api_gateway_usage_plan$",
	"tencentcloud_api_gateway_usage_plan_attachment$",
	"tencentcloud_api_gateway_ip_strategy$",
	"tencentcloud_api_gateway_custom_domain$",
	"tencentcloud_api_gateway_api_key_attachment$",
	"tencentcloud_api_gateway_service_release$",
	"tencentcloud_api_gateway_strategy_attachment$",

	// dayu
	"tencentcloud_dayu_eip$",
	"tencentcloud_dayu_l4_rule$",
	"tencentcloud_dayu_l4_rule_v2$",
	"tencentcloud_dayu_l7_rule$",
	"tencentcloud_dayu_l7_rule_v2$",
	"tencentcloud_dayu_ddos_policy$",
	"tencentcloud_dayu_ddos_policy_v2$",
	"tencentcloud_dayu_ddos_policy_case$",
	"tencentcloud_dayu_ddos_policy_attachment$",
	"tencentcloud_dayu_cc_policy_v2$",
	"tencentcloud_dayu_cc_https_policy$",
	"tencentcloud_dayu_cc_http_policy$",
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
		containercluster.Configure,
		eni.Configure,
		mariadb.Configure,
		cat.Configure,
		sms.Configure,
		dcdb.Configure,
		ses.Configure,
		tcm.Configure,
		teo.Configure,
		tem.Configure,
		lighthouse.Configure,
		cls.Configure,
		dnspod.Configure,
		emr.Configure,
		tcaplus.Configure,
		ssm.Configure,
		sqlserver.Configure,
		scf.Configure,
		redis.Configure,
		postgresql.Configure,
		monitor.Configure,
		mysql.Configure,
		mongodb.Configure,
		tdmq.Configure,
		tke.Configure,
		kms.Configure,
		gaap.Configure,
		elasticsearch.Configure,
		dc.Configure,
		kafka.Configure,
		apigateway.Configure,
		gaap.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
