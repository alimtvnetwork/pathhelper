package normalizeinternal

import "gitlab.com/auk-go/core/constants"

func JoinFixIf(isFix bool, currentPath1, currentPath2 string) string {
	if currentPath2 == constants.EmptyString {
		return FixIfSingle(isFix, currentPath1)
	}

	return JoinPathsFixIf(isFix, currentPath1, currentPath2)
}
