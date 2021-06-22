package envvars

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errcmd"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/pathhelper/internal/consts"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func ApplyEnvVar(
	environmentVariable *pathinsfmt.EnvironmentVariable,
) *errorwrapper.Wrapper {
	if environmentVariable == nil {
		return errnew.EmptyPtr
	}

	variableKeyValueAttach :=
		environmentVariable.Name +
			constants.EqualSymbol +
			environmentVariable.Value

	cmdResult := errcmd.BashArgs(
		consts.Export,
		variableKeyValueAttach,
	)

	return cmdResult.CompiledErrorWrapper()
}
