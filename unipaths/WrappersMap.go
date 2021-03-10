package unipaths

import (
	"sync"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"

	"gitlab.com/evatix-go/pathhelper/unipath"
)

type WrappersMap struct {
	separator   string
	isFinalized bool
	sync.Mutex
	items *map[string]*unipath.Wrapper
}

func (receiver *WrappersMap) IsFinalized() bool {
	return receiver.isFinalized
}

func (receiver *WrappersMap) Lock() {
	receiver.Lock()
}

func (receiver *WrappersMap) Unlock() {
	receiver.Unlock()
}

func (receiver *WrappersMap) AddLock(
	key,
	anyPath string,
) *WrappersMap {
	receiver.Lock()
	defer receiver.Unlock()

	return receiver.addLock(key, anyPath)
}

func (receiver *WrappersMap) addLock(
	key,
	anyPath string,
) *WrappersMap {
	wrapper, has := (*receiver.items)[key]

	if !has {
		wrapper = unipath.New(receiver.separator)
		(*receiver.items)[key] = wrapper
	}

	wrapper.AddLock(anyPath)

	return receiver
}

func (receiver *WrappersMap) Add(
	key,
	anyPath string,
) *WrappersMap {
	wrapper, has := (*receiver.items)[key]

	if !has {
		wrapper = unipath.New(receiver.separator)
		(*receiver.items)[key] = wrapper
	}

	wrapper.Add(anyPath)

	return receiver
}

func (receiver *WrappersMap) Length() int {
	return len(*receiver.items)
}

func (receiver *WrappersMap) HasItems() bool {
	return receiver.Length() > constants.Zero
}

// same one needs to be inserted
func (receiver *WrappersMap) Has(
	key,
	pathSplit string,
) bool {
	wrapper, has := (*receiver.items)[key]

	return has && wrapper.Has(pathSplit)
}

func (receiver *WrappersMap) IsEmpty() bool {
	return receiver.Length() == constants.Zero
}

func (receiver *WrappersMap) IsEqual(wrappersMap *WrappersMap) bool {
	if wrappersMap == nil && receiver == nil {
		return true
	}

	if wrappersMap == nil || receiver == nil {
		return false
	}

	if wrappersMap == receiver {
		return true
	}

	if wrappersMap.isFinalized != receiver.isFinalized {
		return false
	}

	if wrappersMap.Length() != receiver.Length() {
		return false
	}

	if receiver.items == wrappersMap.items {
		return true
	}

	for key, receiverWrapper := range *receiver.items {
		anotherWrapper, has := (*wrappersMap.items)[key]

		if !has {
			return false
		}

		if !anotherWrapper.IsEqual(receiverWrapper) {
			return false
		}
	}

	return true
}

// FinalizeAll all wrappers
func (receiver *WrappersMap) FinalizeAll() {
	if receiver.isFinalized {
		return
	}

	// set finalize error
	receiver.isFinalized = true
	for _, wrapper := range *receiver.items {
		wrapper.Finalize()
	}
}

func (receiver *WrappersMap) GetFinalizePath(
	key string,
) *errstr.Result {
	wrapper, has := (*receiver.items)[key]

	if has {
		return wrapper.GetFinalizePath()
	}

	return errstr.ErrorPtr(
		errtype.
			NotContainsExpectation.
			ErrorNoRefs(key))
}

func (receiver *WrappersMap) GetFinalizePaths() *errstr.ResultsWithErrorCollection {
	if !receiver.IsFinalized() {
		return errstr.NewResultsWithErrorCollectionUsingTypeMessagePtr(
			errtype.Unexpected,
			nonFinalizePathsCannotBeRetrievedMessage)
	}

	length := receiver.Length()
	list := make(
		[]string,
		length,
		length)

	errCollection := errwrappers.Empty()

	i := constants.Zero
	for _, wrapper := range *receiver.items {
		finalizedResult := wrapper.GetFinalizePath()
		errCollection.AddWrapperPtr(finalizedResult.ErrorWrapper)
		list[i] = finalizedResult.Value

		i++
	}

	return &errstr.ResultsWithErrorCollection{
		Values:        &list,
		ErrorWrappers: errCollection,
	}
}

func (receiver *WrappersMap) Items() *map[string]*unipath.Wrapper {
	return receiver.items
}

func (receiver *WrappersMap) ListPtr() *[]*unipath.Wrapper {
	list := make([]*unipath.Wrapper, receiver.Length())

	i := constants.Zero
	for _, wrapper := range *receiver.items {
		list[i] = wrapper
		i++
	}

	return &list
}

func (receiver *WrappersMap) ToStringsPtr(
	separator string,
	isNormalize bool,
) *[]string {
	list := make([]string, receiver.Length())

	i := constants.Zero
	for _, wrapper := range *receiver.items {
		list[i] = wrapper.ToString(
			separator,
			isNormalize)
		i++
	}

	return &list
}

func (receiver *WrappersMap) StringsPtr() *[]string {
	list := make([]string, receiver.Length())

	i := constants.Zero
	for _, wrapper := range *receiver.items {
		list[i] = wrapper.String()
		i++
	}

	return &list
}

func (receiver *WrappersMap) Get(
	key string,
) *unipath.Wrapper {
	return (*receiver.items)[key]
}
