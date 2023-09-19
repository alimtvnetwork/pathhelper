package fsinternal

import (
	"gitlab.com/auk-go/errorwrapper/errdata/errbyte"
	"gitlab.com/auk-go/pathhelper/pathjoin"
)

func Read(location string, filename string) *errbyte.Results {
	compiledFilePath := pathjoin.JoinNormalized(location, filename)

	return ReadFile(compiledFilePath)
}
