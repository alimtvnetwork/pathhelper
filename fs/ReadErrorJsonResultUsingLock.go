package fs

import "gitlab.com/evatix-go/errorwrapper/errdata/errjson"

func ReadErrorJsonResultUsingLock(filePath string) *errjson.Result {
	readWriteMutex.Lock()
	defer readWriteMutex.Unlock()

	return ReadErrorJsonResult(filePath)
}
