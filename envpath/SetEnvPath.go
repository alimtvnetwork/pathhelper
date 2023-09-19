package envpath

import (
	"os"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
)

func SetEnvPath(compiledPath string) *errorwrapper.Wrapper {
	err := os.Setenv(constants.Path, compiledPath)

	if err != nil {
		return errnew.Messages.Many(
			errtype.EditFailed,
			"Failed to Add or Update environment paths.",
			compiledPath)
	}

	return nil
}
