package pathhelper

func GetAbsolutePathsOfPtr(
	basePath string,
	isLongPathFix, isNormalize bool,
	relativePaths *[]string,
) *[]string {
	return GetAsyncProcessed(
		relativePaths,
		func(
			index int,
			relativePath string,
		) (result string) {
			return GetAbsolutePath(
				basePath,
				relativePath,
				isLongPathFix,
				isNormalize)
		})
}
