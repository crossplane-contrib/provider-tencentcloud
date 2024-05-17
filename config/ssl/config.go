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

package ssl

import "github.com/crossplane/upjet/pkg/config"

const shortGroupSsl = "ssl"

// Configure configures the ssl group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("tencentcloud_ssl_certificate", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroupSsl
		r.Kind = "Certificate"
	})
	p.AddResourceConfigurator("tencentcloud_ssl_free_certificate", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroupSsl
		r.Kind = "FreeCertificate"
	})
	p.AddResourceConfigurator("tencentcloud_ssl_pay_certificate", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroupSsl
		r.Kind = "PayCertificate"
	})
}
