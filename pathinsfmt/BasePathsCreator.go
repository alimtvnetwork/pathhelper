package pathinsfmt

import (
	"os"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/createpath"
	"gitlab.com/evatix-go/pathhelper/internal/normalizeinternal"
	"gitlab.com/evatix-go/pathhelper/pathchmod"
)

type BasePathsCreator struct {
	RootDir                     string `json:"RootDir,omitempty"`
	Files                       []string
	IsNormalize                 bool
	lazyFlatFiles               []string
	lazyPathsChmod              *errstr.Hashmap
	lazyFilteredPathFileInfoMap *chmodhelper.FilteredPathFileInfoMap
}

func (it *BasePathsCreator) SimilarPaths() *SimilarPaths {
	return &SimilarPaths{
		RootPath:         it.RootDir,
		RelativePaths:    it.Files,
		IsNormalizeApply: it.IsNormalize,
	}
}

func (it *BasePathsCreator) LazyPathsChmod() *errstr.Hashmap {
	if it.lazyPathsChmod != nil {
		return it.lazyPathsChmod
	}

	it.lazyPathsChmod = it.PathsChmodMap()

	return it.lazyPathsChmod
}

func (it *BasePathsCreator) LazyFlatPaths() []string {
	if it.lazyFlatFiles != nil {
		return it.lazyFlatFiles
	}

	it.lazyFlatFiles = it.FlatPaths()

	return it.lazyFlatFiles
}

func (it *BasePathsCreator) Length() int {
	return len(it.Files)
}

func (it *BasePathsCreator) LengthPlusRoot() int {
	return len(it.Files) + 1
}

func (it *BasePathsCreator) IsEmpty() bool {
	return len(it.Files) == 0
}

func (it *BasePathsCreator) HasAnyItem() bool {
	return len(it.Files) > 0
}

func (it *BasePathsCreator) FlatPaths() []string {
	return *it.FlatPathsPtr()
}

func (it *BasePathsCreator) FlatPathsPtr() *[]string {
	slice := make([]string, it.Length())

	for i, file := range it.Files {
		compiledPath := normalizeinternal.JoinFixIf(
			it.IsNormalize,
			it.RootDir,
			file)
		slice[i] = compiledPath
	}

	return &slice
}

// DeleteAllPaths delete all files in root path
func (it *BasePathsCreator) DeleteAllPaths() *errorwrapper.Wrapper {
	location := it.RootDir
	err := os.RemoveAll(location)

	return errnew.Path(errtype.DeletePathFailed, err, location)
}

func (it *BasePathsCreator) PathsChmodMap() *errstr.Hashmap {
	files := it.FlatPathsPtr()
	hashmap, err := chmodhelper.
		GetFilesChmodRwxFullMap(*files)

	return &errstr.Hashmap{
		Hashmap: hashmap,
		ErrorWrapper: errnew.NewPtr(
			errtype.ExistingChmodReadFailed,
			err),
	}
}

func (it *BasePathsCreator) CreatePaths(mode os.FileMode) (
	[]*os.File,
	*errorwrapper.Wrapper,
) {
	return it.createFiles(
		mode,
		it.FlatPaths())
}

func (it *BasePathsCreator) CreateLazyPathsWithoutMode() (
	[]*os.File,
	*errorwrapper.Wrapper,
) {
	return createpath.CreateMany(
		false,
		it.LazyFlatPaths())
}

func (it *BasePathsCreator) CreatePathsWithoutMode() (
	[]*os.File,
	*errorwrapper.Wrapper,
) {
	return createpath.CreateMany(
		false,
		it.FlatPaths())
}

func (it *BasePathsCreator) ApplyLinuxRecursiveFileModeOnRoot(
	fileMode os.FileMode,
) *errorwrapper.Wrapper {
	return pathchmod.ApplyLinuxRecursiveChmodOnPathUsingFileMode(
		fileMode,
		it.RootDir)
}

func (it *BasePathsCreator) CreateLazyFlatFiles(mode os.FileMode) (
	[]*os.File,
	*errorwrapper.Wrapper,
) {
	return it.createFiles(
		mode,
		it.LazyFlatPaths())
}

func (it *BasePathsCreator) DeleteAllThenCreateLazyFlatFiles(
	mode os.FileMode,
) (
	[]*os.File,
	*errorwrapper.Wrapper,
) {
	deleteAllErr := it.DeleteAllPaths()

	if deleteAllErr.HasError() {
		return []*os.File{}, errnew.EmptyPtr
	}

	return it.createFiles(
		mode,
		it.LazyFlatPaths())
}

func (it *BasePathsCreator) DeleteAllThenCreateFlatFiles(
	mode os.FileMode,
) (
	[]*os.File,
	*errorwrapper.Wrapper,
) {
	deleteAllErr := it.DeleteAllPaths()

	if deleteAllErr.HasError() {
		return []*os.File{}, errnew.EmptyPtr
	}

	return it.createFiles(
		mode,
		it.FlatPaths())
}

func (it *BasePathsCreator) createFiles(
	mode os.FileMode,
	files []string,
) ([]*os.File, *errorwrapper.Wrapper) {
	if len(files) == 0 {
		return []*os.File{}, errnew.EmptyPtr
	}

	return createpath.CreateManySameDirWithFileMode(
		mode,
		false,
		it.RootDir,
		files,
	)
}

func (it *BasePathsCreator) GetFilesInfoMap(
	isSkipOnInvalid bool,
) *chmodhelper.FilteredPathFileInfoMap {
	files := it.FlatPathsPtr()

	return chmodhelper.
		GetExistsFilteredPathFileInfoMap(
			isSkipOnInvalid,
			*files)
}
