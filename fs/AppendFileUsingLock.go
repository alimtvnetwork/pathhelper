package fs

import "gitlab.com/auk-go/errorwrapper"

func AppendFileUsingLock(
	isCreateParentDir bool,
	filePath string,
	content []byte,
) *errorwrapper.Wrapper {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return AppendFile(
		isCreateParentDir,
		filePath,
		content)
}
