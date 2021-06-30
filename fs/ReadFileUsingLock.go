package fs

import "gitlab.com/evatix-go/errorwrapper/errdata/errbyte"

func ReadFileUsingLock(filePath string) *errbyte.Results {
	readWriteMutex.Lock()
	defer readWriteMutex.Unlock()

	return ReadFile(filePath)
}
