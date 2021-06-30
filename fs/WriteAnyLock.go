package fs

import (
	"gitlab.com/evatix-go/core/coreutils/stringutil"
	"gitlab.com/evatix-go/errorwrapper"
)

func WriteAnyLock(
	filePath string,
	content interface{},
) *errorwrapper.Wrapper {
	anyToString := stringutil.AnyToString(
		content)

	readWriteMutex.Lock()
	defer readWriteMutex.Unlock()

	return WriteFile(
		filePath,
		[]byte(anyToString))
}
