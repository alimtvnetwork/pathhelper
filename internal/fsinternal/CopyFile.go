package fsinternal

import (
	"gitlab.com/evatix-go/errorwrapper"
)

// CopyFile Future ref: https://stackoverflow.com/a/21067803
func CopyFile(srcPath, dstPath string) *errorwrapper.Wrapper {
	return CopyFileContents(srcPath, dstPath)
}
