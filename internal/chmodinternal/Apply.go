package chmodinternal

import (
	"os"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/msgtype"
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func Apply(
	isSkipOnInvalid bool,
	mode os.FileMode,
	rootDir string,
	filePaths []string,
) *errorwrapper.Wrapper {
	fileMode := chmodhelper.NewUsingFileMode(mode)

	if osconsts.IsLinux {
		chmodErr := fileMode.
			LinuxApplyRecursive(
				isSkipOnInvalid,
				rootDir)

		return errnew.Path(
			errtype.ChmodApplyFailed,
			chmodErr,
			rootDir)
	}

	sliceErr := corestr.EmptySimpleSlice()
	// for other os apply using each file.
	for _, filePath := range filePaths {
		err := fileMode.ApplyChmod(isSkipOnInvalid, filePath)

		if err != nil {
			sliceErr.Add(err.Error())
		}
	}

	toErr := msgtype.SliceToError(sliceErr.Items)

	return errnew.Path(
		errtype.ChmodApplyFailed,
		toErr,
		rootDir)
}
