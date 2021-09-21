package fs

import "gitlab.com/evatix-go/errorwrapper"

func AppendStringFileUsingLock(
	isCreateParentDir bool,
	filePath string,
	content string,
) *errorwrapper.Wrapper {
	readWriteMutex.Lock()
	defer readWriteMutex.Unlock()

	return AppendFile(
		isCreateParentDir,
		filePath,
		[]byte(content))
}
