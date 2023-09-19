package fs

import (
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/errorwrapper"
)

func ReadJsonParseSelfInjectorUsingLock(
	filePath string,
	jsonParseSelfInjector corejson.JsonParseSelfInjector,
) *errorwrapper.Wrapper {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return ReadJsonParseSelfInjector(
		filePath,
		jsonParseSelfInjector)
}
