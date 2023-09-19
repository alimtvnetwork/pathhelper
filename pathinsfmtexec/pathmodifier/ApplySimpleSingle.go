package pathmodifier

import (
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errwrappers"
	"gitlab.com/auk-go/pathhelper/pathinsfmt"
)

func ApplySimpleSingle(
	modifier *pathinsfmt.PathModifier,
	location string,
) *errorwrapper.Wrapper {
	errCollection := errwrappers.Empty()

	locations := []string{
		location,
	}

	isSuccess := ApplySimple(
		false,
		errCollection,
		modifier,
		locations)

	if !isSuccess {
		return errCollection.GetAsErrorWrapperPtr()
	}

	return nil
}
