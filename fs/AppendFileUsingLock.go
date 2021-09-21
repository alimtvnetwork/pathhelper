package fs

import "gitlab.com/evatix-go/errorwrapper"

func AppendFileUsingLock(
	isCreateParentDir bool,
	filePath string,
	content []byte,
) *errorwrapper.Wrapper {
	readWriteMutex.Lock()
	defer readWriteMutex.Unlock()

	return AppendFile(
		isCreateParentDir,
		filePath,
		content)
}
