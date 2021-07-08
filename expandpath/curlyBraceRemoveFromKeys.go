package expandpath

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreutils/stringutil"
)

func curlyBraceRemoveFromKeys(slice []string) []EnvKeyInfo {
	length := len(slice)
	newSlice := make(
		[]EnvKeyInfo,
		length)

	for i, key := range slice {
		name := key
		hasTwoChars := len(name) >= 2
		hasCurlyBrace := hasTwoChars &&
			stringutil.IsStartsWith(
				name[constants.One:],
				constants.CurlyStart,
				false)

		if hasCurlyBrace {
			replacer := name[:constants.Two]
			name = strings.Replace(
				name,
				replacer,
				constants.EmptyString,
				constants.One)
			name = strings.Replace(
				name,
				constants.CurlyEnd,
				constants.EmptyString,
				constants.One)
		} else if hasTwoChars {
			name = name[1:]
		}

		newSlice[i] = EnvKeyInfo{
			GivenAs:        key,
			SimplifiedName: name,
		}
	}

	return newSlice
}
