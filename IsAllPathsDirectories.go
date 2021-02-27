package pathhelper

func IsAllPathsDirectories(fullPaths *[]string) bool {
	if fullPaths == nil {
		return false
	}

	convertedFileInfos := GetEachPathConvertedToFileInfoWrapper(
		fullPaths)

	for _, wrapper := range *convertedFileInfos {
		if wrapper != nil && !wrapper.IsDirectory {
			return false
		}
	}

	return true
}
