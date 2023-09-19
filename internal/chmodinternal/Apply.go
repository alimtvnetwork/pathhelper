package chmodinternal

import (
	"os"

	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/osconsts"
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
)

func Apply(
	isSkipOnInvalid bool,
	mode os.FileMode,
	rootDir string,
	filePaths []string,
) *errorwrapper.Wrapper {
	fileMode := chmodhelper.New.RwxWrapper.UsingFileMode(mode)

	if osconsts.IsLinux {
		chmodErr := fileMode.
			LinuxApplyRecursive(
				isSkipOnInvalid,
				rootDir)

		return errnew.
			Path.
			Error(
				errtype.ChmodApplyFailed,
				chmodErr,
				rootDir)
	}

	sliceErr := corestr.Empty.SimpleSlice()
	// for other os apply using each file.
	for _, filePath := range filePaths {
		err := fileMode.ApplyChmod(isSkipOnInvalid, filePath)

		if err != nil {
			sliceErr.Add(err.Error())
		}
	}

	toErr := errcore.SliceToError(sliceErr.Items)

	return errnew.
		Path.
		Error(
			errtype.ChmodApplyFailed,
			toErr,
			rootDir)
}
