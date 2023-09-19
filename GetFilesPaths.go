package pathhelper

import (
	"gitlab.com/auk-go/core/osconsts"

	"gitlab.com/auk-go/pathhelper/fileinfo"
)

// GetFileNames
//
// returns file names on the path (non-lazy execution).
func GetFileNames(path string, isNormalize bool) *fileinfo.FileNamesCollection {
	return fileinfo.NewFileNamesUsing(
		path,
		osconsts.PathSeparator,
		isNormalize)
}
