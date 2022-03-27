package fs

import (
	"io/ioutil"
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func writeExistingFileContentUsingFileMode(
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

	return nil
}
