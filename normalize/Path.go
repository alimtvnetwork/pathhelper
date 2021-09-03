package normalize

import "gitlab.com/evatix-go/core/constants"

// Path By default apply long path fix and regular normalize using os.PathSeparator
func Path(givenPath string) string {
	return pathUsingSeparator(
		true,
		false,
		constants.PathSeparator,
		givenPath,
	)
}
