package fs

import (
	"gitlab.com/auk-go/errorwrapper"
)

func JsonReadUnmarshal(
	filePath string,
	toPointer interface{},
) *errorwrapper.Wrapper {
	readContents := ReadFile(filePath)

	if readContents.HasError() {
		return readContents.ErrorWrapper
	}

	return readContents.Deserialize(
		toPointer)
}
