package fs

import (
	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func appendFileContent(filePath string, content []byte) (finalErrorWrapper *errorwrapper.Wrapper) {
	chmod, err := chmodhelper.GetExistingChmod(filePath)
	if err != nil {
		return errnew.PathMessages(
			errtype.ExistingChmodReadFailed,
			filePath,
			err.Error())
	}

	osFile := GetOsFile(
		finalErrorWrapper,
		filePath,
		FlagAppendOrWrite,
		chmod)

	if osFile.HasError() {
		return osFile.ErrorWrapper
	}

	defer osFile.AttachDeferCloseOnRequire()
	_, appendingErr := osFile.OsFile.Write(content)

	if appendingErr != nil {
		return errnew.PathMessages(
			errtype.FileAppend,
			filePath,
			"fs.appendFileContent",
			"Failed append file contents.",
			appendingErr.Error())
	}

	return errnew.EmptyPtr
}
