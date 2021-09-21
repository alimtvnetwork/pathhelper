package fs

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/pathhelper/internal/consts"
)

func writeNewFileContent(
	isCreateParentDir bool,
	filePath string,
	content []byte,
) *errorwrapper.Wrapper {
	return writeNewFileContentUsingFileMode(
		isCreateParentDir,
		filePath,
		content,
		consts.DefaultFileMode)
}
