package volumeactions

import "github.com/nexclipper/gophercloud"

func actionURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("volumes", id, "action")
}
