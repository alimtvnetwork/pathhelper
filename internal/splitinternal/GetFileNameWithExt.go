package splitinternal

// reference example : https://play.golang.org/p/oT6eWNZAeEi
func GetFileNameWithExt(currentPath string) (fileName string) {
	i := LastSlash(
		currentPath)

	return currentPath[i+1:]
}
