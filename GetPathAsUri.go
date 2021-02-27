package pathhelper

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/pathhelper/normalize"
)

func GetPathAsUri(path string, isNormalizePath bool) string {
	if isNormalizePath {
		path = normalize.Path(path)
	}

	return constants.UriSchemePrefixStandard + strings.ReplaceAll(
		path,
		constants.BackSlash,
		constants.ForwardSlash)
}
