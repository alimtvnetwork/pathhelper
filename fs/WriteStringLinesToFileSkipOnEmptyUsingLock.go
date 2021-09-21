package fs

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

func WriteStringLinesToFileSkipOnEmptyUsingLock(
	isCreateParentDir bool,
	filePath string,
	contentLines []string,
) *errorwrapper.Wrapper {
	if len(contentLines) == 0 {
		return errnew.EmptyPtr
	}

	content := strings.Join(
		contentLines,
		constants.NewLineUnix)

	return WriteFileLock(
		isCreateParentDir,
		filePath,
		[]byte(content))
}
