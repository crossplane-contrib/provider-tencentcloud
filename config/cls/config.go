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

package cls

import "github.com/crossplane/upjet/pkg/config"

const shortGroupCls = "cls"

// Configure configures the cls group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("tencentcloud_cls_logset", func(r *config.Resource) {
		r.ShortGroup = shortGroupCls
		r.Kind = "Logset"
	})

	p.AddResourceConfigurator("tencentcloud_cls_topic", func(r *config.Resource) {
		r.ShortGroup = shortGroupCls
		r.Kind = "Topic"
		r.References["logset_id"] = config.Reference{
			Type: "Logset",
		}
	})

	p.AddResourceConfigurator("tencentcloud_cls_index", func(r *config.Resource) {
		r.ShortGroup = shortGroupCls
		r.Kind = "Index"
		r.References["topic_id"] = config.Reference{
			Type: "Topic",
		}
	})

	p.AddResourceConfigurator("tencentcloud_cls_cos_shipper", func(r *config.Resource) {
		r.ShortGroup = shortGroupCls
		r.Kind = "CosShipper"
		r.References["topic_id"] = config.Reference{
			Type: "Topic",
		}
	})

	p.AddResourceConfigurator("tencentcloud_cls_config", func(r *config.Resource) {
		r.ShortGroup = shortGroupCls
		r.Kind = "Config"
	})

	p.AddResourceConfigurator("tencentcloud_cls_config_extra", func(r *config.Resource) {
		r.ShortGroup = shortGroupCls
		r.Kind = "ConfigExtra"
		r.References["topic_id"] = config.Reference{
			Type: "Topic",
		}
		r.References["logset_id"] = config.Reference{
			Type: "Logset",
		}
	})

	p.AddResourceConfigurator("tencentcloud_cls_machine_group", func(r *config.Resource) {
		r.ShortGroup = shortGroupCls
		r.Kind = "MachineGroup"
	})

	p.AddResourceConfigurator("tencentcloud_cls_config_attachment", func(r *config.Resource) {
		r.ShortGroup = shortGroupCls
		r.Kind = "ConfigAttachment"
		r.References["config_id"] = config.Reference{
			Type: "Config",
		}
		r.References["group_id"] = config.Reference{
			Type: "MachineGroup",
		}
	})
}
