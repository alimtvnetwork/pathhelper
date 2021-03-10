package envpath

import (
	"strings"
	"sync"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/fileinfo"
)

type ExecutableEnvironmentPath struct {
	Variable         string
	Expanded         string
	pathsCollections *fileinfo.FileNamesCollection
	fileWrappers     *fileinfo.Wrappers
	sync.Mutex
}

// returns all pathsCollection paths on that env directory once,
// caches it and returns that in later function calls
func (eep *ExecutableEnvironmentPath) GetCachedFileNamesCollection() *fileinfo.FileNamesCollection {
	// checking if fileInfos already generated
	if eep.pathsCollections != nil {
		return eep.pathsCollections
	}

	eep.pathsCollections = eep.GetFileNamesCollection()

	return eep.pathsCollections
}

func (eep *ExecutableEnvironmentPath) Length() int {
	return eep.GetCachedFileNamesCollection().Length()
}

func (eep *ExecutableEnvironmentPath) GetFileNamesCollection() *fileinfo.FileNamesCollection {
	return fileinfo.NewFileNamesUsing(
		eep.Expanded,
		osconsts.PathSeparator,
		true)
}

// returns all directories path on that env directory,
// no nested or resursive paths
func (eep *ExecutableEnvironmentPath) GetDirectories() []*string {
	var directories []*string

	arrayFromPath := strings.Split(eep.Expanded, constants.PathSeparator)

	for _, arrayItem := range arrayFromPath {
		directories = append(directories, &arrayItem)
	}

	return directories
}

// returns all pathsCollection paths on which contains the given string. If no path is found, returns empty array.
func (eep *ExecutableEnvironmentPath) GetFilePathsContains(
	separator,
	contains string,
) *[]string {
	var filePathThatContains = make([]string, 0, eep.Length())
	filesPaths := *eep.GetCachedFileNamesCollection().GetFilePaths(separator)

	for _, eachPath := range filesPaths {
		if strings.Contains(eachPath, contains) {
			filePathThatContains = append(filePathThatContains, eachPath)
		}
	}

	return &filePathThatContains
}
