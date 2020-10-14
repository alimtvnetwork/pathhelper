package pathhelper

// Replace both double slashes to single slash (// -> /, \\ -> \)
func RemovingDoubleSeparator(path string) string {

	doubleBackSlashToSingle := RemovingDoubleBackSlash(path)
	doubleForwardSlashToSingle := RemovingDoubleSlash(doubleBackSlashToSingle)

	return doubleForwardSlashToSingle
}
