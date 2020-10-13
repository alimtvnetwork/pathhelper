package pathhelper

import (
	"path"
	"strings"

	"gitlab.com/evatix-go/pathhelper/constants"
	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
)

var normalizeMap = map[string]string{
	constants.UriSchemePrefixStandard: "",
}

func NormalizePath(Path string) string {
	if pathhelpercore.IsEmptyPath(Path) {
		return Path
	}

	// if Path does not contain file prefixes
	if !HasPathIssues(Path) {
		return path.Clean(Path)
	}

	// when Path contains prefix
	pathWithoutPrefix := strings.Replace(Path, constants.UriSchemePrefixStandard, "", 1)
	pathWithoutPrefix = strings.Replace(pathWithoutPrefix, constants.UriSchemePrefixTwoSlashes, "", 1)

	// when givenPath contains UriSchemePrefixTwoSlashes
	pathWithoutPrefix = strings.Replace(pathWithoutPrefix, constants.UriSchemePrefixTwoSlashes, "", 1)

	// removing doubles
	pathWithoutDouble := RemovingDoubleSeparator(pathWithoutPrefix)

	// replacing with correct separator
	var pathWithCorrectSeparator string

	if IsWindows() {
		pathWithCorrectSeparator = strings.ReplaceAll(pathWithoutDouble, constants.ForwardSlash, constants.BackSlash)
	} else {
		pathWithCorrectSeparator = strings.ReplaceAll(pathWithoutDouble, constants.BackSlash, constants.ForwardSlash)
	}

	return RemovingDoubleSeparator(strings.TrimSpace(pathWithCorrectSeparator))
}
