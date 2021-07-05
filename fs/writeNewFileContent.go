package fs

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/pathhelper/internal/consts"
)

func writeNewFileContent(
	filePath string,
	content []byte,
) *errorwrapper.Wrapper {
	return writeNewFileContentUsingFileMode(
		filePath,
		content,
		consts.DefaultFileMode)
}
