package fs

import (
	"encoding/json"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func JsonReadUnmarshal(
	filePath string,
	unmarshallObjectRef interface{},
) *errorwrapper.Wrapper {
	readContents := ReadFile(filePath)

	if readContents.HasError() {
		return readContents.ErrorWrapper
	}

	err := json.Unmarshal(
		*readContents.Values,
		unmarshallObjectRef)

	if err != nil {
		return errnew.PathMessages(
			errtype.Unmarshalling,
			filePath,
			err.Error(),
			"Contents:\n",
			readContents.String())
	}

	return errnew.EmptyPtr
}
