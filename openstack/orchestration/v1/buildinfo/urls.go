package buildinfo

import "github.com/nexclipper/gophercloud"

func getURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("build_info")
}
