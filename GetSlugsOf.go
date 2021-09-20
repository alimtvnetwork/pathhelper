package pathhelper

func GetSlugsOf(
	slugFixer, spaceSlugFixer rune,
	paths []string,
) []string {
	return GetAsyncProcessed(
		paths,
		func(
			index int,
			currentPath string,
		) (result string) {
			return GetSlug(
				slugFixer,
				spaceSlugFixer,
				currentPath,
			)
		})
}
