package pathhelpercore

import (
	"os"

	"gitlab.com/evatix-go/pathhelper/enums"
)

type DirectoryResult struct {
	FileInfoWrapper   *FileInfoWrapper
	Error             *error
	RawPath           string
	FileModeRequested os.FileMode
	HasIssues         bool
	IsIgnoredAction   bool
	Action            enums.PerformingAction
}
