package fs

import "gitlab.com/evatix-go/errorwrapper"

func WriteFileLock(
	filePath string,
	content []byte,
) *errorwrapper.Wrapper {
	writerMutex.Lock()
	defer writerMutex.Unlock()

	return WriteFile(
		filePath,
		content)
}
