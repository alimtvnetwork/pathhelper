package fs

import (
	"os"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func appendFileContent(filePath string, content []byte) *errorwrapper.Wrapper {
	chmod, err := chmodhelper.GetExistingChmod(filePath)
	if err != nil {
		return errnew.MessagesPtr(
			errtype.File,
			"fs.WriteFile",
			err.Error())
	}

	appendingFile, appendingFileErr := os.OpenFile(
		"temp.txt",
		os.O_APPEND|os.O_WRONLY,
		chmod)

	if appendingFile != nil {
		defer appendingFile.Close()
	}

	if appendingFileErr != nil {
		return errnew.MessagesPtr(
			errtype.FileRead,
			"fs.AppendFile",
			filePath,
			"Failed append file.",
			appendingFileErr.Error())
	}

	_, appendingErr := appendingFile.Write(content)

	if appendingErr != nil {
		return errnew.MessagesPtr(
			errtype.FileAppend,
			"fs.AppendFile",
			filePath,
			"Failed append file contents.",
			appendingErr.Error())
	}

	return errnew.EmptyPtr
}
