package envpath

import (
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

func AddOrUpdateEnvPathsPtr(addOrUpdateEnvPaths *[]string) *errorwrapper.Wrapper {
	if addOrUpdateEnvPaths == nil || len(*addOrUpdateEnvPaths) == 0 {
		return errnew.EmptyPtr
	}

	envPaths := ReadEnvPathsPtr()
	hashset := corestr.NewHashsetUsingStrings(
		envPaths)

	hashset.AddStringsPtr(addOrUpdateEnvPaths)
	compiledPath := hashsetEnvPathToSingleString(hashset)

	return SetEnvPath(compiledPath)
}
