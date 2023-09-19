package fs

import (
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/pathhelper/pathjoin"
)

func Write(
	isCreateParentDir bool,
	location string,
	filename string,
	content []byte,
) *errorwrapper.Wrapper {
	compileFilePath := pathjoin.JoinNormalized(location, filename)

	return WriteFile(
		isCreateParentDir,
		compileFilePath,
		content)
}
