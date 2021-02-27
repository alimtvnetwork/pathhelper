package pathhelper

import (
	"os/exec"

	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/errorwrapper/errdata/errbool"
	"gitlab.com/evatix-go/errorwrapper/errtype"

	"gitlab.com/evatix-go/core/constants"
)

// Creates symbolicLink of the source at the provided destination path for linux system. If destination doesn't exist it will panic.
// sourcePath example: "/home/a/test.txt"; destinationPath example: "/home/a/go/test.txt"
// destination need to have read and write permission for the user.
func CreateSymbolicLinkLinux(sourcePath, destinationPath string) *errbool.Result {
	if osconsts.IsWindows {
		return errbool.
			EmptyErrorResultPtr(false)
	}

	cmd := exec.Command(
		constants.SymbolicLinkCreationCommandName,
		constants.SymbolicLinkCreationArgument,
		sourcePath,
		destinationPath)

	if cmd == nil {
		return errbool.
			NewSimplePtr(
				false,
				errtype.CommandExecutionNotFound)
	}

	_, err := cmd.Output()

	if err != nil {
		resultWithError := errbool.
			NewErrorWithType(
				false,
				errtype.SymbolicLink,
				err)

		return &resultWithError
	}

	return errbool.
		EmptyErrorResultPtr(true)
}
