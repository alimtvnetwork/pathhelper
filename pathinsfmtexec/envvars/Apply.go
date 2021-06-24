package envvars

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func Apply(environmentVariables *pathinsfmt.BaseEnvironmentVariables) *errorwrapper.Wrapper {
	if environmentVariables == nil || environmentVariables.IsEmptyEnvVars() {
		return errnew.EmptyPtr
	}

	for _, envVar := range environmentVariables.EnvVars {
		errW := ApplyEnvVar(&envVar)
		if errW.HasError() {
			return errW
		}
	}

	return errnew.EmptyPtr
}
