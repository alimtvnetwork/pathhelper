package pathwrapper

import (
	"io/ioutil"
	"os"
	"path"
	"strings"

	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errdata/errbool"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/pathhelper/internal/splitinternal"
	"gitlab.com/evatix-go/pathhelper/pathext"
)

type Wrapper string

func (it *Wrapper) Value() string {
	return string(*it)
}

func (it *Wrapper) String() string {
	return string(*it)
}

func (it *Wrapper) GetFileInfo() (
	os.FileInfo,
	*errorwrapper.Wrapper,
) {
	fileInfo, err := os.Stat(it.String())

	return fileInfo, errnew.ErrPtr(err)
}

func (it *Wrapper) GetDirectory() *errstr.Result {
	info, e := it.GetFileInfo()

	if e.HasError() {
		return errstr.ErrorWrapper(e)
	}

	currentPath := it.String()

	if info.IsDir() {
		return errstr.New(currentPath, e)
	}

	// file
	return errstr.New(path.Dir(currentPath), e)
}

func (it *Wrapper) DirStatus() *errbool.Result {
	info, e := it.GetFileInfo()

	if e.HasError() {
		return errbool.ErrorWrapper(e)
	}

	return errbool.New(
		info.IsDir(),
		e)
}

func (it *Wrapper) GetBaseDir() string {
	return splitinternal.GetBaseDir(it.Value())
}

func (it *Wrapper) Parent() *Wrapper {
	baseDir := splitinternal.GetBaseDir(it.Value())
	parent := Wrapper(baseDir)

	return &parent
}

func (it *Wrapper) IsDir() bool {
	resultPtr := it.DirStatus()

	if resultPtr.ErrorWrapper.HasError() {
		return false
	}

	return resultPtr.Value
}

func (it *Wrapper) IsFile() bool {
	resultPtr := it.DirStatus()

	if resultPtr.ErrorWrapper.HasError() {
		return false
	}

	// not dir means file.
	return !resultPtr.Value
}

func (it *Wrapper) IsExist() bool {
	resultPtr := it.DirStatus()

	if resultPtr.ErrorWrapper.HasError() {
		return false
	}

	// not dir means file.
	return true
}

func (it *Wrapper) BothExtensions() (dotExt, ext string) {
	return splitinternal.GetBothExtension(it.String())
}

// DotExtension .mp4 reference: https://stackoverflow.com/a/64122557
func (it *Wrapper) DotExtension() string {
	dotExt, _ := splitinternal.GetBothExtension(it.String())

	return dotExt
}

// Extension mp4 reference: https://stackoverflow.com/a/64122557
func (it *Wrapper) Extension() string {
	_, ext := splitinternal.GetBothExtension(it.String())

	return ext
}

func (it *Wrapper) ExtensionWrapper() *pathext.Wrapper {
	return pathext.NewPtr(it.String())
}

// GetDirectoriesDefault Get all directory on that root path only, no nested or recursive visit.
func (it *Wrapper) GetDirectoriesDefault() *errstr.Results {
	return it.GetDirectories(osconsts.PathSeparator)
}

func (it *Wrapper) GetAllPathsDefault() *errstr.Results {
	return it.GetAllPaths(osconsts.PathSeparator)
}

// GetAllPaths Get all paths on that root path only, no nested or recursive visit.
func (it *Wrapper) GetAllPaths(separator string) *errstr.Results {
	fileInfos, errW := it.GetDirFileInfos()
	if errW.HasError() {
		return errstr.EmptyResultsWithError(
			errW)
	}

	rootPath := it.GetDirectory().Value

	// file
	results := make([]string, 0, len(*fileInfos))
	for _, info := range *fileInfos {
		currentPath := rootPath + separator + info.Name()
		results = append(results, currentPath)
	}

	return errstr.EmptyErrorResults(
		results...)
}

// GetDirectories Get all directory on that root path only, no nested or recursive visit.
func (it *Wrapper) GetDirectories(separator string) *errstr.Results {
	fileInfos, errW := it.GetDirFileInfos()
	if errW.HasError() {
		return errstr.EmptyResultsWithError(
			errW)
	}

	rootPath := it.GetDirectory().Value

	// file
	results := make([]string, 0, len(*fileInfos))
	for _, info := range *fileInfos {
		if info.IsDir() {
			currentPath := rootPath + separator + info.Name()
			results = append(results, currentPath)
		}
	}

	return errstr.EmptyErrorResults(
		results...)
}

func (it *Wrapper) GetAFilePathAsString(
	separator string,
	nesting ...string,
) string {
	rootPath := it.GetDirectory().Value
	nestingCombined := strings.Join(nesting, separator)

	currentPath := rootPath +
		separator +
		nestingCombined

	return currentPath
}

// GetAFilePath Get a file path combining file path.
func (it *Wrapper) GetAFilePath(
	separator string,
	nesting ...string,
) Wrapper {
	return Wrapper(it.GetAFilePathAsString(separator, nesting...))
}

func (it *Wrapper) GetNestedDirectories(
	separator string,
	nesting ...string,
) *errstr.Results {
	fileInfos, errW := it.GetDirFileInfos()
	if errW.HasError() {
		return &errstr.Results{
			Values:       nil,
			ErrorWrapper: errW,
		}
	}

	rootPath := it.GetDirectory().Value
	nestingCombined := strings.Join(nesting, separator)

	// file
	results := make([]string, 0, len(*fileInfos))
	for _, info := range *fileInfos {
		if info.IsDir() {
			currentPath := rootPath +
				separator +
				nestingCombined +
				separator +
				info.Name()
			results = append(results, currentPath)
		}
	}

	return errstr.EmptyErrorResults(
		results...)
}

// GetFilesDefault Get all files on that root path only, no nested or recursive visit.
func (it *Wrapper) GetFilesDefault() *errstr.Results {
	return it.GetFiles(osconsts.PathSeparator)
}

// GetFiles Get all files on that root path only, no nested or recursive visit.
func (it *Wrapper) GetFiles(separator string) *errstr.Results {
	fileInfos, errW := it.GetDirFileInfos()
	if errW.HasError() {
		return errstr.EmptyResultsWithError(
			errW)
	}

	rootPath := it.
		GetDirectory().
		Value

	// file
	results := make([]string, 0, len(*fileInfos))
	for _, info := range *fileInfos {
		if !info.IsDir() {
			currentPath := rootPath +
				separator +
				info.Name()
			results = append(results, currentPath)
		}
	}

	return errstr.EmptyErrorResults(
		results...)
}

func (it *Wrapper) GetDirFileInfos() (
	*[]os.FileInfo, *errorwrapper.Wrapper,
) {
	directoryResult := it.GetDirectory()

	if directoryResult.ErrorWrapper.HasError() {
		return nil, directoryResult.
			ErrorWrapper
	}

	fileInfos, err := ioutil.ReadDir(directoryResult.Value)

	if err != nil {
		return nil, errnew.ErrPtr(err)
	}

	return &fileInfos, errnew.EmptyPtr
}
