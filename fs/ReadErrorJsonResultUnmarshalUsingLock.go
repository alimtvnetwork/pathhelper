package fs

import "gitlab.com/auk-go/errorwrapper"

func ReadErrorJsonResultUnmarshalUsingLock(
	filePath string,
	unmarshalObject interface{},
) *errorwrapper.Wrapper {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return ReadErrorJsonResultUnmarshal(filePath, unmarshalObject)
}
