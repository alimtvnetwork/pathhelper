package envpath

import (
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

func RemoveEnvPathsPtr(removeEnvPaths *[]string) *errorwrapper.Wrapper {
	if removeEnvPaths == nil || len(*removeEnvPaths) == 0 {
		return errnew.EmptyPtr
	}

	envPaths := ReadEnvPathsPtr()
	hashset := corestr.NewHashsetUsingStrings(
		envPaths)

	for _, removeEnvPath := range *removeEnvPaths {
		hashset.Remove(removeEnvPath)
	}

	compiledPath := hashsetEnvPathToSingleString(hashset)

	return SetEnvPath(compiledPath)
}
