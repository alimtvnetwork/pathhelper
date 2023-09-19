package normalize

import "gitlab.com/auk-go/core/osconsts"

func PathFixWithoutLongPathIf(
	isNormalize bool,
	givenPath string,
) string {
	if !isNormalize {
		return givenPath
	}

	return PathUsingSeparatorIf(
		false,
		false,
		isNormalize,
		osconsts.PathSeparator,
		givenPath)
}
