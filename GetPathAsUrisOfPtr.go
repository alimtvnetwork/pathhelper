package pathhelper

func GetPathAsUrisOfPtr(
	isNormalizePath bool,
	paths *[]string,
) *[]string {
	return GetAsyncProcessed(
		paths,
		func(
			index int,
			currentPath string,
		) (result string) {
			return GetPathAsUri(
				currentPath,
				isNormalizePath)
		})
}
