package pathinsfmtexec

import (
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/pathhelper/envpath"
	"gitlab.com/auk-go/pathhelper/pathinsfmt"
)

func ApplyEnvPaths(baseEnvPaths *pathinsfmt.BaseEnvPaths) *errorwrapper.Wrapper {
	if baseEnvPaths == nil || len(baseEnvPaths.EnvPaths) == 0 {
		return nil
	}

	return envpath.AddOrUpdateEnvPaths(baseEnvPaths.EnvPaths...)
}
