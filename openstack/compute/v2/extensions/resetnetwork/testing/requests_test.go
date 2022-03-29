package testing

import (
	"testing"

	"github.com/nexclipper/gophercloud/openstack/compute/v2/extensions/resetnetwork"
	th "github.com/nexclipper/gophercloud/testhelper"
	"github.com/nexclipper/gophercloud/testhelper/client"
)

const serverID = "b16ba811-199d-4ffd-8839-ba96c1185a67"

func TestResetNetwork(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	mockResetNetworkResponse(t, serverID)

	err := resetnetwork.ResetNetwork(client.ServiceClient(), serverID).ExtractErr()
	th.AssertNoErr(t, err)
}
