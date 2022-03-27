package fs

import (
	"encoding/json"
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func JsonWriteMarshal(
	isCreateParentDir,
	isSkipOnNilObject bool,
	isKeepExistingFileModeOnExist bool,
	dirMode, fileMode os.FileMode,
	filePath string,
	marshallingObjectRef interface{},
) *errorwrapper.Wrapper {
	if marshallingObjectRef == nil && isSkipOnNilObject {
		return nil
	}

	allBytes, err := json.Marshal(marshallingObjectRef)

	if err != nil {
		return errnew.
			Path.
			Messages(
				errtype.Marshalling,
				filePath,
				err.Error(),
			)
	}

	return WriteFileUsingFileMode(
		isCreateParentDir,
		isKeepExistingFileModeOnExist,
		dirMode,
		fileMode,
		filePath,
		allBytes,
	)
}
