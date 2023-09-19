package apachelinuxpath

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/osconsts"

	"gitlab.com/auk-go/pathhelper/knowndir"
	"gitlab.com/auk-go/pathhelper/knowndirget"
)

// GetSitesAvailable returns /etc/apache/sites-available as a string
func GetSitesAvailable() string {
	if osconsts.IsWindows {
		return constants.EmptyString
	}

	return knowndir.SitesAvailable.CombineWith(knowndirget.ApacheLinuxPath())
}
