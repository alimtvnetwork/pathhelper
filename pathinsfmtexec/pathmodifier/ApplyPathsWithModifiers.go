package pathmodifier

import (
	"gitlab.com/auk-go/errorwrapper/errwrappers"
	"gitlab.com/auk-go/pathhelper/pathinsfmt"
)

func ApplyPathsWithModifiers(
	isContinueOnErr bool,
	errCollection *errwrappers.Collection,
	pathsWithModifiers *pathinsfmt.BasePathsWithModifiers,
) (isSuccess bool) {
	if pathsWithModifiers.IsPathWithModifiersUndefined() {
		return true
	}

	isImmediateExit := !isContinueOnErr
	stateTracker := errCollection.StateTracker()
	pathsWithModifiersRanges := pathsWithModifiers.PathWithModifiers

	for i := range pathsWithModifiersRanges {
		isSuccessInner := ApplyPathWithModifier(errCollection,
			&pathsWithModifiersRanges[i])

		if isImmediateExit && !isSuccessInner {
			return false
		}
	}

	return stateTracker.IsSuccess()
}
