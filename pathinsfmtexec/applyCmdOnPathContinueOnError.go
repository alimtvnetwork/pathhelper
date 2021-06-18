package pathinsfmtexec

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errcmd"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func applyCmdOnPathContinueOnError(
	cmdPrefix string,
	paths []string,
) *errorwrapper.Wrapper {
	pathIssues := corestr.NewCollection(constants.Zero)

	for _, currentPath := range paths {
		if currentPath == "" {
			pathIssues.Add("Cannot process empty path.")

			continue
		}

		// chgrp groupName path or chown -R $user:$group /dir
		pathCmd := cmdPrefix +
			constants.Space +
			currentPath

		errWrapper := errcmd.BashArgsErrorWrapper(pathCmd)

		if errWrapper.HasError() {
			pathIssues.Add(pathCmd + " -- failed")
		}
	}

	if pathIssues.IsEmpty() {
		return errnew.EmptyPtr
	}

	return errnew.MessagesPtr(
		errtype.PathRelatedIssue,
		"Failed to execute cmd prefix :"+cmdPrefix,
		pathIssues.Join(constants.CommaSpace))
}
