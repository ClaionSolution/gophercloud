package availabilityzones

import "github.com/nexclipper/gophercloud"

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("os-availability-zone")
}
