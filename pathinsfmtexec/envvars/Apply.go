package envvars

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func Apply(environmentVariables *pathinsfmt.BaseEnvironmentVariables) *errorwrapper.Wrapper {
	if environmentVariables == nil || environmentVariables.IsEmptyEnvVars() {
		return nil
	}

	for i := range environmentVariables.EnvVars {
		errW := ApplyEnvVar(&environmentVariables.EnvVars[i])
		if errW.HasError() {
			return errW
		}
	}

	return nil
}
