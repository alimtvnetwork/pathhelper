package envvars

import (
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/pathhelper/pathinsfmt"
)

func Apply(environmentVariables *pathinsfmt.BaseEnvironmentVariables) *errorwrapper.Wrapper {
	if environmentVariables == nil || environmentVariables.IsEmptyEnvVars() {
		return nil
	}

	for i := range environmentVariables.EnvVars {
		errWrap := ApplyEnvVar(&environmentVariables.EnvVars[i])
		if errWrap.HasError() {
			return errWrap
		}
	}

	return nil
}
