package splitinternal

// reference example : https://play.golang.org/p/EMbLKv5Jyqe
func GetFileNameWithExt(currentPath string) (fileName string) {
	i := LastSlash(
		currentPath)

	return currentPath[i+1:]
}
