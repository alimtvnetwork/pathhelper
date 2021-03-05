package unipath

import (
	"os"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/msgtype"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"

	"gitlab.com/evatix-go/pathhelper/dirinfo"
	"gitlab.com/evatix-go/pathhelper/fileinfo"
	"gitlab.com/evatix-go/pathhelper/internal/recursiveinternal"
	"gitlab.com/evatix-go/pathhelper/internal/splitinternal"
	"gitlab.com/evatix-go/pathhelper/normalize"
	"gitlab.com/evatix-go/pathhelper/pathext"
	"gitlab.com/evatix-go/pathhelper/pathgetter"
	"gitlab.com/evatix-go/pathhelper/pathwrapper"
	"gitlab.com/evatix-go/pathhelper/recursivepaths"
)

const (
	defaultCapacity = constants.N3
)

type Wrapper struct {
	isFinalized    bool
	finalPath      string
	separator      string
	finalizedError *errorwrapper.Wrapper
	collection     *corestr.Collection
}

func New(sep string) *Wrapper {
	return &Wrapper{
		isFinalized:    false,
		finalPath:      "",
		separator:      sep,
		finalizedError: nil,
		collection:     corestr.NewCollection(defaultCapacity),
	}
}

func NewCap(cap int, sep string) *Wrapper {
	return &Wrapper{
		isFinalized:    false,
		finalPath:      "",
		separator:      sep,
		finalizedError: nil,
		collection:     corestr.NewCollection(cap),
	}
}

func NewCapStartingPath(cap int, sep, startingPath string) *Wrapper {
	wrapper := &Wrapper{
		isFinalized:    false,
		finalPath:      "",
		separator:      sep,
		finalizedError: nil,
		collection:     corestr.NewCollection(cap),
	}

	return wrapper.Add(startingPath)
}

func (receiver *Wrapper) IsFinalized() bool {
	return receiver.isFinalized
}

func (receiver *Wrapper) Lock() {
	receiver.collection.Lock()
}

func (receiver *Wrapper) Unlock() *Wrapper {
	receiver.collection.Unlock()

	return receiver
}

func (receiver *Wrapper) Separator() string {
	return receiver.separator
}

// anyPath can contain separators or without separator both are fine
func (receiver *Wrapper) AddLock(
	anyPath string,
) *Wrapper {
	receiver.Lock()
	defer receiver.Unlock()

	return receiver.Add(anyPath)
}

// anyPath can contain separators or without separator both are fine
// One cannot add path after finalize.
func (receiver *Wrapper) Add(
	anyPath string,
) *Wrapper {
	receiver.handleFinalizeError()

	receiver.collection.Add(anyPath)

	return receiver
}

func (receiver *Wrapper) handleFinalizeError() {
	if receiver.IsFinalized() {
		receiver.finalizedError.HandleErrorWithMsg(
			"Finalized unipath cannot add or modify data.")
	}
}

func (receiver *Wrapper) AddStringsPtr(
	stringItems *[]string,
) *Wrapper {
	receiver.handleFinalizeError()
	receiver.collection.AddStringsPtr(stringItems)

	return receiver
}

func (receiver *Wrapper) AddPointerStringsPtr(
	stringItems *[]*string,
) *Wrapper {
	receiver.handleFinalizeError()
	receiver.collection.AddPointerStringsPtr(stringItems)

	return receiver
}

func (receiver *Wrapper) Length() int {
	return receiver.collection.Length()
}

func (receiver *Wrapper) HasItems() bool {
	return receiver.collection.HasItems()
}

// same one needs to be inserted
func (receiver *Wrapper) Has(
	pathSplit string,
) bool {
	return receiver.collection.Has(pathSplit)
}

func (receiver *Wrapper) IsWindowsSeparator() bool {
	return receiver.separator == constants.WindowsPathSeparator
}

func (receiver *Wrapper) IsUnixSeparator() bool {
	return receiver.separator == constants.ForwardSlash
}

func (receiver *Wrapper) IsEmpty() bool {
	return receiver.collection.IsEmpty()
}

func (receiver *Wrapper) IsValid() bool {
	if receiver.collection.IsEmpty() {
		return false
	}

	_, errW := receiver.GetFileInfo()

	return errW.IsEmpty()
}

