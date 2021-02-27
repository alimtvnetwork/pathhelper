package pathfilter

import (
	"sync"

	"gitlab.com/evatix-go/core/coredata/corestr"
)

func getRecursiveFilterDefinitions(
	arg *recursiveFilterGetterParam,
) *[]string {
	linkedCollection :=
		corestr.NewLinkedCollections()
	wg := &sync.WaitGroup{}
	wg.Add(arg.additionalFiltersLength)

	for _, filterPath := range *arg.additionalFilters {
		rootPathPlusFilterPath :=
			arg.rootPathPlusSeparator +
				filterPath
		arg.eachFilterPath =
			rootPathPlusFilterPath

		newFilters :=
			getRecursiveFilterForEachFilterPath(arg)

		linkedCollection.AddStringsPtrAsync(
			wg,
			newFilters,
			false)
	}

	wg.Wait()

	return linkedCollection.ListPtr()
}
