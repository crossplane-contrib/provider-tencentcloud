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

package cfs

import (
	tjconfig "github.com/crossplane/terrajet/pkg/config"
)

// Configure configures the cfs group
func Configure(p *tjconfig.Provider) {
	p.AddResourceConfigurator("tencentcloud_cfs_file_system", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = "cfs"
		r.Kind = "FileSystem"
		r.References["vpc_id"] = tjconfig.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.VPC",
		}
		r.References["subnet_id"] = tjconfig.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.Subnet",
		}
		r.References["access_group_id"] = tjconfig.Reference{
			Type: "AccessGroup",
		}
	})

	p.AddResourceConfigurator("tencentcloud_cfs_access_group", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = "cfs"
		r.Kind = "AccessGroup"
	})

	p.AddResourceConfigurator("tencentcloud_cfs_access_rule", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = "cfs"
		r.Kind = "AccessRule"
		r.References["access_group_id"] = tjconfig.Reference{
			Type: "AccessGroup",
		}
	})

}
