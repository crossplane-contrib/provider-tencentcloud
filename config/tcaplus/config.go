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

package tcaplus

import (
	tjconfig "github.com/crossplane/terrajet/pkg/config"
)

const shortGroupTcaplus = "tcaplus"

// Configure configures the tcaplus group
func Configure(p *tjconfig.Provider) {
	p.AddResourceConfigurator("tencentcloud_tcaplus_cluster", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupTcaplus
		r.Kind = "Cluster"
		r.References["vpc_id"] = tjconfig.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.VPC",
		}
		r.References["subnet_id"] = tjconfig.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.Subnet",
		}
	})

	p.AddResourceConfigurator("tencentcloud_tcaplus_tablegroup", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupTcaplus
		r.Kind = "TableGroup"
		r.References["cluster_id"] = tjconfig.Reference{
			Type: "Cluster",
		}
	})

	p.AddResourceConfigurator("tencentcloud_tcaplus_table", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupTcaplus
		r.Kind = "Table"
		r.References["cluster_id"] = tjconfig.Reference{
			Type: "Cluster",
		}
		r.References["tablegroup_id"] = tjconfig.Reference{
			Type: "TableGroup",
		}
	})

	p.AddResourceConfigurator("tencentcloud_tcaplus_idl", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupTcaplus
		r.Kind = "Idl"
		r.References["cluster_id"] = tjconfig.Reference{
			Type: "Cluster",
		}
		r.References["tablegroup_id"] = tjconfig.Reference{
			Type: "TableGroup",
		}
	})
}
