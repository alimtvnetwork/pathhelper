package fs

import "gitlab.com/evatix-go/errorwrapper"

func AppendStringFileUsingLock(
	filePath string,
	content string,
) *errorwrapper.Wrapper {
	writerMutex.Lock()
	defer writerMutex.Unlock()

	return AppendFile(filePath, []byte(content))
}
