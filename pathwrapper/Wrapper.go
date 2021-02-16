package pathwrapper

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errdata/errbool"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errnew"

	"gitlab.com/evatix-go/pathhelper/pathext"
)

type Wrapper string

func (receiver Wrapper) Value() string {
	return string(receiver)
}

func (receiver Wrapper) String() string {
	return string(receiver)
}

func (receiver Wrapper) GetFileInfo() (
	os.FileInfo,
	*errorwrapper.Wrapper,
) {
	fileInfo, err := os.Stat(receiver.String())

	return fileInfo, errnew.ErrPtr(err)
}

func (receiver Wrapper) GetDirectory() *errstr.Result {
	info, e := receiver.GetFileInfo()

	if e.HasError() {
		return errstr.ErrorWrapperPtr(e)
	}

	currentPath := receiver.String()

	if info.IsDir() {
		return errstr.NewUsingWrapperPtr(currentPath, e)
	}

	// file
	return errstr.NewUsingWrapperPtr(path.Dir(currentPath), e)
}

func (receiver Wrapper) DirStatus() *errbool.Result {
	info, e := receiver.GetFileInfo()

	if e.HasError() {
		return errbool.ErrorWrapperPtr(e)
	}

	return errbool.NewUsingWrapperPtr(
		info.IsDir(),
		e)
}

func (receiver Wrapper) IsDir() bool {
	resultPtr := receiver.DirStatus()

	if resultPtr.ErrorWrapper.HasError() {
		return false
	}

	return resultPtr.Value
}

func (receiver Wrapper) IsFile() bool {
	resultPtr := receiver.DirStatus()

	if resultPtr.ErrorWrapper.HasError() {
		return false
	}

	// not dir means file.
	return !resultPtr.Value
}

func (receiver Wrapper) IsExist() bool {
	resultPtr := receiver.DirStatus()

	if resultPtr.ErrorWrapper.HasError() {
		return false
	}

	// not dir means file.
	return true
}

// .mp4 reference: https://stackoverflow.com/a/64122557
func (receiver Wrapper) DotExtension() string {
	return filepath.Ext(receiver.String())
}

// mp4 reference: https://stackoverflow.com/a/64122557
func (receiver Wrapper) Extension() string {
	ext := filepath.Ext(receiver.String())

	if len(ext) > 0 && ext[0] == constants.Dot[0] {
		return ext[1:]
	}

	return ext
}

func (receiver Wrapper) ExtensionWrapper() *pathext.Wrapper {
	return pathext.NewPtr(receiver.String())
}

// Get all directory on that root path only, no nested or recursive visit.
func (receiver Wrapper) GetDirectories(separator string) *errstr.Results {
	fileInfos, errW := receiver.getFileInfos()
	if errW.HasError() {
		return &errstr.Results{
			Values:       nil,
			ErrorWrapper: errW,
		}
	}

	rootPath := receiver.GetDirectory().Value

	// file
	results := make([]string, 0, len(*fileInfos))
	for _, info := range *fileInfos {
		if info.IsDir() {
			currentPath := rootPath + separator + info.Name()
			results = append(results, currentPath)
		}
	}

	return &errstr.Results{
		Values:       &results,
		ErrorWrapper: errnew.EmptyPtr,
	}
}

// Get a file path combining file path.
func (receiver Wrapper) GetAFilePath(separator string, nesting ...string) Wrapper {
	rootPath := receiver.GetDirectory().Value
	nestingCombined := strings.Join(nesting, separator)

	currentPath := rootPath +
		separator +
		nestingCombined

	return Wrapper(currentPath)
}

func (receiver Wrapper) GetNestedDirectories(
	separator string,
	nesting ...string,
) *errstr.Results {
	fileInfos, errW := receiver.getFileInfos()
	if errW.HasError() {
		return &errstr.Results{
			Values:       nil,
			ErrorWrapper: errW,
		}
	}

	rootPath := receiver.GetDirectory().Value
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

	return &errstr.Results{
		Values:       &results,
		ErrorWrapper: errnew.EmptyPtr,
	}
}

// Get all files on that root path only, no nested or recursive visit.
func (receiver Wrapper) GetFiles(separator string) *errstr.Results {
	fileInfos, errW := receiver.getFileInfos()
	if errW.HasError() {
		return &errstr.Results{
			Values:       nil,
			ErrorWrapper: errW,
		}
	}

	rootPath := receiver.GetDirectory().Value

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

	return &errstr.Results{
		Values:       &results,
		ErrorWrapper: errnew.EmptyPtr,
	}
}

func (receiver Wrapper) getFileInfos() (
	*[]os.FileInfo, *errorwrapper.Wrapper,
) {
	directoryResult := receiver.GetDirectory()

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
