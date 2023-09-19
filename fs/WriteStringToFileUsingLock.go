package fs

import "gitlab.com/auk-go/errorwrapper"

func WriteStringToFileUsingLock(
	isCreateParentDir bool,
	filePath string,
	content string,
) *errorwrapper.Wrapper {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return WriteStringToFile(
		isCreateParentDir,
		filePath,
		content)
}
