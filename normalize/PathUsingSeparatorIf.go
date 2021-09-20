package normalize

func PathUsingSeparatorIf(
	isForceLongPath,
	isLongPathFix,
	isNormalize bool,
	pathSeparator,
	givenPath string,
) string {
	isApplyLongPathFixOnly := !isNormalize &&
		(isLongPathFix || isForceLongPath)

	if isNormalize || isLongPathFix || isForceLongPath {
		givenPath = TrimPrefixUncPathIf(
			true,
			givenPath)
	}

	if isApplyLongPathFixOnly {
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
