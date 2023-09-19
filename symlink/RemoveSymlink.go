package symlink

import (
	"fmt"
	"os"

	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
	"gitlab.com/auk-go/pathhelper/internal/messages"
)

func RemoveSymlink(path string) *errorwrapper.Wrapper {
	result := IsSymLink(path)

	if result.HasError() {
		return result.ErrorWrapper
	}

	if !result.Value {
		invalidSymLink := fmt.Sprintf(messages.InvalidSymlinkMessageFormat, path)

		return errnew.Messages.Many(errtype.SymbolicLink, invalidSymLink)
	}

	err := os.Remove(path)

	if err != nil {
		invalidSymLink := fmt.Sprintf(messages.CannotRemoveSymLink, path)

		return errnew.Messages.Many(errtype.SymbolicLink, invalidSymLink)
	}

	return nil
}
