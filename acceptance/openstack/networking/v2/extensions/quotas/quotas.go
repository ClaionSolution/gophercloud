package quotas

import (
	"github.com/nexclipper/gophercloud"
	"github.com/nexclipper/gophercloud/openstack/networking/v2/extensions/quotas"
)

var updateOpts = quotas.UpdateOpts{
	FloatingIP:        gophercloud.IntToPointer(10),
	Network:           gophercloud.IntToPointer(-1),
	Port:              gophercloud.IntToPointer(11),
	RBACPolicy:        gophercloud.IntToPointer(0),
	Router:            gophercloud.IntToPointer(-1),
	SecurityGroup:     gophercloud.IntToPointer(12),
	SecurityGroupRule: gophercloud.IntToPointer(13),
	Subnet:            gophercloud.IntToPointer(14),
	SubnetPool:        gophercloud.IntToPointer(15),
}
