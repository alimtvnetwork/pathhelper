package fs

import "gitlab.com/evatix-go/errorwrapper"

func JsonReadUnmarshalLock(
	filePath string,
	unmarshallObjectRef interface{},
) *errorwrapper.Wrapper {
	readWriteMutex.Lock()
	defer readWriteMutex.Unlock()

	return JsonReadUnmarshal(
		filePath,
		unmarshallObjectRef)
}
