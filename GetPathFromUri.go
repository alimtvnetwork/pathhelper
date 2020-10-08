package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/constants"
)

var uriRemovePrefixes = []string{
	constants.UriSchemePrefixStandard,
	constants.UriSchemePrefixTwoSlashes,
}

func GetPathFromUri(path string, isNormalizePath bool) string {
	return RemoveFromPath(path, &uriRemovePrefixes, isNormalizePath)
}
