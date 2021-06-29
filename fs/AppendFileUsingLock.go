package fs

import "gitlab.com/evatix-go/errorwrapper"

func AppendFileUsingLock(
	filePath string,
	content []byte,
) *errorwrapper.Wrapper {
	writerMutex.Lock()
	defer writerMutex.Unlock()

	return AppendFile(filePath, content)
}
