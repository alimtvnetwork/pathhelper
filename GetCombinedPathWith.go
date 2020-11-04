package pathhelper

import (
	"strings"

	"gitlab.com/evatix-go/pathhelper/constants"
)

// todo should be deleted
// Returns path as string after combining all provided paths
// By default isIgnorePath: true, isNormalize: true, and Path separator depends on the OS
func GetCombinePathsWith(paths ...string) string {
	return GetCombinedPath(
		constants.PathSeparator,
		true,
		true,
		true,
		strings.Join(paths, constants.PathSeparator))
}
