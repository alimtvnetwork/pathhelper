package apachelinuxpath

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/osconsts"
	"gitlab.com/auk-go/errorwrapper/errtype"

	"gitlab.com/auk-go/pathhelper/knowndir"
	"gitlab.com/auk-go/pathhelper/knowndirget"
)

// GetConfAvailable returns /etc/apache/conf-available as a string
func GetConfAvailable() string {
	if osconsts.IsWindows {
		errtype.NotSupportInWindows.PanicNoRefs(constants.EmptyString)
	}

	return knowndir.ConfAvailable.CombineWith(knowndirget.ApacheLinuxPath())
}
