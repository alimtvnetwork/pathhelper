package pathhelper

import "gitlab.com/evatix-go/pathhelper/pathhelpercore"

func getFileInfos(path string) []*pathhelpercore.FileInfoWrapper {
	var fileInfos []*pathhelpercore.FileInfoWrapper
	paths := GetFilesPaths(path) 

	for _, eachPath := range paths {
		fileInfos = append(fileInfos, pathhelpercore.NewFileWrapperInfo(*eachPath))
	}

	return fileInfos
}
