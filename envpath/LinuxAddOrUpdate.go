package envpath

import (
	"gitlab.com/auk-go/errorwrapper"
)

func LinuxAddOrUpdate(isApplyEnvironmentSource bool, envPaths ...string) *errorwrapper.Wrapper {
	if len(envPaths) == 0 {
		return nil
	}

	return LinuxAddOrUpdatePtr(envPaths, isApplyEnvironmentSource)
}
