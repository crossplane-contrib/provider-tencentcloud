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

package monitor

import "github.com/crossplane/upjet/pkg/config"

const shortGroupMonitor = "monitor"

// Configure configures the monitor group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("tencentcloud_monitor_alarm_notice", func(r *config.Resource) {
		r.ShortGroup = shortGroupMonitor
		r.Kind = "AlarmNotice"
	})

	p.AddResourceConfigurator("tencentcloud_monitor_alarm_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroupMonitor
		r.Kind = "AlarmPolicy"
		r.References["notice_ids"] = config.Reference{
			Type: "AlarmNotice",
		}
	})

	p.AddResourceConfigurator("tencentcloud_monitor_binding_receiver", func(r *config.Resource) {
		r.ShortGroup = shortGroupMonitor
		r.Kind = "BindingReceiver"
	})

	p.AddResourceConfigurator("tencentcloud_monitor_grafana_instance", func(r *config.Resource) {
		r.ShortGroup = shortGroupMonitor
		r.Kind = "GrafanaInstance"
		r.References["vpc_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.VPC",
		}
		r.References["subnet_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.Subnet",
		}
	})

	p.AddResourceConfigurator("tencentcloud_monitor_grafana_integration", func(r *config.Resource) {
		r.ShortGroup = shortGroupMonitor
		r.Kind = "GrafanaIntegration"
		r.References["instance_id"] = config.Reference{
			Type: "GrafanaInstance",
		}
	})

	p.AddResourceConfigurator("tencentcloud_monitor_grafana_notification_channel", func(r *config.Resource) {
		r.ShortGroup = shortGroupMonitor
		r.Kind = "GrafanaNotificationChannel"
		r.References["instance_id"] = config.Reference{
			Type: "GrafanaInstance",
		}
	})

	p.AddResourceConfigurator("tencentcloud_monitor_grafana_plugin", func(r *config.Resource) {
		r.ShortGroup = shortGroupMonitor
		r.Kind = "GrafanaPlugin"
		r.References["instance_id"] = config.Reference{
			Type: "GrafanaInstance",
		}
	})

	p.AddResourceConfigurator("tencentcloud_monitor_grafana_sso_account", func(r *config.Resource) {
		r.ShortGroup = shortGroupMonitor
		r.Kind = "GrafanaSSOAccount"
		r.References["instance_id"] = config.Reference{
			Type: "GrafanaInstance",
		}
	})

	p.AddResourceConfigurator("tencentcloud_monitor_policy_binding_object", func(r *config.Resource) {
		r.ShortGroup = shortGroupMonitor
		r.Kind = "PolicyBindingObject"
		r.References["policy_id"] = config.Reference{
			Type: "AlarmPolicy",
		}
	})

	p.AddResourceConfigurator("tencentcloud_monitor_policy_group", func(r *config.Resource) {
		r.ShortGroup = shortGroupMonitor
		r.Kind = "PolicyGroup"
	})

	p.AddResourceConfigurator("tencentcloud_monitor_tmp_instance", func(r *config.Resource) {
		r.ShortGroup = shortGroupMonitor
		r.Kind = "TmpInstance"
		r.References["vpc_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.VPC",
		}
		r.References["subnet_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.Subnet",
		}
	})

	p.AddResourceConfigurator("tencentcloud_monitor_tmp_alert_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroupMonitor
		r.Kind = "TmpAlertRule"
		r.References["instance_id"] = config.Reference{
			Type: "TmpInstance",
		}
	})

	p.AddResourceConfigurator("tencentcloud_monitor_tmp_cvm_agent", func(r *config.Resource) {
		r.ShortGroup = shortGroupMonitor
		r.Kind = "TmpCvmAgent"
		r.References["instance_id"] = config.Reference{
			Type: "TmpInstance",
		}
	})

	p.AddResourceConfigurator("tencentcloud_monitor_tmp_exporter_integration", func(r *config.Resource) {
		r.ShortGroup = shortGroupMonitor
		r.Kind = "TmpExporterIntegration"
		r.References["instance_id"] = config.Reference{
			Type: "TmpInstance",
		}
	})

	p.AddResourceConfigurator("tencentcloud_monitor_tmp_recording_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroupMonitor
		r.Kind = "TmpRecordingRule"
	})

	p.AddResourceConfigurator("tencentcloud_monitor_tmp_scrape_job", func(r *config.Resource) {
		r.ShortGroup = shortGroupMonitor
		r.Kind = "TmpScrapeJob"
		r.References["instance_id"] = config.Reference{
			Type: "TmpInstance",
		}
	})

	p.AddResourceConfigurator("tencentcloud_monitor_tmp_tke_alert_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroupMonitor
		r.Kind = "TmpTkeAlertPolicy"
		r.References["instance_id"] = config.Reference{
			Type: "TmpInstance",
		}
	})

	p.AddResourceConfigurator("tencentcloud_monitor_tmp_tke_cluster_agent", func(r *config.Resource) {
		r.ShortGroup = shortGroupMonitor
		r.Kind = "TmpTkeClusterAgent"
		r.References["instance_id"] = config.Reference{
			Type: "TmpInstance",
		}
	})

	p.AddResourceConfigurator("tencentcloud_monitor_tmp_tke_config", func(r *config.Resource) {
		r.ShortGroup = shortGroupMonitor
		r.Kind = "TmpTkeConfig"
		r.References["instance_id"] = config.Reference{
			Type: "TmpInstance",
		}
	})

	p.AddResourceConfigurator("tencentcloud_monitor_tmp_tke_global_notification", func(r *config.Resource) {
		r.ShortGroup = shortGroupMonitor
		r.Kind = "TmpTkeGlobalNotification"
		r.References["instance_id"] = config.Reference{
			Type: "TmpInstance",
		}
	})

	p.AddResourceConfigurator("tencentcloud_monitor_tmp_tke_record_rule_yaml", func(r *config.Resource) {
		r.ShortGroup = shortGroupMonitor
		r.Kind = "TmpTkeRecordRuleYaml"
		r.References["instance_id"] = config.Reference{
			Type: "TmpInstance",
		}
	})

	p.AddResourceConfigurator("tencentcloud_monitor_tmp_tke_template", func(r *config.Resource) {
		r.ShortGroup = shortGroupMonitor
		r.Kind = "TmpTkeTemplate"
	})

	p.AddResourceConfigurator("tencentcloud_monitor_tmp_tke_template_attachment", func(r *config.Resource) {
		r.ShortGroup = shortGroupMonitor
		r.Kind = "TmpTkeTemplateAttachment"
		r.References["template_id"] = config.Reference{
			Type: "TmpTkeTemplate",
		}
	})
}
