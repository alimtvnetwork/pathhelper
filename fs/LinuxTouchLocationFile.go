package fs

import (
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/pathhelper/pathjoin"
)

func LinuxTouchLocationFile(parentPath, fileName string) *errorwrapper.Wrapper {
	joinedPath := pathjoin.JoinNormalized(
		parentPath,
		fileName)

	return LinuxTouchFile(joinedPath)
}
