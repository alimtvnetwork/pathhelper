package createdir

import "gitlab.com/auk-go/errorwrapper"

func AllUptoParentManyDefaultMode(locations ...string) *errorwrapper.Wrapper {
	return AllUptoParentMany(DefaultDirectoryFileMode, locations...)
}
