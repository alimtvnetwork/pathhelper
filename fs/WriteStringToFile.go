package fs

import (
	"gitlab.com/auk-go/errorwrapper"
)

func WriteStringToFile(
	isCreateParentDir bool,
	filePath string,
	content string,
) *errorwrapper.Wrapper {
	return WriteFile(
		isCreateParentDir,
		filePath,
		[]byte(content))
}
