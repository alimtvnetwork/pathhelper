package pathhelper

func NormalizePathUsingSeparatorIf(isLongPathFix, isNormalize bool, pathSeparator, givenPath string) string {
	isApplyLongPathFix := !isNormalize &&
		isLongPathFix

	if isApplyLongPathFix {
		return GetLongPathFixed(givenPath)
	}

	return NormalizePathUsingSeparator(
		pathSeparator,
		givenPath,
		isLongPathFix)
}
