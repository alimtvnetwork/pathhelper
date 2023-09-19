package symlink

import (
	"os"

	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
)

func Create(path, linkName string) *errorwrapper.Wrapper {
	err := os.Symlink(path, linkName)

	if err != nil {
		return errnew.Ref.TwoWithError(
			errtype.SymbolicLink,
			err,
			"Source Symbolic Link",
			path,
			"Symbolic Link Place",
			linkName)
	}

	return nil
}
