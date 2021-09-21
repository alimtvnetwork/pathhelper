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
	filePath string,
	content []byte,
	mode os.FileMode,
) *errorwrapper.Wrapper {
	createDirErr := errnew.EmptyPtr
	if isCreateParentDir {
		createDirErr = createdir.AllUptoParentDefault(
			filePath)
	}

	if createDirErr.HasError() {
		return createDirErr
	}

	writeErr := ioutil.WriteFile(
		filePath,
		content,
		mode)

	if writeErr != nil {
		return errnew.PathMessages(
			errtype.FileWrite,
			filePath,
			"fs.WriteFile",
			"Failed write file contents.",
			writeErr.Error())
	}

	return errnew.EmptyPtr
}
