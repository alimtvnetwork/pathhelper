package fs

import "gitlab.com/evatix-go/errorwrapper"

func WriteStringToFileUsingLock(
	filePath string,
	content string,
) *errorwrapper.Wrapper {
	readWriteMutex.Lock()
	defer readWriteMutex.Unlock()

	return WriteStringToFile(
		filePath,
		content)
}
