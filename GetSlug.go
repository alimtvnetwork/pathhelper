package pathhelper

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/pathhelper/ispath"
)

// GetSlug from given path,
//
// @fixingSlugRune if char is not alphabet (a-z) o number
func GetSlug(fixingSlugRune, spaceFixingSlug rune, path string) string {
	if ispath.Empty(path) {
		return path
	}

	runes := []rune(path)

	for i, eachChar := range runes {
		if constants.AsciiSpace[eachChar] == constants.One {
			runes[i] = spaceFixingSlug

			continue
		} else if isAlphabetChar(eachChar) || isNumber(eachChar) {
			runes[i] = eachChar

			continue
		}

		runes[i] = fixingSlugRune
	}

	slugFixed := string(runes)
	fixingSlug := string(fixingSlugRune)
	fixingSlugRepeat3 := strings.Repeat(
		fixingSlug,
		constants.Capacity3)
	fixingSlugRepeat2 := strings.Repeat(
		fixingSlug,
		constants.Capacity2)

	slugFixed = strings.ReplaceAll(
		slugFixed,
		fixingSlugRepeat3,
		fixingSlug)

	slugFixed = strings.ReplaceAll(
		slugFixed,
		fixingSlugRepeat2,
		fixingSlug)

	spaceFixing := string(spaceFixingSlug)
	spaceFixingRepeat3 := strings.Repeat(
		spaceFixing,
		constants.Capacity3)
	spaceFixingRepeat2 := strings.Repeat(
		spaceFixing,
		constants.Capacity2)

	slugFixed = strings.ReplaceAll(
		slugFixed,
		spaceFixingRepeat3,
		spaceFixing)

	return strings.ReplaceAll(
		slugFixed,
		spaceFixingRepeat2,
		spaceFixing)
}
