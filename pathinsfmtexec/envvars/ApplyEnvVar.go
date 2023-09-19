package envvars

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errcmd"
	"gitlab.com/auk-go/pathhelper/internal/consts"
	"gitlab.com/auk-go/pathhelper/pathinsfmt"
)

func ApplyEnvVar(
	environmentVariable *pathinsfmt.EnvironmentVariable,
) *errorwrapper.Wrapper {
	if environmentVariable == nil {
		return nil
	}

	variableKeyValueAttach :=
		environmentVariable.Name +
			constants.EqualSymbol +
			environmentVariable.Value

	cmdResult := errcmd.New.BashScript.ArgsDefault(
		consts.Export,
		variableKeyValueAttach,
	)

	return cmdResult.CompiledErrorWrapper()
}
