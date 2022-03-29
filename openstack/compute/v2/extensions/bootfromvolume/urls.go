package bootfromvolume

import "github.com/nexclipper/gophercloud"

func createURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("servers")
}
