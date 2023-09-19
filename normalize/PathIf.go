package normalize

import "gitlab.com/auk-go/core/osconsts"

func PathIf(
	isNormalize bool,
	givenPath string,
) string {
	return PathUsingSeparatorIf(
		false,
		isNormalize,
		isNormalize,
		osconsts.PathSeparator,
		givenPath,
	)
}
