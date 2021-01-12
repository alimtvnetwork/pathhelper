package pathhelper

import (
	"strings"
	"sync"

	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
)

type ExecutableEnvironmentPath struct {
	Variable     string
	Expanded     string
	files        []*string
	filesAsInfos []*pathhelpercore.FileInfoWrapper
}

var executableEnvPathMutex = &sync.Mutex{}

// returns all files paths on that env directory once, caches it and returns that in later function calls
func (eep *ExecutableEnvironmentPath) GetLazyFilePaths() []*string {
	// checking if fileInfos already generated
	if !pathhelpercore.IsEmptyArrayPtr(eep.files) {
		return eep.files
	}

	// generate if not generated already
	eep.files = GetFilesPaths(eep.Expanded)

	return eep.files
}

// returns all files fileInfo on that env directory
func (eep *ExecutableEnvironmentPath) GetFileInfosMap() map[string]*ExecutableEnvironmentPath {
	// checking if fileInfos already generated
	if !pathhelpercore.IsEmptyArrayForFileInfo(eep.filesAsInfos) {
		return map[string]*ExecutableEnvironmentPath{
			eep.Variable: eep,
		}
	}

	// otherwise generate
	executableEnvPathMutex.Lock()
	eep.filesAsInfos = getFileInfos(eep.Expanded) // todo add recover
	executableEnvPathMutex.Unlock()

	return map[string]*ExecutableEnvironmentPath{
		eep.Variable: eep,
	}
}

// returns all directories path on that env directory
func (eep *ExecutableEnvironmentPath) GetDirectories() []*string {
	var directories []*string

	arrayFromPath := strings.Split(eep.Expanded, constants.PathSeparator)

	for _, arrayItem := range arrayFromPath {
		directories = append(directories, &arrayItem)
	}

	return directories
}

// returns all files paths on which contains the given string. If no path is found, returns empty array.
func (eep *ExecutableEnvironmentPath) GetFilesContains(contains string) []*string {
	var filePathThatContains = []*string{}

	// Check if contains
	for _, eachPath := range eep.files {
		if strings.Contains(*eachPath, contains) {
			filePathThatContains = append(filePathThatContains, eachPath)
		}
	}

	return filePathThatContains
}
