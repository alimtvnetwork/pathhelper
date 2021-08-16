package fs

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper"
)

func WriteEmptyStringLock(
	filePath string,
) *errorwrapper.Wrapper {
	return WriteFileLock(
		filePath,
		[]byte(constants.EmptyString))
}
