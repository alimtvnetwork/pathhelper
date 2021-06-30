package fs

import "gitlab.com/evatix-go/errorwrapper"

func AppendStringFileUsingLock(
	filePath string,
	content string,
) *errorwrapper.Wrapper {
	readWriteMutex.Lock()
	defer readWriteMutex.Unlock()

	return AppendFile(
		filePath,
		[]byte(content))
}
