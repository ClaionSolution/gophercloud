//go:build acceptance || blockstorage
// +build acceptance blockstorage

package extensions

import (
	"testing"

	"github.com/nexclipper/gophercloud/acceptance/clients"
	"github.com/nexclipper/gophercloud/acceptance/tools"
	"github.com/nexclipper/gophercloud/openstack/blockstorage/extensions/services"
	th "github.com/nexclipper/gophercloud/testhelper"
)

func TestServicesList(t *testing.T) {
	clients.RequireAdmin(t)

	blockClient, err := clients.NewBlockStorageV3Client()
	th.AssertNoErr(t, err)

	allPages, err := services.List(blockClient, services.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)

	allServices, err := services.ExtractServices(allPages)
	th.AssertNoErr(t, err)

	for _, service := range allServices {
		tools.PrintResource(t, service)
	}
}
