package fs

import (
	"gitlab.com/evatix-go/errorwrapper"
)

func WriteFileLock(
	isCreateParentDir bool,
	filePath string,
	content []byte,
) *errorwrapper.Wrapper {
	readWriteMutex.Lock()
	defer readWriteMutex.Unlock()

	return WriteFile(
		isCreateParentDir,
		filePath,
		content)
}
