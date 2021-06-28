package fs

import (
	"io/ioutil"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/filemode"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func WriteFile(
	filePath string,
	content []byte,
) *errorwrapper.Wrapper {
	if content == nil {
		return errnew.MessagesPtr(
			errtype.NullOrEmptyReference,
			"Cannot write empty or nil contents to the file.",
			filePath)
	}

	if IsPathExists(filePath) {
		chmod, err := chmodhelper.GetExistingChmod(filePath)
		if err != nil {
			return errnew.MessagesPtr(
				errtype.File,
				"fs.WriteFile",
				err.Error())
		}

		writeErr := ioutil.WriteFile(
			filePath,
			content,
			chmod)

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
