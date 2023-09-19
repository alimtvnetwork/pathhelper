package nginxlinuxpath

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/osconsts"

	"gitlab.com/auk-go/pathhelper/knowndir"
	"gitlab.com/auk-go/pathhelper/knowndirget"
)

// GetMimeTypes
//
// returns /etc/nginx/mime.types as a string
func GetMimeTypes() string {
	if osconsts.IsWindows {
		return constants.EmptyString
	}

	if defaultMimeTypesPath.IsInitialized() {
		return defaultMimeTypesPath.String()
	}

	mimePath := knowndir.MimeTypes.CombineWith(
		knowndirget.NginxLinuxPath())

	return defaultMimeTypesPath.GetPlusSetOnUninitialized(
		mimePath)
}
