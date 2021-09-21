package fs

import "gitlab.com/evatix-go/errorwrapper"

func WriteStringToFileUsingLock(
	isCreateParentDir bool,
	filePath string,
	content string,
) *errorwrapper.Wrapper {
	readWriteMutex.Lock()
	defer readWriteMutex.Unlock()

	return WriteStringToFile(
		isCreateParentDir,
		filePath,
		content)
}
