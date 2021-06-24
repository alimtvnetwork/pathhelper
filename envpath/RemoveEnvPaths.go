package envpath

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

func RemoveEnvPaths(removeEnvPaths ...string) *errorwrapper.Wrapper {
	if len(removeEnvPaths) == 0 {
		return errnew.EmptyPtr
	}

	return RemoveEnvPathsPtr(&removeEnvPaths)
}
