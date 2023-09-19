package fs

import (
	"gitlab.com/auk-go/errorwrapper/errdata/errbyte"
	"gitlab.com/auk-go/pathhelper/pathjoin"
)

func ReadUsingLock(location string, filename string) *errbyte.Results {
	compiledFilePath := pathjoin.JoinNormalized(location, filename)

	globalMutex.Lock()
	defer globalMutex.Unlock()

	return ReadFile(compiledFilePath)
}
