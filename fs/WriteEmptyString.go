package fs

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/errorwrapper"
)

func WriteEmptyString(
	isCreateParentDir bool,
	filePath string,
) *errorwrapper.Wrapper {
	return WriteFile(
		isCreateParentDir,
		filePath,
		[]byte(constants.EmptyString))
}
