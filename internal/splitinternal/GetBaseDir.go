package splitinternal

// No slash at the end
// reference example : https://play.golang.org/p/EMbLKv5Jyqe
func GetBaseDir(currentPath string) (baseDir string) {
	i := LastSlash(
		currentPath)

	return currentPath[:i]
}
