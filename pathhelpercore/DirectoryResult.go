package pathhelpercore

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"

	"gitlab.com/evatix-go/pathhelper/enums"
)

type DirectoryResult struct {
	FileInfoWrapper   *FileInfoWrapper
	Error             errorwrapper.Wrapper
	RawPath           string
	FileModeRequested *os.FileMode
	HasIssues         bool
	IsIgnoredAction   bool
	Action            enums.PerformingAction
}

func NewEmptyDirectoryResult() *DirectoryResult {
	return &DirectoryResult{
		FileInfoWrapper:   nil,
		Error:             errnew.Empty,
		RawPath:           "",
		FileModeRequested: nil,
		HasIssues:         false,
		IsIgnoredAction:   true,
		Action:            enums.EmptyDirectoryResult,
	}
}
