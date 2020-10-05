package pathhelper

func IsPathExist(path string) bool {
	fileInfoWrapper := GetFileInfoWrapper(path)

	return fileInfoWrapper.IsPathExists()
}
