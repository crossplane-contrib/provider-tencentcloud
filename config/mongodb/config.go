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

package mongodb

import "github.com/crossplane/upjet/pkg/config"

const shortGroupMongodb = "mongodb"

// Configure configures the mongodb group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("tencentcloud_mongodb_instance", func(r *config.Resource) {
		r.ShortGroup = shortGroupMongodb
		r.Kind = "Instance"
		r.References["vpc_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.VPC",
		}
		r.References["subnet_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.Subnet",
		}
	})

	p.AddResourceConfigurator("tencentcloud_mongodb_sharding_instance", func(r *config.Resource) {
		r.ShortGroup = shortGroupMongodb
		r.Kind = "ShardingInstance"
		r.References["vpc_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.VPC",
		}
		r.References["subnet_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.Subnet",
		}
	})

	p.AddResourceConfigurator("tencentcloud_mongodb_standby_instance", func(r *config.Resource) {
		r.ShortGroup = shortGroupMongodb
		r.Kind = "StandbyInstance"
		r.References["father_instance_id"] = config.Reference{
			Type: "Instance",
		}
	})

}
