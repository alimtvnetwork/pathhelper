package fs

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/errorwrapper"
)

func WriteEmptyStringLock(
	isCreateParentDir bool,
	filePath string,
) *errorwrapper.Wrapper {
	return WriteFileLock(
		isCreateParentDir,
		filePath,
		[]byte(constants.EmptyString))
}
