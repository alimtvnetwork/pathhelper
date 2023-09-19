package fsinternal

import (
	"gitlab.com/auk-go/core/filemode"
	"gitlab.com/auk-go/errorwrapper"
)

// WriteFileDefault
//
//	Default chmod dir - 0755, file - 0644
func WriteFileDefault(
	filePath string,
	content []byte,
) *errorwrapper.Wrapper {
	return WriteFile(
		filemode.DirDefault,
		filemode.FileDefault,
		filePath,
		content)
}
