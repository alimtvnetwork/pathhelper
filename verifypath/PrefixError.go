package verifypath

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
	"gitlab.com/auk-go/pathhelper/ispath"
	"gitlab.com/auk-go/pathhelper/normalize"
)

func PrefixError(
	isVerifyExistence bool,
	homeDirPrefix, currentFullPath string,
) (fixedPath string, verificationErrorWrapper *errorwrapper.Wrapper) {
	prefixFix := normalize.Path(homeDirPrefix)
	currentPathFix := normalize.Path(currentFullPath)

	if isVerifyExistence && ispath.NotExists(currentPathFix) {
		// error
		return constants.EmptyString, errnew.
			Path.
			Messages(
				errtype.FileInvalidOrMissing,
				currentPathFix,
				"given path is not valid in the file system.",
				"path must contain user home prefix and don't contain any relative path!",
				"user home prefix:",
				prefixFix,
			)
	}

	if normalize.HasPrefix(prefixFix, currentPathFix) {
		return currentPathFix, nil
	}

	// homeDirPrefix doesn't match
	return constants.EmptyString, errnew.Ref.Messages(
		errtype.PathSyntaxIssue,
		"Prefix",
		prefixFix,
		"current path homeDirPrefix missing (\""+homeDirPrefix+"\")",
		"path needs to have user home dir!",
		"current path:",
		currentPathFix,
		"homeDirPrefix",
		prefixFix,
	)
}
