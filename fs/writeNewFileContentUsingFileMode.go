package fs

import (
	"io/ioutil"
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/createdir"
)

func writeNewFileContentUsingFileMode(
	isCreateParentDir bool,
	dirMode, fileMode os.FileMode,
	filePath string,
	content []byte,
) *errorwrapper.Wrapper {
	var createDirErr *errorwrapper.Wrapper
	if isCreateParentDir {
		createDirErr = createdir.AllUptoParent(
			filePath,
			dirMode)
	}

	if createDirErr.HasError() {
		return createDirErr
	}

	writeErr := ioutil.WriteFile(
		filePath,
		content,
		fileMode)

	if writeErr != nil {
		return errnew.Path.
			Messages(
				errtype.FileWrite,
				filePath,
				"fs.WriteFile",
				"Failed write file contents.",
				writeErr.Error())
	}

	return nil
}
