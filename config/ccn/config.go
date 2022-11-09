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

package ccn

import (
	tjconfig "github.com/crossplane/terrajet/pkg/config"
)

// Configure configures the ccn group
func Configure(p *tjconfig.Provider) {
	p.AddResourceConfigurator("tencentcloud_ccn", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = "ccn"
		r.Kind = "CCN"
	})

	p.AddResourceConfigurator("tencentcloud_ccn_attachment", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = "ccn"
		r.Kind = "Attachment"
	})

	p.AddResourceConfigurator("tencentcloud_ccn_bandwidth_limit", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = "ccn"
		r.Kind = "BandwidthLimit"
		r.References["ccn_id"] = tjconfig.Reference{
			Type: "CCN",
		}
	})
}
