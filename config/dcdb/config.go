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

package dcdb

import "github.com/crossplane/upjet/pkg/config"

const shortGroupDcdb = "dcdb"

// Configure configures the dcdb group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("tencentcloud_dcdb_hourdb_instance", func(r *config.Resource) {
		r.ShortGroup = shortGroupDcdb
		r.Kind = "HourdbInstance"
		r.References["vpc_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.VPC",
		}
		r.References["subnet_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.Subnet",
		}
	})

	p.AddResourceConfigurator("tencentcloud_dcdb_account", func(r *config.Resource) {
		r.ShortGroup = shortGroupDcdb
		r.Kind = "DcdbAccount"
		r.References["instance_id"] = config.Reference{
			Type: "HourdbInstance",
		}
	})

	p.AddResourceConfigurator("tencentcloud_dcdb_security_group_attachment", func(r *config.Resource) {
		r.ShortGroup = shortGroupDcdb
		r.Kind = "SecurityGroupAttachment"
		r.References["instance_id"] = config.Reference{
			Type: "HourdbInstance",
		}
		r.References["security_group_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.SecurityGroup",
		}
	})
}
