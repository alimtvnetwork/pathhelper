package envpath

import "gitlab.com/auk-go/errorwrapper"

func LinuxRemovePtr(envPaths []string, isApplyEnvironmentSource bool) *errorwrapper.Wrapper {
	return linuxCrudEnvPath(
		linuxEnvRemoveAction,
		envPaths,
		isApplyEnvironmentSource)
}
