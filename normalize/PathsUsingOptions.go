package normalize

import "gitlab.com/auk-go/core/osconsts"

func PathsUsingOptions(
	options *Options,
	locations []string,
) []string {
	if len(locations) == 0 {
		return []string{}
	}

	newItems := make([]string, len(locations))

	for i, location := range locations {
		newItems[i] = PathUsingSeparatorIf(
			options.IsForceLongPathFix,
			options.IsLongPathFix,
			options.IsNormalize,
			osconsts.PathSeparator,
			location)
	}

	return newItems
}
