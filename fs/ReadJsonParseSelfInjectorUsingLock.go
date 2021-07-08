package fs

import (
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/errorwrapper"
)

func ReadJsonParseSelfInjectorUsingLock(
	filePath string,
	jsonParseSelfInjector corejson.JsonParseSelfInjector,
) *errorwrapper.Wrapper {
	readWriteMutex.Lock()
	defer readWriteMutex.Unlock()

	return ReadJsonParseSelfInjector(
		filePath,
		jsonParseSelfInjector)
}
