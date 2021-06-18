package fsinternal

import (
	"io/ioutil"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func WriteFile(filePath string, content []byte) *errorwrapper.Wrapper {
	if content == nil {
		return errnew.MessagesPtr(
			errtype.NullOrEmptyReference,
			"Cannot write empty or nil contents to the file.",
			filePath)
	}

	chmod, err := chmodhelper.GetExistingChmod(filePath)
	if err != nil {
		return errnew.MessagesPtr(
			errtype.File,
			"fsinternal.WriteFile",
			err.Error())
	}

	writeErr := ioutil.WriteFile(
		filePath,
		content,
		chmod)

	if writeErr != nil {
		return errnew.MessagesPtr(
			errtype.FileWrite,
			"fsinternal.WriteFile",
			filePath,
			"Failed write file contents.",
			writeErr.Error())
	}

	return errnew.EmptyPtr
}
