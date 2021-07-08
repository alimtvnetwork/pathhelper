package fs

import "gitlab.com/evatix-go/errorwrapper"

func ReadErrorJsonResultUnmarshalUsingLock(
	filePath string,
	unmarshalObject interface{},
) *errorwrapper.Wrapper {
	readWriteMutex.Lock()
	defer readWriteMutex.Unlock()

	return ReadErrorJsonResultUnmarshal(filePath, unmarshalObject)
}
