package strindexinternal

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

func OfPtr(
	lines *[]string,
	findingString *string,
	startsAtIndex int,
	isCaseSensitive bool,
) int {
	if lines == nil || *lines == nil {
		return constants.InvalidNotFoundCase
	}

	length := len(*lines)

	if !isCaseSensitive {
		// insensitive
		for i := startsAtIndex; i < length; i++ {
			if strings.EqualFold((*lines)[i], *findingString) {
				return i
			}
		}

		return constants.InvalidValue
	}

	for i := startsAtIndex; i < length; i++ {
		if (*lines)[i] == *findingString {
			return i
		}
	}

	return constants.InvalidNotFoundCase
}

func OfPtrSimple(
	lines *[]string,
	findingString string,
	startsAtIndex int,
	isCaseSensitive bool,
) int {
	if lines == nil || *lines == nil {
		return constants.InvalidNotFoundCase
	}

	length := len(*lines)

	if length == 0 {
		return constants.InvalidNotFoundCase
	}

	if !isCaseSensitive {
		// insensitive
		for i := startsAtIndex; i < length; i++ {
			if strings.EqualFold((*lines)[i], findingString) {
				return i
			}
		}

		return constants.InvalidValue
	}

	for i := startsAtIndex; i < length; i++ {
		if (*lines)[i] == findingString {
			return i
		}
	}

	return constants.InvalidNotFoundCase
}

func Of(
	lines []string,
	findingString string,
	startsAtIndex int,
	isCaseSensitive bool,
) int {
	if lines == nil {
		return constants.InvalidNotFoundCase
	}

	length := len(lines)

	if length == 0 {
		return constants.InvalidNotFoundCase
	}

	if !isCaseSensitive {
		// insensitive
		for i := startsAtIndex; i < length; i++ {
			if strings.EqualFold(lines[i], findingString) {
				return i
			}
		}

		return constants.InvalidValue
	}

	for i := startsAtIndex; i < length; i++ {
		if lines[i] == findingString {
			return i
		}
	}

	return constants.InvalidNotFoundCase
}
