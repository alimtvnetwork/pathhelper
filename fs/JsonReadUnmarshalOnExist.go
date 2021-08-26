package fs

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

func JsonReadUnmarshalOnExist(
	filePath string,
	unmarshallObjectRef interface{},
) *errorwrapper.Wrapper {
	if !IsPathExistsUsingLock(filePath) {
		return errnew.EmptyPtr
	}

	return JsonReadUnmarshal(filePath, unmarshallObjectRef)
}
