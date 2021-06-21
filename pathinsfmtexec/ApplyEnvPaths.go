package pathinsfmtexec

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/pathhelper/envpath"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func ApplyEnvPaths(baseEnvPaths *pathinsfmt.BaseEnvPaths) *errorwrapper.Wrapper {
	if baseEnvPaths == nil || len(baseEnvPaths.EnvPaths) == 0 {
		return errnew.EmptyPtr
	}

	return envpath.AddOrUpdateEnvPaths(baseEnvPaths.EnvPaths...)
}



func ApplySymbolicLink(symbolicLink *pathinsfmt.SymbolicLink) *errorwrapper.Wrapper {
	if symbolicLink == nil {
		return errnew.EmptyPtr
	}

	return sym(baseEnvPaths.EnvPaths...)
}
