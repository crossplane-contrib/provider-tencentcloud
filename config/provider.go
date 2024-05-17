/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/crossplane-contrib/provider-tencentcloud/config/apigateway"
	"github.com/crossplane-contrib/provider-tencentcloud/config/as"
	"github.com/crossplane-contrib/provider-tencentcloud/config/audit"
	"github.com/crossplane-contrib/provider-tencentcloud/config/cam"
	"github.com/crossplane-contrib/provider-tencentcloud/config/cat"
	"github.com/crossplane-contrib/provider-tencentcloud/config/cbs"
	"github.com/crossplane-contrib/provider-tencentcloud/config/ccn"
	"github.com/crossplane-contrib/provider-tencentcloud/config/cdh"
	"github.com/crossplane-contrib/provider-tencentcloud/config/cdn"
	"github.com/crossplane-contrib/provider-tencentcloud/config/cfs"
	"github.com/crossplane-contrib/provider-tencentcloud/config/clb"
	"github.com/crossplane-contrib/provider-tencentcloud/config/cls"
	"github.com/crossplane-contrib/provider-tencentcloud/config/cos"
	"github.com/crossplane-contrib/provider-tencentcloud/config/cvm"
	"github.com/crossplane-contrib/provider-tencentcloud/config/cynosdb"
	"github.com/crossplane-contrib/provider-tencentcloud/config/dayu"
	"github.com/crossplane-contrib/provider-tencentcloud/config/dcdb"
	"github.com/crossplane-contrib/provider-tencentcloud/config/dnspod"
	"github.com/crossplane-contrib/provider-tencentcloud/config/elasticsearch"
	"github.com/crossplane-contrib/provider-tencentcloud/config/emr"
	"github.com/crossplane-contrib/provider-tencentcloud/config/eni"
	"github.com/crossplane-contrib/provider-tencentcloud/config/gaap"
	"github.com/crossplane-contrib/provider-tencentcloud/config/kafka"
	"github.com/crossplane-contrib/provider-tencentcloud/config/kms"
	"github.com/crossplane-contrib/provider-tencentcloud/config/lighthouse"
	"github.com/crossplane-contrib/provider-tencentcloud/config/mariadb"
	"github.com/crossplane-contrib/provider-tencentcloud/config/mongodb"
	"github.com/crossplane-contrib/provider-tencentcloud/config/monitor"
	"github.com/crossplane-contrib/provider-tencentcloud/config/mysql"
	"github.com/crossplane-contrib/provider-tencentcloud/config/postgresql"
	"github.com/crossplane-contrib/provider-tencentcloud/config/privatedns"
	"github.com/crossplane-contrib/provider-tencentcloud/config/redis"
	"github.com/crossplane-contrib/provider-tencentcloud/config/scf"
	"github.com/crossplane-contrib/provider-tencentcloud/config/ses"
	"github.com/crossplane-contrib/provider-tencentcloud/config/sms"
	"github.com/crossplane-contrib/provider-tencentcloud/config/sqlserver"
	"github.com/crossplane-contrib/provider-tencentcloud/config/ssl"
	"github.com/crossplane-contrib/provider-tencentcloud/config/ssm"
	"github.com/crossplane-contrib/provider-tencentcloud/config/tcaplus"
	"github.com/crossplane-contrib/provider-tencentcloud/config/tcm"
	"github.com/crossplane-contrib/provider-tencentcloud/config/tcr"
	"github.com/crossplane-contrib/provider-tencentcloud/config/tdmq"
	"github.com/crossplane-contrib/provider-tencentcloud/config/tem"
	"github.com/crossplane-contrib/provider-tencentcloud/config/teo"
	"github.com/crossplane-contrib/provider-tencentcloud/config/tke"
	"github.com/crossplane-contrib/provider-tencentcloud/config/vod"
	"github.com/crossplane-contrib/provider-tencentcloud/config/vpc"

	ujconfig "github.com/crossplane/upjet/pkg/config"
)

const (
	resourcePrefix = "tencentcloud"
	modulePath     = "github.com/crossplane-contrib/provider-tencentcloud"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("tencentcloud.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		apigateway.Configure,
		as.Configure,
		audit.Configure,
		cam.Configure,
		cat.Configure,
		cbs.Configure,
		ccn.Configure,
		cdh.Configure,
		cdn.Configure,
		cfs.Configure,
		clb.Configure,
		cls.Configure,
		cos.Configure,
		cvm.Configure,
		cynosdb.Configure,
		dayu.Configure,
		dcdb.Configure,
		dnspod.Configure,
		elasticsearch.Configure,
		emr.Configure,
		eni.Configure,
		gaap.Configure,
		kafka.Configure,
		kms.Configure,
		lighthouse.Configure,
		mariadb.Configure,
		mongodb.Configure,
		monitor.Configure,
		mysql.Configure,
		postgresql.Configure,
		privatedns.Configure,
		redis.Configure,
		scf.Configure,
		ses.Configure,
		sms.Configure,
		sqlserver.Configure,
		ssl.Configure,
		ssm.Configure,
		tcaplus.Configure,
		tcm.Configure,
		tcr.Configure,
		tdmq.Configure,
		tem.Configure,
		teo.Configure,
		tke.Configure,
		vod.Configure,
		vpc.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
