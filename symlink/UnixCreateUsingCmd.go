package symlink

import (
	"gitlab.com/auk-go/core/codestack"
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/osconsts"
	"gitlab.com/auk-go/errorwrapper/errcmd"
	"gitlab.com/auk-go/errorwrapper/errdata/errbool"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
)

// Creates symbolicLink of the source at the provided destination path for linux system. If destination doesn't exist it will panic.
// sourcePath example: "/home/a/test.txt"; destinationPath example: "/home/a/go/test.txt"
// destination need to have read and write permission for the user.
func UnixCreateUsingCmd(sourcePath, destinationPath string) *errbool.Result {
	if osconsts.IsWindows {
		return errbool.New.Result.ErrorWrapper(
			errnew.
				Type.
				UsingStackSkip(
					codestack.Skip1,
					errtype.NotSupportInWindows))
	}

	symLink := errcmd.ArgsJoin(
		constants.SymbolicLinkCreationCommandName,
		constants.SymbolicLinkCreationArgument,
		sourcePath,
		destinationPath)

	cmdOnceResult := errcmd.New.BashScript.LinesResult(symLink)

	return errbool.New.Result.Create(
		cmdOnceResult.IsEmptyError(),
		cmdOnceResult.ErrorWrapper(),
	)
}
