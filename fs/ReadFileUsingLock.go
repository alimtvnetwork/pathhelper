package fs

import "gitlab.com/auk-go/errorwrapper/errdata/errbyte"

func ReadFileUsingLock(filePath string) *errbyte.Results {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return ReadFile(filePath)
}
