package pathjoin

import "gitlab.com/auk-go/core/osconsts"

// JoinNormalizedThreeIf normalized applied auto
func JoinNormalizedThreeIf(
	isNormalize bool,
	first, second, third string,
) string {
	if isNormalize {
		return JoinNormalizedThree(first, second, third)
	}

	finalPath := first +
		osconsts.PathSeparator +
		second +
		osconsts.PathSeparator +
		third

	return finalPath
}
