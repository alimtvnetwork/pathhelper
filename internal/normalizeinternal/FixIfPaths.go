package normalizeinternal

import (
	"path"
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

func FixIfPaths(isFix bool, givenPaths ...string) string {
	if len(givenPaths) == 0 {
		return constants.EmptyString
	}

	joinedPath := path.Join(givenPaths...)

	if isFix {
		joinedPath = strings.ReplaceAll(joinedPath, constants.ForwardSlash, constants.BackSlash)

		return path.Clean(joinedPath)
	}

	return joinedPath
}
