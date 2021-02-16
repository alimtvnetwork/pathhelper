package normalize

func PathUsingSeparatorIf(isLongPathFix, isNormalize bool, pathSeparator, givenPath string) string {
	isApplyLongPathFix := !isNormalize &&
		isLongPathFix

	if isApplyLongPathFix {
		return GetLongPathFixed(givenPath)
	}

	return PathUsingSeparator(
		pathSeparator,
		givenPath,
		isLongPathFix)
}
