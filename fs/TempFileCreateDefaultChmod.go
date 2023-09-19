package fs

import (
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/pathhelper/internal/consts"
)

func TempFileCreateDefaultChmod(
	existingFileErrorWrapper *errorwrapper.Wrapper, // can be nil, current function return error
	relativeFilePath string,
) *OsFile {
	return TempFileCreate(
		existingFileErrorWrapper,
		relativeFilePath,
		consts.DefaultFileMode)
}
