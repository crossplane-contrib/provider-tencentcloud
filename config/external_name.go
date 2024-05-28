/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	// api gateway
	"tencentcloud_api_gateway_service":               config.IdentifierFromProvider,
	"tencentcloud_api_gateway_api":                   config.IdentifierFromProvider,
	"tencentcloud_api_gateway_api_key":               config.IdentifierFromProvider,
	"tencentcloud_api_gateway_usage_plan":            config.IdentifierFromProvider,
	"tencentcloud_api_gateway_usage_plan_attachment": config.IdentifierFromProvider,
	"tencentcloud_api_gateway_ip_strategy":           config.IdentifierFromProvider,
	"tencentcloud_api_gateway_custom_domain":         config.IdentifierFromProvider,
	"tencentcloud_api_gateway_api_key_attachment":    config.IdentifierFromProvider,
	"tencentcloud_api_gateway_service_release":       config.IdentifierFromProvider,
	"tencentcloud_api_gateway_strategy_attachment":   config.IdentifierFromProvider,
	// as
	"tencentcloud_as_scaling_config": config.IdentifierFromProvider,
	"tencentcloud_as_scaling_group":  config.IdentifierFromProvider,
	"tencentcloud_as_attachment":     config.IdentifierFromProvider,
	"tencentcloud_as_scaling_policy": config.IdentifierFromProvider,
	"tencentcloud_as_schedule":       config.IdentifierFromProvider,
	"tencentcloud_as_lifecycle_hook": config.IdentifierFromProvider,
	"tencentcloud_as_notification":   config.IdentifierFromProvider,
	// audit
	"tencentcloud_audit": config.IdentifierFromProvider,
	// cam
	"tencentcloud_cam_role_by_name":                   config.NameAsIdentifier,
	"tencentcloud_cam_role_sso":                       config.IdentifierFromProvider,
	"tencentcloud_cam_policy_by_name":                 config.NameAsIdentifier,
	"tencentcloud_cam_role_policy_attachment_by_name": config.IdentifierFromProvider,
	"tencentcloud_cam_oidc_sso":                       config.IdentifierFromProvider,
	"tencentcloud_cam_group":                          config.IdentifierFromProvider,
	"tencentcloud_cam_user":                           config.IdentifierFromProvider,
	"tencentcloud_cam_group_membership":               config.IdentifierFromProvider,
	"tencentcloud_cam_group_policy_attachment":        config.IdentifierFromProvider,
	"tencentcloud_cam_user_policy_attachment":         config.IdentifierFromProvider,
	"tencentcloud_cam_saml_provider":                  config.IdentifierFromProvider,
	// cat
	"tencentcloud_cat_task_set": config.IdentifierFromProvider,
	// cbs
	"tencentcloud_cbs_storage":                    config.IdentifierFromProvider,
	"tencentcloud_cbs_storage_set":                config.IdentifierFromProvider,
	"tencentcloud_cbs_storage_attachment":         config.IdentifierFromProvider,
	"tencentcloud_cbs_snapshot":                   config.IdentifierFromProvider,
	"tencentcloud_cbs_snapshot_policy":            config.IdentifierFromProvider,
	"tencentcloud_cbs_snapshot_policy_attachment": config.IdentifierFromProvider,
	// ccn
	"tencentcloud_ccn":                 config.IdentifierFromProvider,
	"tencentcloud_ccn_attachment":      config.IdentifierFromProvider,
	"tencentcloud_ccn_bandwidth_limit": config.IdentifierFromProvider,
	// cdh
	"tencentcloud_cdh_instance": config.IdentifierFromProvider,
	// cfs
	"tencentcloud_cfs_file_system":  config.IdentifierFromProvider,
	"tencentcloud_cfs_access_group": config.IdentifierFromProvider,
	"tencentcloud_cfs_access_rule":  config.IdentifierFromProvider,
	// clb
	"tencentcloud_clb_instance":                         config.IdentifierFromProvider,
	"tencentcloud_clb_listener":                         config.IdentifierFromProvider,
	"tencentcloud_clb_listener_rule":                    config.IdentifierFromProvider,
	"tencentcloud_clb_attachment":                       config.IdentifierFromProvider,
	"tencentcloud_clb_customized_config":                config.IdentifierFromProvider,
	"tencentcloud_clb_log_set":                          config.IdentifierFromProvider,
	"tencentcloud_clb_log_topic":                        config.IdentifierFromProvider,
	"tencentcloud_clb_redirection":                      config.IdentifierFromProvider,
	"tencentcloud_clb_snat_ip":                          config.IdentifierFromProvider,
	"tencentcloud_clb_target_group":                     config.IdentifierFromProvider,
	"tencentcloud_clb_target_group_attachment":          config.IdentifierFromProvider,
	"tencentcloud_clb_target_group_instance_attachment": config.IdentifierFromProvider,
	"tencentcloud_lb":                                   config.IdentifierFromProvider,
	"tencentcloud_alb_server_attachment":                config.IdentifierFromProvider,
	// cls
	"tencentcloud_cls_logset":            config.IdentifierFromProvider,
	"tencentcloud_cls_topic":             config.IdentifierFromProvider,
	"tencentcloud_cls_index":             config.IdentifierFromProvider,
	"tencentcloud_cls_cos_shipper":       config.IdentifierFromProvider,
	"tencentcloud_cls_config":            config.IdentifierFromProvider,
	"tencentcloud_cls_config_extra":      config.IdentifierFromProvider,
	"tencentcloud_cls_machine_group":     config.IdentifierFromProvider,
	"tencentcloud_cls_config_attachment": config.IdentifierFromProvider,
	// cos
	"tencentcloud_cos_bucket":                               config.IdentifierFromProvider,
	"tencentcloud_cos_bucket_policy":                        config.IdentifierFromProvider,
	"tencentcloud_cos_bucket_object":                        config.IdentifierFromProvider,
	"tencentcloud_cos_bucket_domain_certificate_attachment": config.IdentifierFromProvider,
	// cvm
	"tencentcloud_instance":          config.IdentifierFromProvider,
	"tencentcloud_instance_set":      config.IdentifierFromProvider,
	"tencentcloud_key_pair":          config.IdentifierFromProvider,
	"tencentcloud_placement_group":   config.IdentifierFromProvider,
	"tencentcloud_reserved_instance": config.IdentifierFromProvider,
	"tencentcloud_image":             config.IdentifierFromProvider,
	// cynosdb
	"tencentcloud_cynosdb_cluster":           config.IdentifierFromProvider,
	"tencentcloud_cynosdb_readonly_instance": config.IdentifierFromProvider,
	// dayu
	"tencentcloud_dayu_eip":                    config.IdentifierFromProvider,
	"tencentcloud_dayu_l4_rule":                config.IdentifierFromProvider,
	"tencentcloud_dayu_l4_rule_v2":             config.IdentifierFromProvider,
	"tencentcloud_dayu_l7_rule":                config.IdentifierFromProvider,
	"tencentcloud_dayu_l7_rule_v2":             config.IdentifierFromProvider,
	"tencentcloud_dayu_ddos_policy":            config.IdentifierFromProvider,
	"tencentcloud_dayu_ddos_policy_v2":         config.IdentifierFromProvider,
	"tencentcloud_dayu_ddos_policy_case":       config.IdentifierFromProvider,
	"tencentcloud_dayu_ddos_policy_attachment": config.IdentifierFromProvider,
	"tencentcloud_dayu_cc_policy_v2":           config.IdentifierFromProvider,
	"tencentcloud_dayu_cc_https_policy":        config.IdentifierFromProvider,
	"tencentcloud_dayu_cc_http_policy":         config.IdentifierFromProvider,
	// dc
	"tencentcloud_dcx":                  config.IdentifierFromProvider,
	"tencentcloud_dc_gateway":           config.IdentifierFromProvider,
	"tencentcloud_dc_gateway_ccn_route": config.IdentifierFromProvider,
	// dcdb
	"tencentcloud_dcdb_hourdb_instance":           config.IdentifierFromProvider,
	"tencentcloud_dcdb_account":                   config.IdentifierFromProvider,
	"tencentcloud_dcdb_security_group_attachment": config.IdentifierFromProvider,
	// dnspod
	"tencentcloud_dnspod_record":          config.IdentifierFromProvider,
	"tencentcloud_dnspod_domain_instance": config.IdentifierFromProvider,
	// elasticsearch
	"tencentcloud_elasticsearch_instance": config.IdentifierFromProvider,
	// emr
	"tencentcloud_emr_cluster": config.IdentifierFromProvider,
	// eni
	"tencentcloud_eni":            config.IdentifierFromProvider,
	"tencentcloud_eni_attachment": config.IdentifierFromProvider,
	// gaap
	"tencentcloud_gaap_proxy":           config.IdentifierFromProvider,
	"tencentcloud_gaap_certificate":     config.IdentifierFromProvider,
	"tencentcloud_gaap_layer4_listener": config.IdentifierFromProvider,
	"tencentcloud_gaap_layer7_listener": config.IdentifierFromProvider,
	"tencentcloud_gaap_realserver":      config.IdentifierFromProvider,
	"tencentcloud_gaap_http_domain":     config.IdentifierFromProvider,
	"tencentcloud_gaap_http_rule":       config.IdentifierFromProvider,
	"tencentcloud_gaap_security_policy": config.IdentifierFromProvider,
	"tencentcloud_gaap_security_rule":   config.IdentifierFromProvider,
	// kafka
	"tencentcloud_ckafka_instance": config.IdentifierFromProvider,
	"tencentcloud_ckafka_user":     config.IdentifierFromProvider,
	"tencentcloud_ckafka_topic":    config.IdentifierFromProvider,
	"tencentcloud_ckafka_acl":      config.IdentifierFromProvider,
	// kms
	"tencentcloud_kms_key":          config.IdentifierFromProvider,
	"tencentcloud_kms_external_key": config.IdentifierFromProvider,
	// lighthouse
	"tencentcloud_lighthouse_instance": config.IdentifierFromProvider,
	// mariadb
	"tencentcloud_mariadb_hour_db_instance":             config.IdentifierFromProvider,
	"tencentcloud_mariadb_dedicatedcluster_db_instance": config.IdentifierFromProvider,
	"tencentcloud_mariadb_security_groups":              config.IdentifierFromProvider,
	"tencentcloud_mariadb_parameters":                   config.IdentifierFromProvider,
	"tencentcloud_mariadb_log_file_retention_period":    config.IdentifierFromProvider,
	"tencentcloud_mariadb_account":                      config.IdentifierFromProvider,
	// mongodb
	"tencentcloud_mongodb_instance":          config.IdentifierFromProvider,
	"tencentcloud_mongodb_sharding_instance": config.IdentifierFromProvider,
	"tencentcloud_mongodb_standby_instance":  config.IdentifierFromProvider,
	// monitor
	"tencentcloud_monitor_alarm_notice":                 config.IdentifierFromProvider,
	"tencentcloud_monitor_alarm_policy":                 config.IdentifierFromProvider,
	"tencentcloud_monitor_policy_group":                 config.IdentifierFromProvider,
	"tencentcloud_monitor_binding_receiver":             config.IdentifierFromProvider,
	"tencentcloud_monitor_grafana_instance":             config.IdentifierFromProvider,
	"tencentcloud_monitor_grafana_integration":          config.IdentifierFromProvider,
	"tencentcloud_monitor_grafana_notification_channel": config.IdentifierFromProvider,
	"tencentcloud_monitor_grafana_plugin":               config.IdentifierFromProvider,
	"tencentcloud_monitor_grafana_sso_account":          config.IdentifierFromProvider,
	"tencentcloud_monitor_policy_binding_object":        config.IdentifierFromProvider,
	"tencentcloud_monitor_tmp_instance":                 config.IdentifierFromProvider,
	"tencentcloud_monitor_tmp_alert_rule":               config.IdentifierFromProvider,
	"tencentcloud_monitor_tmp_cvm_agent":                config.IdentifierFromProvider,
	"tencentcloud_monitor_tmp_exporter_integration":     config.IdentifierFromProvider,
	"tencentcloud_monitor_tmp_recording_rule":           config.IdentifierFromProvider,
	"tencentcloud_monitor_tmp_scrape_job":               config.IdentifierFromProvider,
	"tencentcloud_monitor_tmp_tke_alert_policy":         config.IdentifierFromProvider,
	"tencentcloud_monitor_tmp_tke_cluster_agent":        config.IdentifierFromProvider,
	"tencentcloud_monitor_tmp_tke_config":               config.IdentifierFromProvider,
	"tencentcloud_monitor_tmp_tke_global_notification":  config.IdentifierFromProvider,
	"tencentcloud_monitor_tmp_tke_record_rule_yaml":     config.IdentifierFromProvider,
	"tencentcloud_monitor_tmp_tke_template":             config.IdentifierFromProvider,
	"tencentcloud_monitor_tmp_tke_template_attachment":  config.IdentifierFromProvider,
	// postgresql
	"tencentcloud_postgresql_instance":            config.IdentifierFromProvider,
	"tencentcloud_postgresql_readonly_instance":   config.IdentifierFromProvider,
	"tencentcloud_postgresql_readonly_group":      config.IdentifierFromProvider,
	"tencentcloud_postgresql_readonly_attachment": config.IdentifierFromProvider,
	// private dns
	"tencentcloud_private_dns_record": config.IdentifierFromProvider,
	"tencentcloud_private_dns_zone":   config.IdentifierFromProvider,
	// redis
	"tencentcloud_redis_instance":      config.IdentifierFromProvider,
	"tencentcloud_redis_backup_config": config.IdentifierFromProvider,
	// scf
	"tencentcloud_scf_function":  config.IdentifierFromProvider,
	"tencentcloud_scf_namespace": config.IdentifierFromProvider,
	"tencentcloud_scf_layer":     config.IdentifierFromProvider,
	// ses
	"tencentcloud_ses_domain":        config.IdentifierFromProvider,
	"tencentcloud_ses_email_address": config.IdentifierFromProvider,
	"tencentcloud_ses_template":      config.IdentifierFromProvider,
	// sms
	"tencentcloud_sms_sign":     config.IdentifierFromProvider,
	"tencentcloud_sms_template": config.IdentifierFromProvider,
	// sqlserver
	"tencentcloud_sqlserver_basic_instance":        config.IdentifierFromProvider,
	"tencentcloud_sqlserver_instance":              config.IdentifierFromProvider,
	"tencentcloud_sqlserver_db":                    config.IdentifierFromProvider,
	"tencentcloud_sqlserver_readonly_instance":     config.IdentifierFromProvider,
	"tencentcloud_sqlserver_account":               config.IdentifierFromProvider,
	"tencentcloud_sqlserver_account_db_attachment": config.IdentifierFromProvider,
	"tencentcloud_sqlserver_publish_subscribe":     config.IdentifierFromProvider,
	// ssl
	"tencentcloud_ssl_certificate":      config.IdentifierFromProvider,
	"tencentcloud_ssl_free_certificate": config.IdentifierFromProvider,
	"tencentcloud_ssl_pay_certificate":  config.IdentifierFromProvider,
	// ssm
	"tencentcloud_ssm_secret":         config.IdentifierFromProvider,
	"tencentcloud_ssm_secret_version": config.IdentifierFromProvider,
	// tcaplus
	"tencentcloud_tcaplus_cluster":    config.IdentifierFromProvider,
	"tencentcloud_tcaplus_tablegroup": config.IdentifierFromProvider,
	"tencentcloud_tcaplus_table":      config.IdentifierFromProvider,
	"tencentcloud_tcaplus_idl":        config.IdentifierFromProvider,
	// tcm
	"tencentcloud_tcm_mesh":               config.IdentifierFromProvider,
	"tencentcloud_tcm_cluster_attachment": config.IdentifierFromProvider,
	// tcmq
	"tencentcloud_tcmq_queue":     config.IdentifierFromProvider,
	"tencentcloud_tcmq_subscribe": config.IdentifierFromProvider,
	"tencentcloud_tcmq_topic":     config.IdentifierFromProvider,
	// tcr
	"tencentcloud_tcr_instance":       config.IdentifierFromProvider,
	"tencentcloud_tcr_namespace":      config.IdentifierFromProvider,
	"tencentcloud_tcr_repository":     config.IdentifierFromProvider,
	"tencentcloud_tcr_token":          config.IdentifierFromProvider,
	"tencentcloud_tcr_vpc_attachment": config.IdentifierFromProvider,
	// tdmq
	"tencentcloud_tdmq_instance":                  config.IdentifierFromProvider,
	"tencentcloud_tdmq_namespace":                 config.IdentifierFromProvider,
	"tencentcloud_tdmq_topic":                     config.IdentifierFromProvider,
	"tencentcloud_tdmq_role":                      config.IdentifierFromProvider,
	"tencentcloud_tdmq_namespace_role_attachment": config.IdentifierFromProvider,
	// tem
	"tencentcloud_tem_environment": config.IdentifierFromProvider,
	"tencentcloud_tem_app_config":  config.IdentifierFromProvider,
	"tencentcloud_tem_application": config.IdentifierFromProvider,
	"tencentcloud_tem_gateway":     config.IdentifierFromProvider,
	"tencentcloud_tem_workload":    config.IdentifierFromProvider,
	"tencentcloud_tem_scale_rule":  config.IdentifierFromProvider,
	"tencentcloud_tem_log_config":  config.IdentifierFromProvider,
	// teo
	"tencentcloud_teo_zone":                   config.IdentifierFromProvider,
	"tencentcloud_teo_zone_setting":           config.IdentifierFromProvider,
	"tencentcloud_teo_rule_engine":            config.IdentifierFromProvider,
	"tencentcloud_teo_origin_group":           config.IdentifierFromProvider,
	"tencentcloud_teo_application_proxy_rule": config.IdentifierFromProvider,
	// tke
	"tencentcloud_kubernetes_cluster":            config.IdentifierFromProvider,
	"tencentcloud_kubernetes_scale_worker":       config.IdentifierFromProvider,
	"tencentcloud_kubernetes_cluster_attachment": config.IdentifierFromProvider,
	"tencentcloud_kubernetes_node_pool":          config.IdentifierFromProvider,
	"tencentcloud_kubernetes_addon_attachment":   config.IdentifierFromProvider,
	"tencentcloud_kubernetes_auth_attachment":    config.IdentifierFromProvider,
	"tencentcloud_kubernetes_cluster_endpoint":   config.IdentifierFromProvider,
	// vod
	"tencentcloud_vod_adaptive_dynamic_streaming_template": config.IdentifierFromProvider,
	"tencentcloud_vod_procedure_template":                  config.IdentifierFromProvider,
	"tencentcloud_vod_snapshot_by_time_offset_template":    config.IdentifierFromProvider,
	"tencentcloud_vod_image_sprite_template":               config.IdentifierFromProvider,
	"tencentcloud_vod_super_player_config":                 config.IdentifierFromProvider,
	"tencentcloud_vod_sub_application":                     config.IdentifierFromProvider,
	// vpc
	"tencentcloud_vpc":                              config.IdentifierFromProvider,
	"tencentcloud_vpc_acl":                          config.IdentifierFromProvider,
	"tencentcloud_vpc_acl_attachment":               config.IdentifierFromProvider,
	"tencentcloud_subnet":                           config.IdentifierFromProvider,
	"tencentcloud_eip":                              config.IdentifierFromProvider,
	"tencentcloud_eip_association":                  config.IdentifierFromProvider,
	"tencentcloud_ha_vip":                           config.IdentifierFromProvider,
	"tencentcloud_security_group":                   config.IdentifierFromProvider,
	"tencentcloud_security_group_rule":              config.IdentifierFromProvider,
	"tencentcloud_security_group_lite_rule":         config.IdentifierFromProvider,
	"tencentcloud_address_template":                 config.IdentifierFromProvider,
	"tencentcloud_address_template_group":           config.IdentifierFromProvider,
	"tencentcloud_protocol_template":                config.IdentifierFromProvider,
	"tencentcloud_protocol_template_group":          config.IdentifierFromProvider,
	"tencentcloud_route_table":                      config.IdentifierFromProvider,
	"tencentcloud_route_table_entry":                config.IdentifierFromProvider,
	"tencentcloud_route_entry":                      config.IdentifierFromProvider,
	"tencentcloud_nat_gateway":                      config.IdentifierFromProvider,
	"tencentcloud_nat_gateway_snat":                 config.IdentifierFromProvider,
	"tencentcloud_dnat":                             config.IdentifierFromProvider,
	"tencentcloud_vpc_bandwidth_package":            config.IdentifierFromProvider,
	"tencentcloud_vpc_bandwidth_package_attachment": config.IdentifierFromProvider,
	"tencentcloud_vpn_gateway":                      config.IdentifierFromProvider,
	"tencentcloud_vpn_ssl_client":                   config.IdentifierFromProvider,
	"tencentcloud_vpn_customer_gateway":             config.IdentifierFromProvider,
	"tencentcloud_vpn_connection":                   config.IdentifierFromProvider,
	"tencentcloud_vpn_gateway_route":                config.IdentifierFromProvider,
	"tencentcloud_vpn_ssl_server":                   config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
