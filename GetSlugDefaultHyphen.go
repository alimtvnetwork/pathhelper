package pathhelper

import "gitlab.com/auk-go/core/constants"

func GetSlugDefaultHyphen(path string) string {
	return GetSlug(
		true,
		constants.HyphenRune,
		constants.HyphenRune,
		path)
}
