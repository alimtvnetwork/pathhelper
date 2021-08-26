package fileinfo

import (
	"encoding/json"

	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/defaulterr"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

type Wrappers struct {
	RootPath            string
	Separator           string
	Items               *[]*Wrapper
	directories         *Wrappers
	files               *Wrappers
	ErrorWrapper        *errorwrapper.Wrapper
	pathsCollection     *PathsCollection
	recursiveDirs       *PathsCollection
	fileNamesCollection *FileNamesCollection
}

func (wrappers *Wrappers) HasAny() bool {
	return !wrappers.IsEmpty()
}

func (wrappers *Wrappers) RootFiles() *Wrappers {
	if wrappers.files != nil {
		return wrappers.files
	}

	if wrappers.IsEmpty() {
		wrappers.files = EmptyWrappers()
	}

	files := make(
		[]*Wrapper,
		0,
		wrappers.Length())

	for _, wrapper := range *wrappers.Items {
		if !wrapper.IsFile {
			continue
		}

		files = append(files, wrapper)
	}

	filesWrapper := &Wrappers{
		Items:               &files,
		directories:         nil,
		files:               nil,
		ErrorWrapper:        errnew.EmptyPtr,
		pathsCollection:     nil,
		fileNamesCollection: nil,
	}

	filesWrapper.files = filesWrapper
	wrappers.files = filesWrapper

	return wrappers.files
}

func (wrappers *Wrappers) RootDirs() *Wrappers {
	if wrappers.directories != nil {
		return wrappers.directories
	}

	if wrappers.IsEmpty() {
		wrappers.directories = EmptyWrappers()
	}

	dirs := make(
		[]*Wrapper,
		0,
		wrappers.Length())

	for _, wrapper := range *wrappers.Items {
		if !wrapper.IsDirectory {
			continue
		}

		dirs = append(dirs, wrapper)
	}

	dirWrappers := &Wrappers{
		Items:               &dirs,
		directories:         nil,
		files:               nil,
		ErrorWrapper:        errnew.EmptyPtr,
		pathsCollection:     nil,
		fileNamesCollection: nil,
	}

	dirWrappers.directories = dirWrappers
	wrappers.directories = dirWrappers

	return wrappers.directories
}

func (wrappers *Wrappers) PathsCollection() *PathsCollection {
	if wrappers.pathsCollection != nil {
		return wrappers.pathsCollection
	}

	wrappers.pathsCollection = NewPathsUsingWrappers(
		wrappers.RootPath,
		wrappers.Separator,
		wrappers)

	return wrappers.pathsCollection
}

func (wrappers *Wrappers) FileNamesCollection() *FileNamesCollection {
	if wrappers.fileNamesCollection != nil {
		return wrappers.fileNamesCollection
	}

	wrappers.fileNamesCollection = NewFileNamesUsingWrappers(wrappers)

	return wrappers.fileNamesCollection
}

func (wrappers *Wrappers) IsEmpty() bool {
	return wrappers.ErrorWrapper.HasError() ||
		wrappers.Items == nil ||
		len(*wrappers.Items) == 0
}

func (wrappers *Wrappers) Length() int {
	if wrappers.Items == nil || *wrappers.Items == nil {
		return 0
	}

	return len(*wrappers.Items)
}

func (wrappers *Wrappers) IsNameContains(
	name string,
	isCaseSensitive bool,
) bool {
	return wrappers.
		FileNamesCollection().
		IsContains(
			name,
			isCaseSensitive)
}

func (wrappers *Wrappers) JsonModel() *Wrappers {
	return wrappers
}

func (wrappers *Wrappers) JsonModelAny() interface{} {
	return wrappers.JsonModel()
}

func (wrappers *Wrappers) AsJsoner() corejson.Jsoner {
	return wrappers
}

func (wrappers *Wrappers) AsJsonParseSelfInjector() corejson.JsonParseSelfInjector {
	return wrappers
}

func (wrappers *Wrappers) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := wrappers.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (wrappers *Wrappers) Json() *corejson.Result {
	if wrappers.IsEmpty() {
		return corejson.EmptyWithoutErrorPtr()
	}

	jsonBytes, err := json.Marshal(wrappers.JsonModel())

	return corejson.NewPtr(jsonBytes, err)
}

func (wrappers *Wrappers) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*Wrappers, error) {
	if jsonResult == nil || jsonResult.IsEmptyJsonBytes() {
		return nil, defaulterr.UnMarshallingFailedDueToNilOrEmpty
	}

	err := json.Unmarshal(
		*jsonResult.Bytes,
		&wrappers)

	if err != nil {
		return nil, err
	}

	return wrappers, nil
}

// ParseInjectUsingJsonMust Panic if error
func (wrappers *Wrappers) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *Wrappers {
	newUsingJson, err :=
		wrappers.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return newUsingJson
}
