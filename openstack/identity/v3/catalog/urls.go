package catalog

import "github.com/nexclipper/gophercloud"

func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("auth", "catalog")
}
