package fileinfo

import (
	"gitlab.com/evatix-go/core/coreutils/stringutil"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

type FileNamesCollection struct {
	RootPath string
	// TODO fix this to file paths
	names          *[]string
	Error          *errorwrapper.Wrapper
	parentWrappers *Wrappers
}

func NewFileNames(rootPath string, capacity int) *FileNamesCollection {
	paths := make([]string, 0, capacity)

	return &FileNamesCollection{
		RootPath:       rootPath,
		names:          &paths,
		Error:          errnew.EmptyPtr,
		parentWrappers: nil,
	}
}

func NewFileNamesUsingWrappers(wrappers *Wrappers) *FileNamesCollection {
	if wrappers == nil {
		return &FileNamesCollection{
			RootPath:       "",
			names:          nil,
			Error:          errnew.EmptyPtr,
			parentWrappers: wrappers,
		}
	}

	if wrappers.IsEmpty() {
		return &FileNamesCollection{
			RootPath:       wrappers.RootPath,
			names:          nil,
			Error:          wrappers.ErrorWrapper,
			parentWrappers: wrappers,
		}
	}

	names := make([]string, wrappers.Length())
	for i, wrapper := range *wrappers.Items {
		if wrapper.FileInfo == nil {
			continue
		}

		names[i] = wrapper.FileInfo.Name()
	}

	return &FileNamesCollection{
		RootPath:       wrappers.RootPath,
		names:          &names,
		Error:          wrappers.ErrorWrapper,
		parentWrappers: wrappers,
	}
}

func NewFileNamesUsing(
	directoryPath, separator string,
	isNormalize bool,
) *FileNamesCollection {
	wrappers := NewWrappersPtr(
		directoryPath,
		separator,
		isNormalize)

	return NewFileNamesUsingWrappers(wrappers)
}

func (filesNamesCollection *FileNamesCollection) IsEmpty() bool {
	return filesNamesCollection.names == nil ||
		filesNamesCollection.Error.HasError() ||
		len(*filesNamesCollection.names) == 0
}

func (filesNamesCollection *FileNamesCollection) IsContains(
	fileName string,
	isCaseSensitive bool,
) bool {
	return stringutil.IsContainsPtrSimple(
		filesNamesCollection.names,
		fileName,
		0,
		isCaseSensitive)
}

func (filesNamesCollection *FileNamesCollection) Length() int {
	if filesNamesCollection.names == nil || *filesNamesCollection.names == nil {
		return 0
	}

	return len(*filesNamesCollection.names)
}

// Root level files paths, no nested paths.
func (filesNamesCollection *FileNamesCollection) GetFilePaths(
	separator string,
) *[]string {
	filePaths := make([]string, 0, filesNamesCollection.Length())

	for _, name := range *filesNamesCollection.names {
		newPath := filesNamesCollection.RootPath + separator + name
		filePaths = append(filePaths, newPath)
	}

	return &filePaths
}

func (filesNamesCollection *FileNamesCollection) IsParentWrappersEmpty() bool {
	return filesNamesCollection.parentWrappers == nil ||
		filesNamesCollection.parentWrappers.IsEmpty()
}

func (filesNamesCollection *FileNamesCollection) HasParentWrappers() bool {
	return filesNamesCollection.parentWrappers != nil
}

func (filesNamesCollection *FileNamesCollection) ParentWrappers() *Wrappers {
	return filesNamesCollection.parentWrappers
}

func (filesNamesCollection *FileNamesCollection) Add(fileName string) *FileNamesCollection {
	*filesNamesCollection.names = append(*filesNamesCollection.names, fileName)

	return filesNamesCollection
}

func (filesNamesCollection *FileNamesCollection) AddWrapper(wrapper Wrapper) *FileNamesCollection {
	*filesNamesCollection.names = append(*filesNamesCollection.names, wrapper.FileInfo.Name())

	return filesNamesCollection
}

func (filesNamesCollection *FileNamesCollection) OnlyNamesCollection() *[]string {
	return filesNamesCollection.names
}
