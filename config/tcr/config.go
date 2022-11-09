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

package tcr

import (
	tjconfig "github.com/crossplane/terrajet/pkg/config"
)

// Configure configures the tcr group
func Configure(p *tjconfig.Provider) {
	p.AddResourceConfigurator("tencentcloud_tcr_instance", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = "tcr"
		r.Kind = "Instance"
	})

	p.AddResourceConfigurator("tencentcloud_tcr_namespace", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = "tcr"
		r.Kind = "Namespace"
	})

	p.AddResourceConfigurator("tencentcloud_tcr_repository", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = "tcr"
		r.Kind = "Repository"
		r.References["instance_id"] = tjconfig.Reference{
			Type: "Instance",
		}
	})

	p.AddResourceConfigurator("tencentcloud_tcr_token", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = "tcr"
		r.Kind = "Token"
		r.References["instance_id"] = tjconfig.Reference{
			Type: "Instance",
		}
	})

	p.AddResourceConfigurator("tencentcloud_tcr_vpc_attachment", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = "tcr"
		r.Kind = "VpcAttachment"
		r.References["instance_id"] = tjconfig.Reference{
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
