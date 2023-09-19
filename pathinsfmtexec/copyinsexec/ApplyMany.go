package copyinsexec

import (
	"gitlab.com/auk-go/errorwrapper/errwrappers"
	"gitlab.com/auk-go/pathhelper/pathinsfmt"
)

func ApplyMany(
	errCollection errwrappers.Collection,
	copyPaths *pathinsfmt.CopyPaths,
) (isSuccess bool) {
	if copyPaths.IsEmpty() {
		return true
	}

	isImmediateExit := !copyPaths.IsContinueOnError
	stateTracker := errCollection.StateTracker()

	for _, copyPath := range copyPaths.CopyPaths {
		err := Apply(&copyPath)

		errCollection.AddWrapperPtr(err)

		if isImmediateExit && err.HasError() {
			return false
		}
	}

	return stateTracker.IsSuccess()
}
