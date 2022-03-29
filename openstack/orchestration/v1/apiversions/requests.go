package apiversions

import (
	"github.com/nexclipper/gophercloud"
	"github.com/nexclipper/gophercloud/pagination"
)

// ListVersions lists all the Neutron API versions available to end-users
func ListVersions(c *gophercloud.ServiceClient) pagination.Pager {
	return pagination.NewPager(c, listURL(c), func(r pagination.PageResult) pagination.Page {
		return APIVersionPage{pagination.SinglePageBase(r)}
	})
}
