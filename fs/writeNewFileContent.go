package fs

import (
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/pathhelper/internal/consts"
)

func writeNewFileContent(
	isCreateParentDir bool,
	filePath string,
	content []byte,
) *errorwrapper.Wrapper {
	return writeNewFileContentUsingFileMode(
		isCreateParentDir,
		false,
		false,
		consts.DefaultDirMode,
		consts.DefaultFileMode,
		filePath,
		content,
	)
}
