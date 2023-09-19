package namegroup

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errcmd"
	"gitlab.com/auk-go/errorwrapper/errtype"
)

func applyCmdOnPathsReturnOnErr(
	cmdPrefix string,
	paths []string,
) *errorwrapper.Wrapper {
	for _, currentPath := range paths {
		if currentPath == "" {
			return errorwrapper.NewPtr(errtype.EmptyFilePath)
		}

		// chgrp groupName path or chown -R $user:$group /dir
		pathCmd := cmdPrefix +
			constants.Space +
			currentPath

		errWrapper := errcmd.
			New.BashScript.ArgsErr(pathCmd)

		if errWrapper.HasError() {
			return errWrapper
		}
	}

	return nil
}
