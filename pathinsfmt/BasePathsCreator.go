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
)

type BasePathsCreator struct {
	RootDir                     string `json:"RootDir,omitempty"`
	Files                       []string
	IsNormalize                 bool
	lazyFlatFiles               []string
	lazyFilesChmod              *errstr.Hashmap
	lazyFilteredPathFileInfoMap *chmodhelper.FilteredPathFileInfoMap
}

func (it *BasePathsCreator) SimilarPaths() *SimilarPaths {
	return &SimilarPaths{
		RootPath:         it.RootDir,
		RelativePaths:    it.Files,
		IsNormalizeApply: it.IsNormalize,
	}
}

func (it *BasePathsCreator) LazyFilesChmod() *errstr.Hashmap {
	if it.lazyFilesChmod != nil {
		return it.lazyFilesChmod
	}

	it.lazyFilesChmod = it.GetFilesChmodMap()

	return it.lazyFilesChmod
}

func (it *BasePathsCreator) LazyFlatFiles() []string {
	if it.lazyFlatFiles != nil {
		return it.lazyFlatFiles
	}

	it.lazyFlatFiles = it.FlatFiles()

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

func (it *BasePathsCreator) FlatFiles() []string {
	return *it.FlatFilesPtr()
}

func (it *BasePathsCreator) FlatFilesPtr() *[]string {
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

// DeleteAllFiles delete all files in root path
func (it *BasePathsCreator) DeleteAllFiles() *errorwrapper.Wrapper {
	location := it.RootDir
	err := os.RemoveAll(location)

	return errnew.Path(errtype.DeletePathFailed, err, location)
}

func (it *BasePathsCreator) GetFilesChmodMap() *errstr.Hashmap {
	files := it.FlatFilesPtr()
	hashmap, err := chmodhelper.
		GetFilesChmodRwxFullMap(*files)

	return &errstr.Hashmap{
		Hashmap: hashmap,
		ErrorWrapper: errnew.NewPtr(
			errtype.ExistingChmodReadFailed,
			err),
	}
}

func (it *BasePathsCreator) CreateFiles(mode os.FileMode) (
	[]*os.File,
	*errorwrapper.Wrapper,
) {
	return it.createFiles(
		mode,
		it.FlatFiles())
}

func (it *BasePathsCreator) CreateLazyFlatFiles(mode os.FileMode) (
	[]*os.File,
	*errorwrapper.Wrapper,
) {
	return it.createFiles(
		mode,
		it.LazyFlatFiles())
}

func (it *BasePathsCreator) DeleteAllThenCreateLazyFlatFiles(
	mode os.FileMode,
) (
	[]*os.File,
	*errorwrapper.Wrapper,
) {
	deleteAllErr := it.DeleteAllFiles()

	if deleteAllErr.HasError() {
		return []*os.File{}, errnew.EmptyPtr
	}

	return it.createFiles(
		mode,
		it.LazyFlatFiles())
}


func (it *BasePathsCreator) DeleteAllThenCreateFlatFiles(
	mode os.FileMode,
) (
	[]*os.File,
	*errorwrapper.Wrapper,
) {
	deleteAllErr := it.DeleteAllFiles()

	if deleteAllErr.HasError() {
		return []*os.File{}, errnew.EmptyPtr
	}

	return it.createFiles(
		mode,
		it.FlatFiles())
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
	files := it.FlatFilesPtr()

	return chmodhelper.
		GetExistsFilteredPathFileInfoMap(
			isSkipOnInvalid,
			*files)
}
