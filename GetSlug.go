package pathhelper

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/pathhelper/ispath"
)

// GetSlug from given path,
//
// @fixingSlugRune if char is not alphabet (a-z) o number
func GetSlug(
	isDotValid bool,
	fixingSlugRune,
	spaceFixingSlug rune,
	path string,
) string {
	if ispath.Empty(path) {
		return path
	}

	runes := []rune(path)
	for i, eachChar := range runes {
		if constants.AsciiSpace[eachChar] == constants.One {
			// space
			runes[i] = spaceFixingSlug

			continue
		} else if isAlphabetChar(eachChar) || isNumber(eachChar) {
			runes[i] = eachChar

			continue
		} else if isDotValid && eachChar == constants.DotRune {
			runes[i] = eachChar

			continue
		}

		runes[i] = fixingSlugRune
	}

	finalSlug := RuneRepeatFixToSingle(
		fixingSlugRune,
		string(runes))

	return RuneRepeatFixToSingle(
		spaceFixingSlug,
		finalSlug)
}
