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

package tke

import "github.com/crossplane/upjet/pkg/config"

const shortGroupTke = "tke"

// Configure configures the tke group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("tencentcloud_kubernetes_cluster", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroupTke
		r.Kind = "Cluster"
		r.References["vpc_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.VPC",
		}
	})

	p.AddResourceConfigurator("tencentcloud_kubernetes_scale_worker", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroupTke
		r.Kind = "ScaleWorker"
		r.References["cluster_id"] = config.Reference{
			Type: "Cluster",
		}
	})

	p.AddResourceConfigurator("tencentcloud_kubernetes_cluster_attachment", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroupTke
		r.Kind = "ClusterAttachment"
		r.References["cluster_id"] = config.Reference{
			Type: "Cluster",
		}
		r.References["instance_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/cvm/v1alpha1.Instance",
		}
	})

	p.AddResourceConfigurator("tencentcloud_kubernetes_node_pool", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroupTke
		r.Kind = "NodePool"
		r.References["vpc_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.VPC",
		}
		r.References["subnet_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.Subnet",
		}
	})

	p.AddResourceConfigurator("tencentcloud_kubernetes_addon_attachment", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroupTke
		r.Kind = "AddonAttachment"
		r.References["cluster_id"] = config.Reference{
			Type: "Cluster",
		}
	})

	p.AddResourceConfigurator("tencentcloud_kubernetes_auth_attachment", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroupTke
		r.Kind = "AuthAttachment"
		r.References["cluster_id"] = config.Reference{
			Type: "Cluster",
		}
	})

	p.AddResourceConfigurator("tencentcloud_kubernetes_cluster_endpoint", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroupTke
		r.Kind = "ClusterEndpoint"
		r.References["cluster_id"] = config.Reference{
			Type: "Cluster",
		}
		r.References["cluster_intranet_subnet_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-tencentcloud/apis/vpc/v1alpha1.Subnet",
		}
	})
}