func (receiver *Wrapper) IsEqual(wrapper *Wrapper) bool {
	if wrapper == nil && receiver == nil {
		return true
	}

	if wrapper == nil || receiver == nil {
		return false
	}

	if wrapper.isFinalized != receiver.isFinalized {
		return false
	}

	if wrapper.GetFinalizePath() != receiver.GetFinalizePath() {
		return false
	}

	if wrapper.separator != receiver.separator {
		return false
	}

	if wrapper.finalizedError != nil && receiver.finalizedError != nil {
		if !wrapper.finalizedError.IsEquals(receiver.finalizedError) {
			return false
		}
	}

	if wrapper.finalizedError == nil || receiver.finalizedError == nil {
		return false
	}

	return receiver.collection.
		IsEqualsPtr(
			wrapper.collection)
}

func (receiver *Wrapper) Finalize() *errstr.Result {
	if receiver.isFinalized == true {
		// done
		return &errstr.Result{
			Value:        receiver.finalPath,
			ErrorWrapper: receiver.finalizedError,
		}
	}

	// set finalize error
	receiver.isFinalized = true
	receiver.finalizedError = errorwrapper.
		NewPtr(
			errtype.FinalizedResourceCannotAccess)

	finalPath := receiver.
		collection.
		Join(receiver.separator)

	receiver.finalPath = normalize.PathUsingSeparator(
		receiver.separator,
		finalPath,
		true,
		true)

	return &errstr.Result{
		Value:        receiver.finalPath,
		ErrorWrapper: errnew.EmptyPtr,
	}
}

func (receiver *Wrapper) GetFinalizePath() *errstr.Result {
	if !receiver.isFinalized {
		// not finalized
		return &errstr.Result{
			Value:        receiver.finalPath,
			ErrorWrapper: errorwrapper.NewPtr(errtype.CompileFailed),
		}
	}

	return &errstr.Result{
		Value:        receiver.finalPath,
		ErrorWrapper: errnew.EmptyPtr,
	}
}

func (receiver *Wrapper) GetFilesOnPath(
	isNormalize bool,
) *errstr.ResultsWithErrorCollection {
	currentPath := receiver.String()

	return pathgetter.Files(
		receiver.separator,
		currentPath,
		isNormalize)
}

func (receiver *Wrapper) GetBaseDirFiles(
	isNormalize bool,
) *errstr.ResultsWithErrorCollection {
	currentPath := receiver.GetBaseDir()

	return pathgetter.Files(
		receiver.separator,
		currentPath,
		isNormalize)
}

func (receiver *Wrapper) GetDirectoriesOfBaseDir(
	isNormalize bool,
) *errstr.Results {
	currentPath := receiver.GetBaseDir()

	dirs, errWrapper := pathgetter.Dirs(
		receiver.separator,
		currentPath,
		isNormalize)

	return &errstr.Results{
		Values:       dirs,
		ErrorWrapper: errWrapper,
	}
}

// Get Recursive dirs from the basedir path
func (receiver *Wrapper) GetRecursiveDirectoriesOfBaseDir(
	isNormalize bool,
) *errstr.ResultsWithErrorCollection {
	currentPath := receiver.GetBaseDir()

	dirs, errsCollection := recursiveinternal.GetDirectoryPaths(
		receiver.separator,
		currentPath,
		isNormalize)

	return &errstr.ResultsWithErrorCollection{
		Values:        dirs,
		ErrorWrappers: errsCollection,
	}
}

// Get Recursive dirs from the whole path
func (receiver *Wrapper) GetRecursiveDirectories(
	isNormalize bool,
) *errstr.ResultsWithErrorCollection {
	currentPath := receiver.String()

	dirs, errsCollection := recursiveinternal.GetDirectoryPaths(
		receiver.separator,
		currentPath,
		isNormalize)

	return &errstr.ResultsWithErrorCollection{
		Values:        dirs,
		ErrorWrappers: errsCollection,
	}
}

// Get Recursive files from the whole path
func (receiver *Wrapper) GetRecursiveFiles(
	isContinueOnError bool,
) *errstr.ResultsWithErrorCollection {
	currentPath := receiver.String()

	return recursiveinternal.GetFilesPaths(
		receiver.separator,
		currentPath,
		isContinueOnError)
}

// Get Recursive files from the basedir
func (receiver *Wrapper) GetRecursiveFilesOnBaseDir(
	isContinueOnError bool,
) *errstr.ResultsWithErrorCollection {
	currentPath := receiver.GetBaseDir()

	return recursivepaths.Files(
		currentPath,
		isContinueOnError)
}

