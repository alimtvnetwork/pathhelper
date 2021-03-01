package unipaths

import (
	"sync"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errtype"

	"gitlab.com/evatix-go/pathhelper/unipath"
)

const (
	defaultCapacity = constants.N5
)

type WrappersMap struct {
	separator   string
	isFinalized bool
	sync.Mutex
	items *map[string]*unipath.Wrapper
}

func New(
	sep string,
) *WrappersMap {
	list := make(
		map[string]*unipath.Wrapper,
		defaultCapacity)

	return &WrappersMap{
		separator: sep,
		items:     &list,
	}
}

func NewCap(
	cap int, sep string,
) *WrappersMap {
	list := make(
		map[string]*unipath.Wrapper,
		cap)

	return &WrappersMap{
		separator: sep,
		items:     &list,
	}
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
	return receiver.Length() > 0
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
	return receiver.Length() == 0
}

func (receiver *WrappersMap) IsEqual(wrappersMap *WrappersMap) bool {
	if wrappersMap == nil && receiver == nil {
		return true
	}

	if wrappersMap == nil || receiver == nil {
		return false
	}

	if wrappersMap.isFinalized != receiver.isFinalized {
		return false
	}

	if wrappersMap.Length() != receiver.Length() {
		return false
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

func (receiver *WrappersMap) Finalize() {
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

func (receiver *WrappersMap) Items() *map[string]*unipath.Wrapper {
	return receiver.items
}

func (receiver *WrappersMap) ListPtr() *[]*unipath.Wrapper {
	list := make([]*unipath.Wrapper, receiver.Length())

	i := 0
	for _, wrapper := range *receiver.items {
		list[i] = wrapper
		i++
	}

	return &list
}

func (receiver *WrappersMap) StringsPtr() *[]string {
	list := make([]string, receiver.Length())

	i := 0
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
