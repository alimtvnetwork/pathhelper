package fs

import "gitlab.com/auk-go/errorwrapper"

func AppendStringFileUsingLock(
	isCreateParentDir bool,
	filePath string,
	content string,
) *errorwrapper.Wrapper {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return AppendFile(
		isCreateParentDir,
		filePath,
		[]byte(content))
}
