package fs

import (
	"io/ioutil"

	"gitlab.com/evatix-go/core/filemode"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func writeNewFileContent(
	filePath string,
	content []byte,
) *errorwrapper.Wrapper {
	writeErr := ioutil.WriteFile(
		filePath,
		content,
		filemode.X644)

	if writeErr != nil {
		return errnew.MessagesPtr(
			errtype.FileWrite,
			"fs.WriteFile",
			filePath,
			"Failed write file contents.",
			writeErr.Error())
	}

	return errnew.EmptyPtr
}
