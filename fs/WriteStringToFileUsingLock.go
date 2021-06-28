package fs

import "gitlab.com/evatix-go/errorwrapper"

func WriteStringToFileUsingLock(
	filePath string,
	content string,
) *errorwrapper.Wrapper {
	writerMutex.Lock()
	defer writerMutex.Unlock()

	return WriteStringToFile(
		filePath,
		content)
}
