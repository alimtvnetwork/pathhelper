package pathhelper

import (
	"strings"

	"gitlab.com/evatix-go/pathhelper/constants"
)

func GetPathAsUri(path string, isNormalizePath bool) string {
	if isNormalizePath {
		path = NormalizePath(path)
	}

	return constants.UriSchemePrefixStandard + strings.ReplaceAll(
		path,
		constants.BackSlash,
		constants.ForwardSlash)
}
