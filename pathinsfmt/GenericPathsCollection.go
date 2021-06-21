package pathinsfmt

import (
	"sort"
	"sync"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/coreinstruction"
)

type GenericPathsCollection struct {
	Specification                      *coreinstruction.Specification `json:"Specification,omitempty"`
	SimilarPaths                       []SimilarPaths                 `json:"SimilarPaths,omitempty"`
	AllDiffPaths                       []AllDiffPaths                 `json:"AllDiffPaths,omitempty"`
	DynamicPaths                       *DynamicPaths                  `json:"DynamicPaths,omitempty"`
	lazyFlatPaths, lazyFlatPathsSorted []string
}

// Length of len(receiver.SimilarPaths) +
// len(receiver.AllDiffPaths) +
// items in DynamicPaths (not all specific paths)
func (receiver *GenericPathsCollection) Length() int {
	length := len(receiver.SimilarPaths) +
		len(receiver.AllDiffPaths)

	if receiver.DynamicPaths == nil {
		return length
	}

	return length + receiver.DynamicPaths.Length()
}

func (receiver *GenericPathsCollection) IsEmpty() bool {
	return receiver.Length() == 0
}

func (receiver *GenericPathsCollection) HasAnyItem() bool {
	return receiver.Length() > 0
}

func (receiver *GenericPathsCollection) LazyFlatPathsSorted() []string {
	if receiver.lazyFlatPathsSorted != nil {
		return receiver.lazyFlatPathsSorted
	}

	lazyPaths := receiver.LazyFlatPaths()
	sort.Strings(lazyPaths)

	receiver.lazyFlatPathsSorted = lazyPaths

	return receiver.lazyFlatPaths
}

func (receiver *GenericPathsCollection) LazyFlatPaths() []string {
	if receiver.lazyFlatPaths != nil {
		return receiver.lazyFlatPaths
	}

	receiver.lazyFlatPaths = receiver.FlatPaths()

	return receiver.lazyFlatPaths
}

func (receiver *GenericPathsCollection) IsEmptySimilarPaths() bool {
	return receiver.SimilarPaths == nil || len(receiver.SimilarPaths) == 0
}

func (receiver *GenericPathsCollection) SimilarPathsIndividualItemsLength() int {
	length := 0

	if receiver.SimilarPaths == nil {
		return 0
	}

	for _, similarPaths := range receiver.SimilarPaths {
		length += similarPaths.Length()
	}

	return length
}

func (receiver *GenericPathsCollection) SimilarPathsFlatPaths() []string {
	if receiver.IsEmptySimilarPaths() {
		return []string{}
	}

	slice := make(
		[]string,
		constants.Zero,
		receiver.SimilarPathsIndividualItemsLength()+constants.ArbitraryCapacity10)

	for _, similarPaths := range receiver.SimilarPaths {
		if similarPaths.IsEmpty() {
			continue
		}

		slice = append(
			slice,
			similarPaths.FlatPaths()...)
	}

	return slice
}

func (receiver *GenericPathsCollection) IsEmptyAllDiffPaths() bool {
	return receiver.AllDiffPaths == nil || len(receiver.AllDiffPaths) == 0
}

func (receiver *GenericPathsCollection) AllDiffPathsIndividualItemsLength() int {
	length := 0

	if receiver.AllDiffPaths == nil {
		return 0
	}

	for _, allDiff := range receiver.AllDiffPaths {
		length += allDiff.Length()
	}

	return length
}

func (receiver *GenericPathsCollection) AllDiffPathsFlatPaths() []string {
	if receiver.IsEmptyAllDiffPaths() {
		return []string{}
	}

	slice := make(
		[]string,
		constants.Zero,
		receiver.AllDiffPathsIndividualItemsLength()+constants.ArbitraryCapacity10)

	for _, allDiffPaths := range receiver.AllDiffPaths {
		if allDiffPaths.IsEmpty() {
			continue
		}

		slice = append(
			slice,
			allDiffPaths.FlatPaths()...)
	}

	return slice
}

func (receiver *GenericPathsCollection) IsEmptyDynamicPaths() bool {
	return receiver.DynamicPaths == nil || len(receiver.DynamicPaths.AllDiffPaths) == 0
}

func (receiver *GenericPathsCollection) DynamicPathsIndividualItemsLength() int {
	length := 0

	if receiver.DynamicPaths == nil {
		return 0
	}

	for _, allDiff := range receiver.DynamicPaths.AllDiffPaths {
		length += allDiff.Length()
	}

	return length
}

func (receiver *GenericPathsCollection) DynamicPathsFlatPaths() []string {
	if receiver.IsEmptyDynamicPaths() {
		return []string{}
	}

	length := receiver.DynamicPathsIndividualItemsLength()

	slice := make(
		[]string,
		constants.Zero,
		length+constants.ArbitraryCapacity10)

	for _, allDiffPaths := range receiver.DynamicPaths.AllDiffPaths {
		if allDiffPaths.IsEmpty() {
			continue
		}

		slice = append(
			slice,
			allDiffPaths.FlatPaths()...)
	}

	return slice
}

func (receiver *GenericPathsCollection) FlatPaths() []string {
	collections := corestr.NewLinkedCollections()
	wg := &sync.WaitGroup{}

	wg.Add(constants.Capacity3)
	collections.AddAsyncFuncItems(
		wg,
		false,
		receiver.AllDiffPathsFlatPaths,
		receiver.DynamicPathsFlatPaths,
		receiver.SimilarPathsFlatPaths,
	)

	return collections.
		ToCollection(constants.Zero).
		ListStrings()
}

func (receiver *GenericPathsCollection) FlatPathsSorted() []string {
	flatPaths := receiver.
		FlatPaths()

	sort.Strings(
		flatPaths)

	return flatPaths
}
