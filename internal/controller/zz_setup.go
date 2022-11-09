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

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/terrajet/pkg/controller"

	attachment "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/as/attachment"
	lifecyclehook "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/as/lifecyclehook"
	notification "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/as/notification"
	scalingconfig "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/as/scalingconfig"
	scalinggroup "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/as/scalinggroup"
	scalingpolicy "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/as/scalingpolicy"
	schedule "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/as/schedule"
	audit "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/audit/audit"
	group "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cam/group"
	groupmembership "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cam/groupmembership"
	grouppolicyattachment "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cam/grouppolicyattachment"
	oidcsso "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cam/oidcsso"
	policy "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cam/policy"
	role "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cam/role"
	rolepolicyattachment "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cam/rolepolicyattachment"
	rolesso "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cam/rolesso"
	samlprovider "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cam/samlprovider"
	user "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cam/user"
	userpolicyattachment "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cam/userpolicyattachment"
	snapshot "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cbs/snapshot"
	snapshotpolicy "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cbs/snapshotpolicy"
	snapshotpolicyattachment "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cbs/snapshotpolicyattachment"
	storage "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cbs/storage"
	storageattachment "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cbs/storageattachment"
	storageset "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cbs/storageset"
	attachmentccn "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/ccn/attachment"
	bandwidthlimit "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/ccn/bandwidthlimit"
	ccn "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/ccn/ccn"
	instance "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cdh/instance"
	domain "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cdn/domain"
	urlpurge "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cdn/urlpurge"
	urlpush "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cdn/urlpush"
	accessgroup "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cfs/accessgroup"
	accessrule "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cfs/accessrule"
	filesystem "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cfs/filesystem"
	albserverattachment "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/clb/albserverattachment"
	attachmentclb "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/clb/attachment"
	customizedconfig "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/clb/customizedconfig"
	instanceclb "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/clb/instance"
	lb "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/clb/lb"
	listener "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/clb/listener"
	listenerrule "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/clb/listenerrule"
	logset "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/clb/logset"
	logtopic "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/clb/logtopic"
	redirection "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/clb/redirection"
	snatip "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/clb/snatip"
	targetgroup "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/clb/targetgroup"
	targetgroupattachment "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/clb/targetgroupattachment"
	targetgroupinstanceattachment "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/clb/targetgroupinstanceattachment"
	containercluster "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/containercluster/containercluster"
	containerclusterinstance "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/containercluster/containerclusterinstance"
	bucket "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cos/bucket"
	bucketdomaincertificateattachment "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cos/bucketdomaincertificateattachment"
	bucketobject "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cos/bucketobject"
	bucketpolicy "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cos/bucketpolicy"
	image "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cvm/image"
	instancecvm "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cvm/instance"
	instanceset "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cvm/instanceset"
	keypair "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cvm/keypair"
	placementgroup "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cvm/placementgroup"
	reservedinstance "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cvm/reservedinstance"
	cluster "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cynosdb/cluster"
	readonlyinstance "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/cynosdb/readonlyinstance"
	eni "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/eni/eni"
	eniattachment "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/eni/eniattachment"
	record "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/privatedns/record"
	zone "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/privatedns/zone"
	providerconfig "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/providerconfig"
	certificate "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/ssl/certificate"
	freecertificate "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/ssl/freecertificate"
	paycertificate "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/ssl/paycertificate"
	instancetcr "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/tcr/instance"
	namespace "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/tcr/namespace"
	repository "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/tcr/repository"
	token "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/tcr/token"
	vpcattachment "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/tcr/vpcattachment"
	adaptivedynamicstreamingtemplate "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vod/adaptivedynamicstreamingtemplate"
	imagespritetemplate "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vod/imagespritetemplate"
	proceduretemplate "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vod/proceduretemplate"
	snapshotbytimeoffsettemplate "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vod/snapshotbytimeoffsettemplate"
	subapplication "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vod/subapplication"
	superplayerconfig "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vod/superplayerconfig"
	addresstemplate "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/addresstemplate"
	addresstemplategroup "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/addresstemplategroup"
	dnat "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/dnat"
	eip "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/eip"
	eipassociation "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/eipassociation"
	havip "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/havip"
	natgateway "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/natgateway"
	natgatewaysnat "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/natgatewaysnat"
	protocoltemplate "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/protocoltemplate"
	protocoltemplategroup "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/protocoltemplategroup"
	routeentry "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/routeentry"
	routetable "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/routetable"
	routetableentry "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/routetableentry"
	securitygroup "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/securitygroup"
	securitygroupliterule "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/securitygroupliterule"
	securitygrouprule "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/securitygrouprule"
	subnet "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/subnet"
	vpc "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/vpc"
	vpcacl "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/vpcacl"
	vpcbandwidthpackage "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/vpcbandwidthpackage"
	vpcbandwidthpackageattachment "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/vpcbandwidthpackageattachment"
	vpnconnection "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/vpnconnection"
	vpncustomergateway "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/vpncustomergateway"
	vpngateway "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/vpngateway"
	vpngatewayroute "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/vpngatewayroute"
	vpnsslclient "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/vpnsslclient"
	vpnsslserver "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/vpnsslserver"
	vpvaclattachment "github.com/crossplane-contrib/provider-tencentcloud/internal/controller/vpc/vpvaclattachment"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		attachment.Setup,
		lifecyclehook.Setup,
		notification.Setup,
		scalingconfig.Setup,
		scalinggroup.Setup,
		scalingpolicy.Setup,
		schedule.Setup,
		audit.Setup,
		group.Setup,
		groupmembership.Setup,
		grouppolicyattachment.Setup,
		oidcsso.Setup,
		policy.Setup,
		role.Setup,
		rolepolicyattachment.Setup,
		rolesso.Setup,
		samlprovider.Setup,
		user.Setup,
		userpolicyattachment.Setup,
		snapshot.Setup,
		snapshotpolicy.Setup,
		snapshotpolicyattachment.Setup,
		storage.Setup,
		storageattachment.Setup,
		storageset.Setup,
		attachmentccn.Setup,
		bandwidthlimit.Setup,
		ccn.Setup,
		instance.Setup,
		domain.Setup,
		urlpurge.Setup,
		urlpush.Setup,
		accessgroup.Setup,
		accessrule.Setup,
		filesystem.Setup,
		albserverattachment.Setup,
		attachmentclb.Setup,
		customizedconfig.Setup,
		instanceclb.Setup,
		lb.Setup,
		listener.Setup,
		listenerrule.Setup,
		logset.Setup,
		logtopic.Setup,
		redirection.Setup,
		snatip.Setup,
		targetgroup.Setup,
		targetgroupattachment.Setup,
		targetgroupinstanceattachment.Setup,
		containercluster.Setup,
		containerclusterinstance.Setup,
		bucket.Setup,
		bucketdomaincertificateattachment.Setup,
		bucketobject.Setup,
		bucketpolicy.Setup,
		image.Setup,
		instancecvm.Setup,
		instanceset.Setup,
		keypair.Setup,
		placementgroup.Setup,
		reservedinstance.Setup,
		cluster.Setup,
		readonlyinstance.Setup,
		eni.Setup,
		eniattachment.Setup,
		record.Setup,
		zone.Setup,
		providerconfig.Setup,
		certificate.Setup,
		freecertificate.Setup,
		paycertificate.Setup,
		instancetcr.Setup,
		namespace.Setup,
		repository.Setup,
		token.Setup,
		vpcattachment.Setup,
		adaptivedynamicstreamingtemplate.Setup,
		imagespritetemplate.Setup,
		proceduretemplate.Setup,
		snapshotbytimeoffsettemplate.Setup,
		subapplication.Setup,
		superplayerconfig.Setup,
		addresstemplate.Setup,
		addresstemplategroup.Setup,
		dnat.Setup,
		eip.Setup,
		eipassociation.Setup,
		havip.Setup,
		natgateway.Setup,
		natgatewaysnat.Setup,
		protocoltemplate.Setup,
		protocoltemplategroup.Setup,
		routeentry.Setup,
		routetable.Setup,
		routetableentry.Setup,
		securitygroup.Setup,
		securitygroupliterule.Setup,
		securitygrouprule.Setup,
		subnet.Setup,
		vpc.Setup,
		vpcacl.Setup,
		vpcbandwidthpackage.Setup,
		vpcbandwidthpackageattachment.Setup,
		vpnconnection.Setup,
		vpncustomergateway.Setup,
		vpngateway.Setup,
		vpngatewayroute.Setup,
		vpnsslclient.Setup,
		vpnsslserver.Setup,
		vpvaclattachment.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
