package pathgetter

import (
	"sync"

	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

func AllPtr(
	separator string,
	isNormalize bool,
	exploringPaths []string,
) *errstr.Results {
	length := corestr.LengthOfStrings(exploringPaths)

	if length == 0 {
		return errstr.EmptyResults()
	}

	if length == 1 {
		return AllOfSinglePath(
			isNormalize,
			separator,
			exploringPaths[0])
	}

	linkedCollection := corestr.NewLinkedCollections()
	wg := &sync.WaitGroup{}

	for _, expPath := range exploringPaths {
		wg.Add(1)
		allPaths := AllOfSinglePath(
			isNormalize,
			separator,
			expPath)

		if allPaths.HasError() {
			return &errstr.Results{
				Values:       *linkedCollection.ListPtr(),
				ErrorWrapper: allPaths.ErrorWrapper,
			}
		}

		linkedCollection.AddStringsPtrAsync(
			wg,
			allPaths.ValueMust(),
			false)
	}

	wg.Wait()

	return &errstr.Results{
		Values:       *linkedCollection.ListPtr(),
		ErrorWrapper: errnew.EmptyPtr,
	}
}
