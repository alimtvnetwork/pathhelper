package normalize

func PathUsingSeparatorIf(
	isForceLongPath,
	isLongPathFix,
	isNormalize bool,
	pathSeparator,
	givenPath string,
) string {
	isApplyLongPathFix := !isNormalize &&
		(isLongPathFix || isForceLongPath)

	if isNormalize {
		givenPath = TrimPrefixUncPath(
			givenPath)
	}

	if isApplyLongPathFix {
		return getLongPathFixedUsingSeparator(
			pathSeparator,
			givenPath,
			isForceLongPath)
	}

	if !isNormalize {
		return givenPath
	}

	return pathUsingSeparator(
		isLongPathFix,
		isForceLongPath,
		pathSeparator,
		givenPath,
	)
}
