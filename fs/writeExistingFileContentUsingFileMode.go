package fs

import (
	"io/ioutil"
	"os"

	"gitlab.com/auk-go/core/osconsts"
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
	"gitlab.com/auk-go/pathhelper/pathchmod"
)

func writeExistingFileContentUsingFileMode(
	isApplyChmodOnMismatchOnly bool,
	filePath string,
	content []byte,
	fileMode os.FileMode,
) *errorwrapper.Wrapper {
	writeErr := ioutil.WriteFile(
		filePath,
		content,
		fileMode)

	if writeErr != nil {
		return errnew.
			Path.
			Messages(
				errtype.FileWrite,
				filePath,
				"fs.WriteFile",
				"Failed write file contents.",
				writeErr.Error())
	}

	if osconsts.IsWindows {
		return nil
	}

	return pathchmod.ApplyOnOptions(
		true,
		false,
		isApplyChmodOnMismatchOnly,
		fileMode,
		filePath)
}
