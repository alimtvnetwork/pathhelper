package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/fileinfo"
)

// returns file names on the path (non-lazy execution).
func GetFileNames(path string) *fileinfo.FileNamesCollection {
	return fileinfo.NewFileNamesUsing(path)
}
