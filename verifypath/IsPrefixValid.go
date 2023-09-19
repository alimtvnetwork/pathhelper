package verifypath

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/pathhelper/ispath"
	"gitlab.com/auk-go/pathhelper/normalize"
)

func IsPrefixValid(
	isVerifyExistence bool,
	homeDirPrefix, currentFullPath string,
) (fixedPath string, isValid bool) {
	prefix := normalize.Path(homeDirPrefix)
	currentPathFix := normalize.Path(currentFullPath)

	if isVerifyExistence && ispath.NotExists(currentPathFix) {
		// error
		return constants.EmptyString, false
	}

	if normalize.HasPrefix(prefix, currentPathFix) {
		return currentPathFix, true
	}

	// homeDirPrefix doesn't match
	return constants.EmptyString, false
}
