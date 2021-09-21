package fs

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func CopyChmod(
	srcPath,
	dstPath string,
) *errorwrapper.Wrapper {
	srcFileInfo, err := os.Stat(srcPath)

	if IsNotPathExistsUsing(srcFileInfo, err) {
		return errnew.SourceDestinationMessages(
			errtype.ChownUserOrGroupApplyIssue,
			srcPath,
			dstPath,
			err.Error())
	}

	chmodApplyErr := os.Chmod(dstPath, srcFileInfo.Mode())

	if chmodApplyErr != nil {
		return errnew.SourceDestinationMessages(
			errtype.ChmodApplyFailed,
			srcPath,
			dstPath,
			err.Error(),
		)
	}

	return errnew.EmptyPtr
}
