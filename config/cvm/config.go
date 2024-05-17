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

package cvm

import "github.com/crossplane/upjet/pkg/config"

const shortGroupCvm = "cvm"

// Configure configures the cvm group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("tencentcloud_instance", func(r *config.Resource) {
		r.ShortGroup = shortGroupCvm
		r.References["vpc_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.VPC",
		}
		r.References["subnet_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.Subnet",
		}
		if s, ok := r.TerraformResource.Schema["system_disk_id"]; ok {
			s.Optional = false
			s.Computed = true
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{
				"security_groups",
				"key_ids",
			},
		}
	})

	p.AddResourceConfigurator("tencentcloud_instance_set", func(r *config.Resource) {
		r.ShortGroup = shortGroupCvm
		r.Kind = "InstanceSet"
		r.References["vpc_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.VPC",
		}
		r.References["subnet_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.Subnet",
		}
	})

	p.AddResourceConfigurator("tencentcloud_key_pair", func(r *config.Resource) {
		r.ShortGroup = shortGroupCvm
		r.Kind = "KeyPair"
	})

	p.AddResourceConfigurator("tencentcloud_placement_group", func(r *config.Resource) {
		r.ShortGroup = shortGroupCvm
		r.Kind = "PlacementGroup"
	})

	p.AddResourceConfigurator("tencentcloud_reserved_instance", func(r *config.Resource) {
		r.ShortGroup = shortGroupCvm
		r.Kind = "ReservedInstance"
	})

	p.AddResourceConfigurator("tencentcloud_image", func(r *config.Resource) {
		r.ShortGroup = shortGroupCvm
	})
}
