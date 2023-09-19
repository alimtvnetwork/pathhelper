package createdir

import (
	"gitlab.com/auk-go/errorwrapper"
)

func AllDefaultMode(location string) *errorwrapper.Wrapper {
	return All(location, DefaultDirectoryFileMode)
}
