package unipath

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"

	"gitlab.com/evatix-go/pathhelper/normalize"
	"gitlab.com/evatix-go/pathhelper/pathext"
	"gitlab.com/evatix-go/pathhelper/pathwrapper"
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
	if receiver.IsFinalized() {
		receiver.finalizedError.HandleErrorWithMsg(
			"cannot add " + anyPath + ". As it is finalized unipath.")
	}

	receiver.collection.Add(anyPath)

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

func (receiver *Wrapper) IsEmpty() bool {
	return receiver.collection.IsEmpty()
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
	receiver.finalizedError = errorwrapper.NewPtr(
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
	toStr := receiver.ToString(receiver.separator, true)

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
