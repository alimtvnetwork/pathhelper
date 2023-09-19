package fs

import "gitlab.com/auk-go/errorwrapper"

func JsonReadUnmarshalLock(
	filePath string,
	unmarshallObjectRef interface{},
) *errorwrapper.Wrapper {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return JsonReadUnmarshal(
		filePath,
		unmarshallObjectRef)
}
