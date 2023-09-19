package namegroup

import "gitlab.com/auk-go/errorwrapper"

func applyCmdOnPaths(
	cmdPrefix string,
	paths []string,
	isContinueOnError bool,
) *errorwrapper.Wrapper {
	if !isContinueOnError {
		return applyCmdOnPathsReturnOnErr(cmdPrefix, paths)
	}

	return applyCmdOnPathContinueOnError(cmdPrefix, paths)
}
