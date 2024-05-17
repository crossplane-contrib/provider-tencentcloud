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

package sqlserver

import "github.com/crossplane/upjet/pkg/config"

const shortGroupSQLServer = "sqlserver"

// Configure configures the sqlserver group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("tencentcloud_sqlserver_basic_instance", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroupSQLServer
		r.Kind = "BasicInstance"
		r.References["vpc_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.VPC",
		}
		r.References["subnet_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.Subnet",
		}
	})

	p.AddResourceConfigurator("tencentcloud_sqlserver_instance", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroupSQLServer
		r.Kind = "Instance"
		r.References["vpc_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.VPC",
		}
		r.References["subnet_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.Subnet",
		}
	})

	p.AddResourceConfigurator("tencentcloud_sqlserver_db", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroupSQLServer
		r.Kind = "DB"
		r.References["instance_id"] = config.Reference{
			Type: "Instance",
		}
	})

	p.AddResourceConfigurator("tencentcloud_sqlserver_readonly_instance", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroupSQLServer
		r.Kind = "ReadonlyInstance"
		r.References["vpc_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.VPC",
		}
		r.References["subnet_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.Subnet",
		}
		r.References["master_instance_id"] = config.Reference{
			Type: "Instance",
		}
	})

	p.AddResourceConfigurator("tencentcloud_sqlserver_account", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroupSQLServer
		r.Kind = "Account"
		r.References["instance_id"] = config.Reference{
			Type: "Instance",
		}
	})

	p.AddResourceConfigurator("tencentcloud_sqlserver_account_db_attachment", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroupSQLServer
		r.Kind = "AccountDBAttachment"
		r.References["instance_id"] = config.Reference{
			Type: "Instance",
		}
	})

	p.AddResourceConfigurator("tencentcloud_sqlserver_publish_subscribe", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroupSQLServer
		r.Kind = "PublishSubscribe"
	})
}
