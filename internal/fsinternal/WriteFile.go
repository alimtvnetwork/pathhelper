package fsinternal

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
		return NullContentErrorWrap(
			filePath)
	}

	if IsPathExists(filePath) {
		chmod, err := chmodhelper.GetExistingChmod(filePath)
		if err != nil {
			return errnew.Messages.Many(
				errtype.File,
				"fsinternal.WriteFile",
				err.Error())
		}

		writeErr := ioutil.WriteFile(
			filePath,
			content,
			chmod)

		if writeErr != nil {
			return errnew.Messages.Many(
				errtype.FileWrite,
				"fsinternal.WriteFile",
				filePath,
				"Failed write file contents.",
				writeErr.Error())
		}

		return nil
	}

	writeErr := ioutil.WriteFile(
		filePath,
		content,
		filemode.X644)

	if writeErr != nil {
		return errnew.Messages.Many(
			errtype.FileWrite,
			"fsinternal.WriteFile",
			filePath,
			"Failed write file contents.",
			writeErr.Error())
	}

	return nil
}
