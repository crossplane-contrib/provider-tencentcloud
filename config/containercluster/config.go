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

package containercluster

import (
	tjconfig "github.com/crossplane/terrajet/pkg/config"
)

const shortGroupContainerCluster = "containercluster"

// Configure configures the container_cluster group
func Configure(p *tjconfig.Provider) {
	p.AddResourceConfigurator("tencentcloud_container_cluster", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupContainerCluster
		r.Kind = "ContainerCluster"
	})

	p.AddResourceConfigurator("tencentcloud_container_cluster_instance", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = shortGroupContainerCluster
		r.Kind = "ContainerClusterInstance"
		r.References["cluster_id"] = tjconfig.Reference{
			Type: "ContainerCluster",
		}
	})
}
