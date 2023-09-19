package fs

import (
	"os"

	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
)

func CopyChmodChown(
	srcPath,
	dstPath string,
) *errorwrapper.Wrapper {
	srcFileInfo, err := os.Stat(srcPath)

	if IsNotPathExistsUsing(srcFileInfo, err) {
		return errnew.SrcDst.Error(
			errtype.ChownUserOrGroupApplyIssue,
			err,
			srcPath,
			dstPath,
		)
	}

	return copyChmodChownInternal(srcPath, dstPath, srcFileInfo)
}
