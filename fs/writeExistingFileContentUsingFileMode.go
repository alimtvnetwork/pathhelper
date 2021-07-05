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
	mode os.FileMode,
) *errorwrapper.Wrapper {
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
