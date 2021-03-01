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

	if isApplyLongPathFix {
		return GetLongPathFixedPtr(
			pathSeparator,
			givenPath,
			isForceLongPath)
	}

	if !isNormalize {
		return givenPath
	}

	return PathUsingSeparator(
		pathSeparator,
		givenPath,
		isLongPathFix,
		isForceLongPath)
}
