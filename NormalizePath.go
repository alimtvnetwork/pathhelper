package pathhelper

import (
	"path"
	"strings"

	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
)

func NormalizePath(givenPath string) string {
	if pathhelpercore.IsEmptyPath(givenPath) {
		return givenPath
	}

	// if givenPath does not contain file prefixes
	if !HasPathIssues(givenPath) {
		return path.Clean(givenPath)
	}

	// when givenPath contains file
	pathWithoutPrefix := strings.Replace(givenPath, Prefix, "", 1)

	// removing doubles
	pathWithoutDouble := RemovingDouble(pathWithoutPrefix)

	// replacing with correct separator
	var pathWithCorrectSeparator string

	if IsWindows() {
		pathWithCorrectSeparator = strings.ReplaceAll(pathWithoutDouble, Slash, BackSlash)
	} else {
		pathWithCorrectSeparator = strings.ReplaceAll(pathWithoutDouble, BackSlash, Slash)
	}

	return RemovingDouble(strings.TrimSpace(pathWithCorrectSeparator))
}
