package pathhelper

func GetSlugsOfPtr(
	separatorOfChoice string,
	paths *[]string,
) *[]string {
	return GetAsyncProcessed(
		paths,
		func(
			index int,
			currentPath string,
		) (result string) {
			return GetSlug(
				currentPath,
				separatorOfChoice)
		})
}
