package pathinsfmtexec

import (
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
	"gitlab.com/evatix-go/pathhelper/symlink"
)

func ApplySymbolicLinksUsingErrorCollection(
	errorCollection *errwrappers.Collection,
	symLinks *pathinsfmt.SymbolicLinks,
) (isSuccess bool) {
	if symLinks == nil || symLinks.IsEmpty() {
		return true
	}

	errCount := errorCollection.Length()
	if symLinks.IsContinueOnError {
		for _, symLink := range symLinks.SymbolicLinks {
			errW := symlink.CreateUsingSymbolicLink(&symLink)

			errorCollection.AddWrapperPtr(errW)
		}

		return errorCollection.Length() == errCount
	}

	// immediate exit
	for _, symLink := range symLinks.SymbolicLinks {
		errW := symlink.CreateUsingSymbolicLink(&symLink)

		if errW.HasError() {
			errorCollection.AddWrapperPtr(errW)

			return false
		}
	}

	return errorCollection.Length() == errCount
}
