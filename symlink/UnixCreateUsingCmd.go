package symlink

import (
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/errorwrapper/errcmd"
	"gitlab.com/evatix-go/errorwrapper/errdata/errbool"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/internal/argsinternal"

	"gitlab.com/evatix-go/core/constants"
)

// Creates symbolicLink of the source at the provided destination path for linux system. If destination doesn't exist it will panic.
// sourcePath example: "/home/a/test.txt"; destinationPath example: "/home/a/go/test.txt"
// destination need to have read and write permission for the user.
func UnixCreateUsingCmd(sourcePath, destinationPath string) *errbool.Result {
	if osconsts.IsWindows {
		return errbool.NewSimplePtr(false, errtype.NotSupportInWindows)
	}

	symLink := argsinternal.Join(constants.SymbolicLinkCreationCommandName,
		constants.SymbolicLinkCreationArgument,
		sourcePath,
		destinationPath)

	cmdOnceResult := errcmd.BashScripts(symLink).CompiledResult()

	return &errbool.Result{
		Value:        cmdOnceResult.IsEmptyError(),
		ErrorWrapper: cmdOnceResult.ErrorWrapper(),
	}
}
