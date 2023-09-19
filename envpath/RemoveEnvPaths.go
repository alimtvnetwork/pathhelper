package envpath

import (
	"gitlab.com/auk-go/errorwrapper"
)

func RemoveEnvPaths(removeEnvPaths ...string) *errorwrapper.Wrapper {
	if len(removeEnvPaths) == 0 {
		return nil
	}

	return RemoveEnvPathsPtr(&removeEnvPaths)
}
