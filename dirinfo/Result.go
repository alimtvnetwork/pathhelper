package fileinfo

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"

	"gitlab.com/evatix-go/pathhelper/enums"
	"gitlab.com/evatix-go/pathhelper/fileinfo"
)

type Result struct {
	FileInfoWrapper   *fileinfo.Wrapper
	Error             errorwrapper.Wrapper
	RawPath           string
	FileModeRequested *os.FileMode
	HasIssues         bool
	IsIgnoredAction   bool
	Action            enums.PerformingAction
}

func NewEmptyDirectoryResult() *Result {
	return &Result{
		FileInfoWrapper:   nil,
		Error:             errnew.Empty,
		RawPath:           "",
		FileModeRequested: nil,
		HasIssues:         false,
		IsIgnoredAction:   true,
		Action:            enums.EmptyDirectoryResult,
	}
}
