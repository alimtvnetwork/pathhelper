package normalizeinternal

import "gitlab.com/evatix-go/core/constants"

func FixIf(isFix bool, currentPath1, currentPath2 string) string {
	if currentPath2 == constants.EmptyString {
		return FixIfPaths(isFix, currentPath1)
	}

	return FixIfPaths(isFix, currentPath1, currentPath2)
}
