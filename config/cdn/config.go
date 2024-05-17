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

package cdn

import "github.com/crossplane/upjet/pkg/config"

const shortGroupCdn = "cdn"

// Configure configures the cdn group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("tencentcloud_cdn_domain", func(r *config.Resource) {
		r.ShortGroup = shortGroupCdn
		r.Kind = "Domain"
	})

	p.AddResourceConfigurator("tencentcloud_cdn_url_purge", func(r *config.Resource) {
		r.ShortGroup = shortGroupCdn
		r.Kind = "UrlPurge"
	})

	p.AddResourceConfigurator("tencentcloud_cdn_url_push", func(r *config.Resource) {
		r.ShortGroup = shortGroupCdn
		r.Kind = "UrlPush"
	})
}
