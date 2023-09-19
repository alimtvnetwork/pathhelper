package copyrecursive

import (
	"io/ioutil"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"

	"gitlab.com/auk-go/pathhelper/pathsconst"
)

// CopyToTempDir copies everything from root to a
// random tmp directory and returns the tmp root path
func CopyToTempDir(root string) (string, *errorwrapper.Wrapper) {
	tmpRoot, err := ioutil.TempDir(
		constants.EmptyString,
		pathsconst.TestDirPatternName)
	if err != nil {
		return constants.EmptyString, errnew.Path.Error(
			errtype.Copy, err, root)
	}

	errWrap := NewCopier(root, tmpRoot, Options{
		IsSkipOnExist:      false,
		IsRecursive:        true,
		IsClearDestination: false,
		IsUseShellOrCmd:    false,
		IsNormalize:        false,
		IsExpandVar:        false,
	}).Copy()

	return tmpRoot, errWrap
}
