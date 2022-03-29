package testing

import (
	"github.com/nexclipper/gophercloud"
	"github.com/nexclipper/gophercloud/testhelper"
)

func createClient() *gophercloud.ServiceClient {
	return &gophercloud.ServiceClient{
		ProviderClient: &gophercloud.ProviderClient{TokenID: "abc123"},
		Endpoint:       testhelper.Endpoint(),
	}
}
