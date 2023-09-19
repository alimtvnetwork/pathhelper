package envpath

import "gitlab.com/auk-go/errorwrapper"

func LinuxAddOrUpdatePtr(envPaths []string, isApplyEnvironmentSource bool) *errorwrapper.Wrapper {
	return linuxCrudEnvPath(
		linuxEnvAddOrUpdateAction,
		envPaths,
		isApplyEnvironmentSource)
}
