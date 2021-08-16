package fs

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/errorwrapper"
)

func WriteSimpleSliceToFileUsingLock(
	filePath string,
	simpleSlice *corestr.SimpleSlice,
) *errorwrapper.Wrapper {
	if simpleSlice.IsEmpty() {
		return WriteEmptyStringLock(
			filePath)
	}

	content := strings.Join(
		simpleSlice.Items,
		constants.NewLineUnix)

	return WriteFileLock(
		filePath,
		[]byte(content))
}
