package envpath

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

func AddOrUpdateEnvPaths(addOrUpdateEnvPaths ...string) *errorwrapper.Wrapper {
	if len(addOrUpdateEnvPaths) == 0 {
		return errnew.EmptyPtr
	}

	return AddOrUpdateEnvPathsPtr(&addOrUpdateEnvPaths)
}
