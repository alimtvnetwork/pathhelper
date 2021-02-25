package pathgetter

import (
	"sync"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

func AllPtr(
	separator string,
	isNormalize bool,
	exploringPaths *[]string,
) (*[]string, *errorwrapper.Wrapper) {
	length := corestr.LengthOfStrings(exploringPaths)

	if length == 0 {
		return &(constants.EmptyStrings),
			errnew.EmptyPtr
	}

	if length == 1 {
		return AllOfSinglePath(
			separator,
			isNormalize,
			(*exploringPaths)[0])
	}

	linkedCollection := corestr.NewLinkedCollections()
	wg := &sync.WaitGroup{}
	wg.Add(length)

	for _, expPath := range *exploringPaths {
		allPaths, errW := AllOfSinglePath(
			separator,
			isNormalize,
			expPath)

		if errW.HasError() {
			return linkedCollection.ListPtr(),
				errW
		}

		linkedCollection.AddStringsPtrAsync(
			wg,
			allPaths,
			false)
	}

	wg.Wait()

	return linkedCollection.ListPtr(), errnew.EmptyPtr
}
