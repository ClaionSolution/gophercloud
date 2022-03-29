package serviceassets

import "github.com/nexclipper/gophercloud"

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("services", id, "assets")
}
