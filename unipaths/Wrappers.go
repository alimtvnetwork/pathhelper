package unipaths

import (
	"sync"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"

	"gitlab.com/evatix-go/pathhelper/unipath"
)

type Wrappers struct {
	separator   string
	isFinalized bool
	sync.Mutex
	items *[]*unipath.Wrapper
}

func (receiver *Wrappers) IsFinalized() bool {
	return receiver.isFinalized
}

func (receiver *Wrappers) Lock() {
	receiver.Lock()
}

func (receiver *Wrappers) Unlock() {
	receiver.Unlock()
}

func (receiver *Wrappers) AddPathAs(
	givenPath string,
) *Wrappers {
	wrapper := unipath.NewUsingPath(
		givenPath,
		receiver.separator)

	*receiver.items = append(
		*receiver.items,
		wrapper)

	return receiver
}

func (receiver *Wrappers) AddPathsLock(
	givenPaths ...string,
) *Wrappers {
	receiver.Lock()
	defer receiver.Unlock()

	if givenPaths == nil {
		return receiver
	}

	return receiver.
		AddPathsPtr(&givenPaths)
}

func (receiver *Wrappers) AddPaths(
	givenPaths ...string,
) *Wrappers {
	if givenPaths == nil {
		return receiver
	}

	return receiver.
		AddPathsPtr(&givenPaths)
}

func (receiver *Wrappers) AddPathsPtr(
	givenPaths *[]string,
) *Wrappers {
	if givenPaths == nil {
		return receiver
	}

	for _, currentPath := range *givenPaths {
		wrapper := unipath.NewUsingPath(
			currentPath,
			receiver.separator)

		*receiver.items = append(
			*receiver.items,
			wrapper,
		)
	}

	return receiver
}

func (receiver *Wrappers) AddWrapper(
	wrapper *unipath.Wrapper,
) *Wrappers {
	if wrapper == nil {
		return receiver
	}

	*receiver.items = append(
		*receiver.items,
		wrapper)

	return receiver
}

func (receiver *Wrappers) AddWrapperLock(
	wrapper *unipath.Wrapper,
) *Wrappers {
	receiver.Lock()
	defer receiver.Unlock()

	if wrapper == nil {
		return receiver
	}

	*receiver.items = append(
		*receiver.items,
		wrapper)

	return receiver
}

func (receiver *Wrappers) Length() int {
	return len(*receiver.items)
}

func (receiver *Wrappers) HasItems() bool {
	return receiver.Length() > 0
}

func (receiver *Wrappers) IsEmpty() bool {
	return receiver.Length() == 0
}

func (receiver *Wrappers) IsEqual(wrappers *Wrappers) bool {
	if wrappers == nil && receiver == nil {
		return true
	}

	if wrappers == nil || receiver == nil {
		return false
	}

	if wrappers == receiver {
		return true
	}

	if wrappers.isFinalized != receiver.isFinalized {
		return false
	}

	if wrappers.Length() != receiver.Length() {
		return false
	}

	if receiver.items == wrappers.items {
		return true
	}

	for index, receiverWrapper := range *receiver.items {
		anotherWrapper := (*wrappers.items)[index]

		if anotherWrapper == nil && receiverWrapper == nil {
			continue
		}

		if anotherWrapper == nil || receiverWrapper == nil {
			return false
		}

		if !anotherWrapper.IsEqual(receiverWrapper) {
			return false
		}
	}

	return true
}

func (receiver *Wrappers) FinalizeAll() {
	if receiver.isFinalized {
		return
	}

	// set finalize error
	for _, wrapper := range *receiver.items {
		wrapper.Finalize()
	}

	receiver.isFinalized = true
}

func (receiver *Wrappers) GetFinalizePaths() *errstr.ResultsWithErrorCollection {
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

	for i, wrapper := range *receiver.items {
		finalizedResult := wrapper.GetFinalizePath()
		errCollection.AddWrapperPtr(finalizedResult.ErrorWrapper)
		list[i] = finalizedResult.Value
	}

	return &errstr.ResultsWithErrorCollection{
		Values:        &list,
		ErrorWrappers: errCollection,
	}
}

func (receiver *Wrappers) Items() *[]*unipath.Wrapper {
	return receiver.items
}

func (receiver *Wrappers) ListPtr() *[]*unipath.Wrapper {
	list := make([]*unipath.Wrapper, receiver.Length())

	i := 0
	for _, wrapper := range *receiver.items {
		list[i] = wrapper
		i++
	}

	return &list
}

func (receiver *Wrappers) ToStringsPtr(
	separator string,
	isNormalize bool,
) *[]string {
	list := make([]string, receiver.Length())

	i := 0
	for _, wrapper := range *receiver.items {
		list[i] = wrapper.ToString(
			separator,
			isNormalize)
		i++
	}

	return &list
}

func (receiver *Wrappers) StringsPtr() *[]string {
	list := make([]string, receiver.Length())

	i := constants.Zero
	for _, wrapper := range *receiver.items {
		list[i] = wrapper.String()

		i++
	}

	return &list
}

func (receiver *Wrappers) StringsCollectionPtr() *corestr.Collection {
	return corestr.NewCollectionUsingStrings(
		receiver.StringsPtr(),
		false,
	)
}

func (receiver *Wrappers) GetAt(
	index int,
) *unipath.Wrapper {
	return (*receiver.items)[index]
}

func (receiver *Wrappers) GetSafeAt(
	index int,
) *unipath.Wrapper {
	if index > constants.InvalidNotFoundCase && index <= receiver.Length()-1 {
		return (*receiver.items)[index]
	}

	return nil
}
