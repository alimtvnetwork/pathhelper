package normalize

import "gitlab.com/evatix-go/core/osconsts"

func PathsUsingSingleIf(
	isNormalizeLongPathForce bool,
	locations []string,
) []string {
	if len(locations) == 0 {
		return []string{}
	}

	if !isNormalizeLongPathForce {
		return locations
	}

	newItems := make([]string, len(locations))

	for i, location := range locations {
		newItems[i] = PathUsingSeparatorIf(
			isNormalizeLongPathForce,
			isNormalizeLongPathForce,
			isNormalizeLongPathForce,
			osconsts.PathSeparator,
			location)
	}

	return newItems
}
