package normalizeinternal

import (
	"path"
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

func FixIf(isFix bool, currentPath1, currentPath2 string) string {
	if isFix && currentPath2 == "" {
		return path.Clean(currentPath1)
	}

	joinedPath := path.Join(currentPath1, currentPath2)

	if isFix {
		joinedPath = strings.ReplaceAll(joinedPath, constants.ForwardSlash, constants.BackSlash)

		return path.Clean(joinedPath)
	}

	return joinedPath
}
