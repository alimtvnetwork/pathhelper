package pathhelper

// todo delete
// Takes in input of a directory under users directory and returns its absolute path.
func GetUserPathOf(directoryName string) string {
	outputPath := GetCombinePathWith(GetUserPath(), directoryName)

	return outputPath
}
