package pathgetter

import (
	"sync"

	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/coreindexes"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
)

func FilesUsingPathsPtr(
	separator string,
	isNormalize bool,
	exploringPaths *[]string,
) *errstr.ResultsWithErrorCollection {
	length := corestr.LengthOfStrings(exploringPaths)

	if length == 0 {
		return errstr.
			EmptyResultsWithErrorCollectionPtr()
	}

	if length == 1 {
		exploringPath :=
			(*exploringPaths)[coreindexes.First]

		return Files(
			separator,
			exploringPath,
			isNormalize)
	}

	linkedCollections :=
		corestr.NewLinkedCollections()
	wg := &sync.WaitGroup{}
	wg.Add(length)
	errWrappers := errwrappers.NewCap2()

	for _, eachExploringPath := range *exploringPaths {
		allPaths := Files(
			separator,
			eachExploringPath,
			isNormalize)

		if allPaths.HasError() {
			errWrappers.AddCollections(allPaths.ErrorWrappers)
			wg.Done()

			continue
		}

		linkedCollections.AddStringsPtrAsync(
			wg,
			allPaths.ValueMust(),
			false)
	}

	wg.Wait()

	return &errstr.ResultsWithErrorCollection{
		Values:        linkedCollections.ListPtr(),
		ErrorWrappers: errWrappers,
	}
}
