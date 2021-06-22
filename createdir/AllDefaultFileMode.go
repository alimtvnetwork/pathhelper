package createdir

import (
	"gitlab.com/evatix-go/errorwrapper"
)

func AllDefaultFileMode(location string) *errorwrapper.Wrapper {
	return All(location, DefaultFileMode)
}
