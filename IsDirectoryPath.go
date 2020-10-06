package pathhelper

// Slow performance used os.FileInfo, TODO: improve performance in future
func IsDirectoryPath(path string) bool {
	fileInfoWrapper := GetFileInfoWrapper(path)

	return fileInfoWrapper.IsPathExists() && fileInfoWrapper.IsDirectory
}
