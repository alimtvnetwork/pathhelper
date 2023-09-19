package fs

import (
	"gitlab.com/auk-go/errorwrapper"
)

func JsonReadUnmarshalOnExist(
	filePath string,
	unmarshallObjectRef interface{},
) *errorwrapper.Wrapper {
	if !IsPathExistsUsingLock(filePath) {
		return nil
	}

	return JsonReadUnmarshal(filePath, unmarshallObjectRef)
}
