package fsinternal

import (
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/pathhelper/pathjoin"
)

func Write(
	location string,
	filename string,
	content []byte,
) *errorwrapper.Wrapper {
	compileFilePath := pathjoin.JoinNormalized(
		location,
		filename) // todo check how to remove this reference

	return WriteFileDefault(
		compileFilePath, content)
}
