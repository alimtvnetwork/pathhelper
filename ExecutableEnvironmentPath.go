package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/constants"
	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
	"io/ioutil"
	"strings"
	"sync"
)

type ExecutableEnvironmentPath struct {
	Variable     string
	Expanded     string
	files        []*string
	filesAsInfos []*pathhelpercore.FileInfoWrapper
}

var mutex = &sync.Mutex{}

// returns all files paths on that env directory once, caches it and returns that in later function calls
func (eep *ExecutableEnvironmentPath) GetLazyFilePaths() []*string {
	// checking if fileInfos already generated
	if !pathhelpercore.IsEmptyArrayPtr(eep.files) {
		return eep.files
	}

	// generate if not generated already
	eep.files = GenerateFiles(eep.Expanded)

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
	mutex.Lock()
	eep.filesAsInfos = generateFileInfos(eep.Expanded) // todo add recover
	mutex.Unlock()

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

// returns filepaths as []*string. non-lazy execution.
func GenerateFiles(path string) []*string {
	var fileNames []*string

	files, err := ioutil.ReadDir(path)

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fileName := file.Name()
		fileNames = append(fileNames, &fileName)
	}

	return fileNames
}

func generateFileInfos(path string) []*pathhelpercore.FileInfoWrapper {
	var fileInfos []*pathhelpercore.FileInfoWrapper
	paths := GenerateFiles(path) // double mutex???

	for _ , eachPath := range paths {
		fileInfos = append(fileInfos, pathhelpercore.NewFileWrapperInfo(*eachPath))
	}
	
	return fileInfos
}