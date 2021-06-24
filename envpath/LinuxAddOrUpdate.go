package envpath

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

func LinuxAddOrUpdate(isApplyEnvironmentSource bool, envPaths ...string) *errorwrapper.Wrapper {
	if len(envPaths) == 0 {
		return errnew.EmptyPtr
	}

	return LinuxAddOrUpdatePtr(&envPaths, isApplyEnvironmentSource)
}
