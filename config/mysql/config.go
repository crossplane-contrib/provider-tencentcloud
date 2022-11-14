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

package mysql

import (
	tjconfig "github.com/crossplane/terrajet/pkg/config"
)

const shortGroupMysql = "mysql"

// Configure configures the mysql group
func Configure(p *tjconfig.Provider) {
	p.AddResourceConfigurator("tencentcloud_mysql_instance", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupMysql
		r.Kind = "Instance"
		r.References["vpc_id"] = tjconfig.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.VPC",
		}
		r.References["subnet_id"] = tjconfig.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.Subnet",
		}
	})

	p.AddResourceConfigurator("tencentcloud_mysql_account", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupMysql
		r.Kind = "Account"
		r.References["mysql_id"] = tjconfig.Reference{
			Type: "Instance",
		}
	})

	p.AddResourceConfigurator("tencentcloud_mysql_account_privilege", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupMysql
		r.Kind = "AccountPrivilege"
		r.References["mysql_id"] = tjconfig.Reference{
			Type: "Instance",
		}
	})

	p.AddResourceConfigurator("tencentcloud_mysql_backup_policy", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupMysql
		r.Kind = "BackupPolicy"
		r.References["mysql_id"] = tjconfig.Reference{
			Type: "Instance",
		}
	})

	p.AddResourceConfigurator("tencentcloud_mysql_privilege", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupMysql
		r.Kind = "Privilege"
		r.References["mysql_id"] = tjconfig.Reference{
			Type: "Instance",
		}
	})

	p.AddResourceConfigurator("tencentcloud_mysql_readonly_instance", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupMysql
		r.Kind = "ReadonlyInstance"
		r.References["master_instance_id"] = tjconfig.Reference{
			Type: "Instance",
		}
		r.References["vpc_id"] = tjconfig.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.VPC",
		}
		r.References["subnet_id"] = tjconfig.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.Subnet",
		}
	})
}
