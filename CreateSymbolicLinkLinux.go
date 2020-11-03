package pathhelper

import (
	"fmt"
	"gitlab.com/evatix-go/pathhelper/constants"
	"os/exec"
)

// Creates symbolicLink of the source at the provided destination path for linux system. If destination doesn't exist it will panic.
// sourcePath example: "/home/a/test.txt"; destinationPath example: "/home/a/go/test.txt"
// destination need to have read and write permission for the user.
func CreateSymbolicLinkLinux(sourcePath, destinationPath string) {
	if !IsUnix() {
		return
	}

	_, err := exec.Command(
		constants.SymbolicLinkCreationCommandName,
		constants.SymbolicLinkCreationArgument,
		sourcePath,
		destinationPath).Output()

	if err != nil {
		panicMessage := fmt.Sprintf("Error found in CreateSymbolicLinkLinux function: %s", err)
		panic(panicMessage)
	}
}
