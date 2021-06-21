package pathchmod

import (
	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/enums/scripttype"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errcmd"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/internal/cmdprefix"
)

// ChmodApply Format : chmod -R 777 /var/lib/powerdns
func ChmodApply(
	isRecursive bool,
	isSkipOnError bool,
	wrapper *chmodhelper.RwxWrapper,
	paths ...string,
) *errorwrapper.Wrapper {
	length := len(paths)

	if length == 0 {
		return errnew.EmptyPtr
	}

	chmodPrefix := cmdprefix.Chmod(isRecursive, wrapper)

	if isSkipOnError {
		errCollection := errwrappers.Empty()

		for _, currentPath := range paths {
			command := errcmd.ArgsJoin(chmodPrefix, currentPath)
			errCollection.AddScriptErrors(scripttype.Bash, command)
		}

		return errCollection.GetAsErrorWrapperPtr()
	}

	for _, currentPath := range paths {
		command := errcmd.ArgsJoin(chmodPrefix, currentPath)
		compiledResult := errcmd.BashScripts(command).CompiledResult()

		if compiledResult.HasError() {
			return compiledResult.ErrorWrapper()
		}
	}

	return errnew.EmptyPtr
}
