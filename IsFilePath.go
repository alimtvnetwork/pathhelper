package pathhelper

func IsFilePath(path string) bool {
	fileInfoWrapper := GetFileInfoWrapper(path)

	return fileInfoWrapper.IsPathExists() && fileInfoWrapper.IsFile
}
