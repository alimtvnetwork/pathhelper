package fs

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/errorwrapper"
)

func WriteStringLinesToFileSkipOnEmptyUsingLock(
	isCreateParentDir bool,
	filePath string,
	contentLines []string,
) *errorwrapper.Wrapper {
	if len(contentLines) == 0 {
		return nil
	}

	content := strings.Join(
		contentLines,
		constants.NewLineUnix)

	return WriteFileLock(
		isCreateParentDir,
		filePath,
		[]byte(content))
}
