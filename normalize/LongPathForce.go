package normalize

import "gitlab.com/evatix-go/core/constants"

// LongPathForce By default apply
//
// long path fix,
// force long path fix on windows
// and regular normalize using os.PathSeparator
func LongPathForce(givenPath string) string {
	return PathUsingSeparator(
		constants.PathSeparator,
		givenPath,
		true,
		true)
}
