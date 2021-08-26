package fileinfo

import (
	"encoding/json"
	"strings"

	"gitlab.com/evatix-go/core"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coreindexes"
	"gitlab.com/evatix-go/core/defaulterr"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
)

type PathsCollection struct {
	rootPath          string
	pathWrappers      *[]*SimplePathWrapper
	allRecursivePaths *errstr.ResultsWithErrorCollection
	allRecursiveFiles *errstr.ResultsWithErrorCollection
	allRecursiveDirs  *errstr.ResultsWithErrorCollection
	directories       *[]string
	files             *[]string
	separator         string
	ErrorWrapper      *errorwrapper.Wrapper
	parentWrappers    *Wrappers
}

func NewPaths(rootPath, separator string, capacity int) *PathsCollection {
	paths := make([]*SimplePathWrapper, 0, capacity)

	return &PathsCollection{
		rootPath:     rootPath,
		pathWrappers: &paths,
		separator:    separator,
	}
}

func NewPathsUsingWrappers(
	rootPath, separator string,
	wrappers *Wrappers,
) *PathsCollection {
	if wrappers == nil {
		return &PathsCollection{
			rootPath:       rootPath,
			pathWrappers:   &[]*SimplePathWrapper{},
			ErrorWrapper:   errnew.EmptyPtr,
			parentWrappers: wrappers,
			separator:      separator,
		}
	}

	if wrappers.IsEmpty() {
		return &PathsCollection{
			rootPath:       rootPath,
			pathWrappers:   &[]*SimplePathWrapper{},
			ErrorWrapper:   wrappers.ErrorWrapper,
			parentWrappers: wrappers,
			separator:      separator,
		}
	}

	paths := make(
		[]*SimplePathWrapper,
		wrappers.Length())

	for i, wrapper := range *wrappers.Items {
		paths[i] = &SimplePathWrapper{
			Path:        wrapper.RawPath,
			IsDirectory: wrapper.IsDirectory,
		}
	}

	return &PathsCollection{
		rootPath:       rootPath,
		pathWrappers:   &paths,
		ErrorWrapper:   wrappers.ErrorWrapper,
		parentWrappers: wrappers,
		separator:      separator,
	}
}

func NewPathsUsing(
	directoryPath, separator string,
	isNormalize bool,
) *PathsCollection {
	wrappers := NewWrappersPtr(
		directoryPath,
		separator,
		isNormalize)

	return NewPathsUsingWrappers(
		directoryPath,
		separator,
		wrappers)
}

func NewPathsUsingPaths(
	rootPath, separator string,
	recursivePaths *[]string,
) *PathsCollection {
	if recursivePaths == nil {
		return NewPaths(
			rootPath,
			separator,
			0)
	}

	wrappers :=
		NewPaths(
			rootPath,
			separator,
			len(*recursivePaths))

	wrappers.allRecursivePaths =
		&errstr.ResultsWithErrorCollection{
			Values:        recursivePaths,
			ErrorWrappers: errwrappers.Empty(),
		}

	return wrappers
}

func (pathsCollection *PathsCollection) Directories() *[]string {
	if pathsCollection.directories != nil {
		return pathsCollection.directories
	}

	if pathsCollection.IsEmpty() {
		pathsCollection.directories =
			core.EmptyStringsPtr()

		return pathsCollection.directories
	}

	directories := make([]string, 0, pathsCollection.Length())

	for _, pathWrapper := range *pathsCollection.pathWrappers {
		if !pathWrapper.IsDirectory {
			continue
		}

		directories = append(directories, pathWrapper.Path)
	}

	pathsCollection.directories = &directories

	return pathsCollection.directories
}

func (pathsCollection *PathsCollection) Files() *[]string {
	if pathsCollection.files != nil {
		return pathsCollection.files
	}

	if pathsCollection.IsEmpty() {
		pathsCollection.files = core.EmptyStringsPtr()

		return pathsCollection.files
	}

	files := make([]string, 0, pathsCollection.Length())

	for _, pathWrapper := range *pathsCollection.pathWrappers {
		if pathWrapper.IsDirectory {
			continue
		}

		files = append(files, pathWrapper.Path)
	}

	pathsCollection.files = &files

	return pathsCollection.files
}

func (pathsCollection *PathsCollection) IsEmpty() bool {
	return pathsCollection.pathWrappers == nil ||
		pathsCollection.ErrorWrapper.HasError() ||
		len(*pathsCollection.pathWrappers) == 0
}

