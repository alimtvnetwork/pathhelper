package envpath

import (
	"gitlab.com/auk-go/core/osconsts"
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
	"gitlab.com/auk-go/pathhelper/internal/fsinternal"
	"gitlab.com/auk-go/pathhelper/internal/messages"
)

// linuxCrudEnvPath
//
// linuxEnvPathCrudFunc could be -
//   - linuxEnvRemoveAction or
//   - linuxEnvAddOrUpdateAction
func linuxCrudEnvPath(
	performingFunc linuxEnvPathCrudFunc,
	envPaths []string,
	isApplyEnvironmentSource bool,
) *errorwrapper.Wrapper {
	if len(envPaths) == 0 {
		return nil
	}

	if !osconsts.IsUnixGroup || !osconsts.IsLinux {
		return errnew.Messages.Many(
			errtype.NotSupportOperatingSystem,
			"linuxCrudEnvPath",
			messages.CannotAddUpdateRemoveEnvPath)
	}

	contentsBytes := fsinternal.ReadFile(etcEnvPath)

	if contentsBytes.HasError() {
		return contentsBytes.ErrorWrapper
	}

	envPathString := contentsBytes.String()

	newEnvPathCompiled := performingFunc(envPaths, envPathString)
	prependPathEqual := compileEnvPathToLinuxRawEnvPathFormat(
		newEnvPathCompiled)

	writeErrorWrapper := fsinternal.WriteStringToFile(
		etcEnvPath,
		prependPathEqual)

	if writeErrorWrapper.HasError() {
		return writeErrorWrapper
	}

	if setEnvErr := SetEnvPath(newEnvPathCompiled); setEnvErr.HasError() {
		return setEnvErr
	}

	if isApplyEnvironmentSource {
		return LinuxApplySourceEnvironment()
	}

	return nil
}
