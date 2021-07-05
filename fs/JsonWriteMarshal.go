package fs

import (
	"encoding/json"
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func JsonWriteMarshal(
	isSkipOnNilObject bool,
	filePath string,
	fileMod os.FileMode,
	isKeepExistingFileModeOnExist bool,
	marshallingObjectRef interface{},
) *errorwrapper.Wrapper {
	if marshallingObjectRef == nil && isSkipOnNilObject {
		return errnew.EmptyPtr
	}

	allBytes, err := json.Marshal(marshallingObjectRef)

	if err != nil {
		return errnew.PathMessages(
			errtype.Marshalling,
			filePath,
			err.Error(),
		)
	}

	return WriteFileUsingFileMode(
		filePath,
		allBytes,
		fileMod,
		isKeepExistingFileModeOnExist)
}
