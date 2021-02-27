package normalize

func PathUsingSeparatorUsingSingleIf(
	isNormalizeLongPathForce bool,
	pathSeparator,
	givenPath string,
) string {
	return PathUsingSeparatorIf(
		isNormalizeLongPathForce,
		isNormalizeLongPathForce,
		isNormalizeLongPathForce,
		pathSeparator,
		givenPath)
}