func (receiver *Wrapper) GetBaseDir() string {
	currentPath := receiver.String()

	return splitinternal.GetBaseDir(
		currentPath)
}

func (receiver *Wrapper) GetBaseDirFileInfo() (os.FileInfo, *errorwrapper.Wrapper) {
	currentPath := receiver.GetBaseDir()
	curFileInfo, err := os.Stat(currentPath)

	if err != nil {
		return curFileInfo,
			errnew.NewPtr(errtype.FileInfo, err)
	}

	return curFileInfo, errnew.EmptyPtr
}

func (receiver *Wrapper) IsBaseDirExists() bool {
	currentFileInfo, errW := receiver.GetBaseDirFileInfo()

	if errW.HasError() {
		return false
	}

	return currentFileInfo.IsDir()
}

func (receiver *Wrapper) IsFileExists() bool {
	currentFileInfo, errW := receiver.GetFileInfo()

	if errW.HasError() {
		return false
	}

	return !currentFileInfo.IsDir()
}

func (receiver *Wrapper) GetBaseDirInfoResult() *dirinfo.Result {
	baseDir := receiver.GetBaseDir()

	return dirinfo.New(baseDir)
}

func (receiver *Wrapper) GetFileInfo() (os.FileInfo, *errorwrapper.Wrapper) {
	filePath := receiver.String()
	curFileInfo, err := os.Stat(filePath)

	if err != nil {
		return curFileInfo, errnew.NewPtr(errtype.FileInfo, err)
	}

	return curFileInfo, errnew.EmptyPtr
}

func (receiver *Wrapper) GetFileInfoWrapper() *fileinfo.Wrapper {
	filePath := receiver.String()

	return fileinfo.New(filePath)
}

func (receiver *Wrapper) Collection() *corestr.Collection {
	return receiver.collection
}

func (receiver *Wrapper) ListPtr() *[]string {
	return receiver.collection.ListPtr()
}

func (receiver *Wrapper) Strings() []string {
	return *receiver.collection.ListPtr()
}

func (receiver *Wrapper) StringsPtr() *[]string {
	return receiver.collection.ListPtr()
}

func (receiver *Wrapper) getFinalizedError() *errorwrapper.Wrapper {
	if receiver.IsFinalized() {
		return receiver.finalizedError
	}

	return errnew.EmptyPtr
}

func (receiver *Wrapper) ToWrapperUpto(
	uptoLastIndexMinus int,
	sep string,
	isNormalize bool,
) *Wrapper {
	currPath := receiver.ToStringUptoLastMinus(
		uptoLastIndexMinus,
		sep,
		isNormalize)

	return New(sep).Add(currPath)
}

func (receiver *Wrapper) ToStringUptoLastMinus(
	uptoLastIndexMinus int,
	sep string,
	isNormalize bool,
) string {
	if uptoLastIndexMinus < 0 {
		msgtype.
			CannotBeNegativeMessage.
			HandleUsingPanic(
				"uptoLastIndexMinus cannot be negative.",
				uptoLastIndexMinus)
	}

	generatedPath := receiver.
		collection.
		Take(receiver.Length() - uptoLastIndexMinus).
		Join(sep)

	generatedPathNext := normalize.PathUsingSeparatorUsingSingleIf(
		isNormalize,
		sep,
		generatedPath)

	return generatedPathNext
}

func (receiver *Wrapper) ToString(
	sep string,
	isNormalize bool,
) string {
	generatedPath := receiver.
		collection.
		Join(sep)

	generatedPathNext := normalize.PathUsingSeparatorUsingSingleIf(
		isNormalize,
		sep,
		generatedPath)

	return generatedPathNext
}

func (receiver *Wrapper) String() string {
	if receiver.IsFinalized() {
		return receiver.finalPath
	}

	toStr := receiver.ToString(
		receiver.separator,
		true)

	return toStr
}

func (receiver *Wrapper) GetAsPathWrapper() *pathwrapper.Wrapper {
	toStr := receiver.String()
	pathWrapper := pathwrapper.Wrapper(toStr)

	return &pathWrapper
}

func (receiver *Wrapper) GetAsPathExt() *pathext.Wrapper {
	toStr := receiver.String()

	return pathext.NewPtr(toStr)
}
