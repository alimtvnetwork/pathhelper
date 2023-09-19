package pathhelper

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/osconsts"

	"gitlab.com/auk-go/pathhelper/normalize"
)

func GetPathAsUri(path string, isNormalizePath bool) string {
	if isNormalizePath {
		path = normalize.PathUsingSeparatorIf(
			false,
			false,
			true,
			osconsts.PathSeparator,
			path)
	}

	return constants.UriSchemePrefixStandard + strings.ReplaceAll(
		path,
		constants.BackSlash,
		constants.ForwardSlash)
}
