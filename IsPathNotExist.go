package pathhelper

func IsPathNotExist(path string) bool {
	fileInfoWrapper := GetFileInfoWrapper(path)

	return fileInfoWrapper.HasError()
}
