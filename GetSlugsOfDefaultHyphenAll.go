package pathhelper

import "gitlab.com/auk-go/core/constants"

func GetSlugsOfDefaultHyphenAll(
	paths ...string,
) []string {
	return GetSlugsOf(
		true,
		constants.HyphenRune,
		constants.HyphenRune,
		paths)
}
