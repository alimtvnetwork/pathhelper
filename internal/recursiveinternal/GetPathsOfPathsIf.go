package recursiveinternal

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/constants/percentconst"
	"gitlab.com/evatix-go/core/defaultcapacity"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
)

func GetPathsOfPathsIf(
	isRecursive bool,
	locations []string,
	isContinueOnEmpty bool,
) *errstr.ResultsWithErrorCollection {
	length := len(locations)
	if length == 0 {
		return errstr.EmptyResultsWithErrorCollectionPtr()
	}

	if !isRecursive {
		return &errstr.ResultsWithErrorCollection{
			Values:        &locations,
			ErrorWrappers: errwrappers.Empty(),
		}
	}

	capacity := defaultcapacity.Predictive(
		length,
		percentconst.DoubleIncrement,
		constants.ArbitraryCapacity50)

	errCollection := errwrappers.Empty()
	addLocations := make(
		[]string,
		constants.Zero,
		capacity)

	for _, location := range locations {
		locationsRecursive, errCollect := GetPathsWithoutSeparator(location, isContinueOnEmpty)

		errCollection.AddCollections(errCollect)

		if locationsRecursive != nil && len(*locationsRecursive) > 0 {
			addLocations = append(addLocations, *locationsRecursive...)
		}
	}

	return &errstr.ResultsWithErrorCollection{
		Values:        &addLocations,
		ErrorWrappers: errCollection,
	}
}
