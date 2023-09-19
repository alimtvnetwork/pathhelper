package namegroup

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errcmd"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
)

func applyCmdOnPathContinueOnError(
	cmdPrefix string,
	paths []string,
) *errorwrapper.Wrapper {
	pathIssues := corestr.New.Collection.Cap(constants.Zero)

	for _, currentPath := range paths {
		if currentPath == "" {
			pathIssues.Add("Cannot process empty path.")

			continue
		}

		// chgrp groupName path or chown -R $user:$group /dir
		pathCmd := cmdPrefix +
			constants.Space +
			currentPath

		errWrapper := errcmd.
			New.BashScript.ArgsErr(pathCmd)

		if errWrapper.HasError() {
			pathIssues.Add(pathCmd + " -- failed")
		}
	}

	if pathIssues.IsEmpty() {
		return nil
	}

	return errnew.Messages.Many(
		errtype.PathRelatedIssue,
		"Failed to execute cmd prefix :"+cmdPrefix,
		pathIssues.Join(constants.CommaSpace))
}