func (pathsCollection *PathsCollection) Length() int {
	if pathsCollection.pathWrappers == nil || *pathsCollection.pathWrappers == nil {
		return 0
	}

	return len(*pathsCollection.pathWrappers)
}

func (pathsCollection *PathsCollection) IsParentWrappersEmpty() bool {
	return pathsCollection.parentWrappers == nil ||
		pathsCollection.parentWrappers.IsEmpty()
}

func (pathsCollection *PathsCollection) HasParentWrappers() bool {
	return pathsCollection.parentWrappers != nil
}

func (pathsCollection *PathsCollection) ParentWrappers() *Wrappers {
	return pathsCollection.parentWrappers
}

func (pathsCollection *PathsCollection) Add(
	wrapper *SimplePathWrapper,
) *PathsCollection {
	*pathsCollection.pathWrappers = append(
		*pathsCollection.pathWrappers,
		wrapper)

	return pathsCollection
}

func (pathsCollection *PathsCollection) AddPtr(
	wrapper *SimplePathWrapper,
) *PathsCollection {
	*pathsCollection.pathWrappers = append(
		*pathsCollection.pathWrappers,
		wrapper)

	return pathsCollection
}

func (pathsCollection *PathsCollection) AddWrapper(
	pathWrapper *SimplePathWrapper,
) *PathsCollection {
	*pathsCollection.pathWrappers = append(
		*pathsCollection.pathWrappers,
		pathWrapper)

	return pathsCollection
}

func (pathsCollection *PathsCollection) Strings() *[]string {
	list := make(
		[]string,
		pathsCollection.Length())

	for i, wrapper := range *pathsCollection.pathWrappers {
		list[i] = wrapper.String()
	}

	return &list
}

func (pathsCollection *PathsCollection) String() string {
	list := make(
		[]string,
		constants.ArbitraryCapacity4)
	compiledPaths := strings.Join(
		*pathsCollection.Strings(),
		constants.NewLineUnix)

	list[coreindexes.I0] = "Root Location :" + pathsCollection.rootPath
	list[coreindexes.I1] = "Separator :" + pathsCollection.separator
	if pathsCollection.ErrorWrapper.HasError() {
		list[coreindexes.I2] = "Error :" + pathsCollection.ErrorWrapper.
			String()
	}

	list[coreindexes.I3] = compiledPaths

	return strings.Join(
		list,
		constants.NewLineUnix)
}

func (pathsCollection *PathsCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(*pathsCollection.JsonModel())
}

func (pathsCollection *PathsCollection) UnmarshalJSON(data []byte) error {
	var dataModel PathsCollectionDataModel
	err := json.Unmarshal(data, &dataModel)

	if err == nil {
		pathsCollection.rootPath = dataModel.RootPath
		pathsCollection.pathWrappers = dataModel.PathWrappers
		pathsCollection.separator = dataModel.Separator
		pathsCollection.ErrorWrapper = dataModel.ErrorWrapper
		pathsCollection.parentWrappers = dataModel.ParentWrappers
	}

	return err
}

func (pathsCollection *PathsCollection) JsonModel() *PathsCollectionDataModel {
	return &PathsCollectionDataModel{
		RootPath:       pathsCollection.rootPath,
		PathWrappers:   pathsCollection.pathWrappers,
		Separator:      pathsCollection.separator,
		ErrorWrapper:   pathsCollection.ErrorWrapper,
		ParentWrappers: pathsCollection.parentWrappers,
	}
}

func (pathsCollection *PathsCollection) JsonModelAny() interface{} {
	return pathsCollection.JsonModel()
}

func (pathsCollection *PathsCollection) Json() *corejson.Result {
	return corejson.NewFromAny(pathsCollection)
}

//goland:noinspection GoLinterLocal
func (pathsCollection *PathsCollection) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*PathsCollection, error) {
	if jsonResult == nil || jsonResult.IsEmptyJsonBytes() {
		return nil, defaulterr.UnMarshallingFailedDueToNilOrEmpty
	}

	err := json.Unmarshal(*jsonResult.Bytes, &pathsCollection)

	if err != nil {
		return nil, err
	}

	return pathsCollection, nil
}

// Panic if error
//goland:noinspection GoLinterLocal
func (pathsCollection *PathsCollection) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *PathsCollection {
	newUsingJson, err :=
		pathsCollection.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return newUsingJson
}

func (pathsCollection *PathsCollection) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := pathsCollection.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (pathsCollection *PathsCollection) AsJsoner() corejson.Jsoner {
	return pathsCollection
}

func (pathsCollection *PathsCollection) AsJsonParseSelfInjector() corejson.JsonParseSelfInjector {
	return pathsCollection
}
