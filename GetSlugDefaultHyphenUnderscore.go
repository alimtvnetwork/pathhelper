package pathhelper

import "gitlab.com/auk-go/core/constants"

func GetSlugDefaultHyphenUnderscore(path string) string {
	return GetSlug(
		true,
		constants.HyphenRune,
		constants.UnderscoreRune,
		path)
}
