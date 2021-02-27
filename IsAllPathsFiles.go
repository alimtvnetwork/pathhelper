package pathhelper

func IsAllPathsFiles(fullPaths *[]string) bool {
	if fullPaths == nil {
		return false
	}

	convertedFileInfos := GetEachPathConvertedToFileInfoWrapper(
		fullPaths)

	for _, wrapper := range *convertedFileInfos {
		if wrapper != nil && !wrapper.IsFile {
			return false
		}
	}

	return true
}
