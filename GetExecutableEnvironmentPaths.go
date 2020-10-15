package pathhelper

import (
	"os"
	"strings"

	"gitlab.com/evatix-go/pathhelper/constants"
)

type ExecutableEnvironmentPath struct {
	Variable     string
	Expanded     string
	files        []*string
	filesAsInfos []*os.FileInfo
}

// // returns all files fileInfo on that env directory
// func (eep *ExecutableEnvironmentPath) GetFiles() []*string {
// 	if isEmptyArrayPtr(eep.filesAsInfos) {
// 		filePaths := eep.GetFiles()
// 		// for each filePaths -> filePath
// 		// generate, hint use pathhelpercore.NewFileInfo(filePath)
// 		// we must use try, incase of panic we should be able to recover here.
// 		eep.filesAsInfos = generated file infos
// 	}

// 	return eep.filesAsInfos
// }

// // returns all directories path on that env directory
// func (eep *ExecutableEnvironmentPath) GetDirectories() []*string {

// }

// // returns all files paths on which contains the given string
// func (eep *ExecutableEnvironmentPath) GetFilesContains(contains string) []*string {

// }

// type ExecutableEnvironmentPaths struct {
// 	pathsMap *map[string]ExecutableEnvironmentPath
// }

func GetExecutableEnvironmentPaths() []string {
	var paths []string

	pathString := os.Getenv(constants.Path)
	paths = strings.Split(pathString, constants.SemiColon)

	return paths
}
