package pathhelper

func IsDirectoryPath(path string) bool {
	fileInfoWrapper := GetFileInfoWrapper(path)

	return fileInfoWrapper.IsPathExists() && fileInfoWrapper.IsDirectory
}
