package pathinsfmt

import (
	"sync"

	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
)

type PathsCreatorCollection struct {
	PathsCreateInstructions []BasePathsCreator `json:"PathsCreateInstructions,omitempty"`
	IsIgnoreOnExist         bool
	IsDeleteAllBeforeCreate bool
	ApplyRwx                *chmodins.RwxOwnerGroupOther
	ApplyUserGroup          *UserGroupName
	lazyFlatPaths           []string
	lazyPathsCreators       []*PathsCreator
}

func (it *PathsCreatorCollection) LazyPathsCreators() []*PathsCreator {
	if it.lazyPathsCreators != nil {
		return it.lazyPathsCreators
	}

	it.lazyPathsCreators = it.PathsCreators()

	return it.LazyPathsCreators()
}

func (it *PathsCreatorCollection) PathsCreators() []*PathsCreator {
	slice := make([]*PathsCreator, it.Length())

	for i, instruction := range it.PathsCreateInstructions {
		slice[i] = &PathsCreator{
			BasePathsCreator: instruction,
			ApplyRwx:         it.ApplyRwx,
			ApplyUserGroup:   it.ApplyUserGroup,
		}
	}

	return slice
}

func (it *PathsCreatorCollection) LazyFlatPathsIf(isLazy bool) []string {
	if isLazy {
		return it.LazyFlatPaths()
	}

	return it.FlatPaths()
}

func (it *PathsCreatorCollection) LazyFlatPaths() []string {
	if it.lazyFlatPaths != nil {
		return it.lazyFlatPaths
	}

	it.lazyFlatPaths = it.FlatPaths()

	return it.lazyFlatPaths
}

// Length yields count of PathsCreateInstructions, not all paths count
func (it *PathsCreatorCollection) Length() int {
	return len(it.PathsCreateInstructions)
}

func (it *PathsCreatorCollection) IsEmpty() bool {
	return len(it.PathsCreateInstructions) == 0
}

func (it *PathsCreatorCollection) HasAnyItem() bool {
	return len(it.PathsCreateInstructions) > 0
}

func (it *PathsCreatorCollection) HasRwx() bool {
	return it.ApplyRwx != nil
}

func (it *PathsCreatorCollection) HasUserGroup() bool {
	return it.ApplyUserGroup != nil &&
		it.ApplyUserGroup.HasUserNameOrGroup()
}

func (it *PathsCreatorCollection) FlatPaths() []string {
	length := it.Length()

	if length == 0 {
		return []string{}
	}

	collection := corestr.NewLinkedCollections()

	wg3 := &sync.WaitGroup{}

	for _, createInstruction := range it.PathsCreateInstructions {
		wg3.Add(1)
		collection.AddAsyncFuncItemsPointer(
			wg3,
			false,
			createInstruction.FlatPathsPtr)
	}

	return collection.
		ToCollection(constants.Zero).
		ListStrings()
}
