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

package mariadb

import "github.com/crossplane/upjet/pkg/config"

const shortGroupMariadb = "mariadb"

// Configure configures the mariadb group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("tencentcloud_mariadb_hour_db_instance", func(r *config.Resource) {
		r.ShortGroup = shortGroupMariadb
		r.Kind = "HourDBInstance"
		r.References["vpc_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.VPC",
		}
		r.References["subnet_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.Subnet",
		}
	})

	p.AddResourceConfigurator("tencentcloud_mariadb_dedicatedcluster_db_instance", func(r *config.Resource) {
		r.ShortGroup = shortGroupMariadb
		r.Kind = "DedicatedclusterDBInstance"
		r.References["vpc_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.VPC",
		}
		r.References["subnet_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.Subnet",
		}
	})

	p.AddResourceConfigurator("tencentcloud_mariadb_security_groups", func(r *config.Resource) {
		r.ShortGroup = shortGroupMariadb
		r.Kind = "SecurityGroups"
	})

	p.AddResourceConfigurator("tencentcloud_mariadb_parameters", func(r *config.Resource) {
		r.ShortGroup = shortGroupMariadb
		r.Kind = "Parameters"
	})

	p.AddResourceConfigurator("tencentcloud_mariadb_log_file_retention_period", func(r *config.Resource) {
		r.ShortGroup = shortGroupMariadb
		r.Kind = "LogFileRetentionPeriod"
	})

	p.AddResourceConfigurator("tencentcloud_mariadb_account", func(r *config.Resource) {
		r.ShortGroup = shortGroupMariadb
		r.Kind = "MariadbAccount"
	})
}
