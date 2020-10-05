package pathhelper

func GetAbsoluteFromExecutableDirectoryPath(relativePath string) string {
	basePath := GetExecutableDirectory()

	return GetAbsolutePath(basePath, relativePath)
}
