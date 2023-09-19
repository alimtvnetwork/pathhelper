package pathchmod

import (
	"os"

	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
)

func ApplyLinuxRecursiveChmodOnPathUsingFileMode(
	mode os.FileMode,
	rootDir string,
) *errorwrapper.Wrapper {
	fileMode := chmodhelper.
		New.
		RwxWrapper.
		UsingFileMode(mode)

	chmodErr := fileMode.
		LinuxApplyRecursive(false, rootDir)

	return errnew.
		Path.
		Error(errtype.ChmodApplyFailed, chmodErr, rootDir)
}
