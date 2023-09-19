package pathinsfmtexec

import (
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errwrappers"
	"gitlab.com/auk-go/pathhelper/pathinsfmt"
	"gitlab.com/auk-go/pathhelper/symlink"
)

func ApplySymbolicLinks(symLinks *pathinsfmt.SymbolicLinks) *errorwrapper.Wrapper {
	if symLinks == nil || symLinks.IsEmpty() {
		return nil
	}

	if symLinks.IsContinueOnError {
		errCollection := errwrappers.Empty()

		for _, symLink := range symLinks.SymbolicLinks {
			errWrap := symlink.CreateUsingSymbolicLink(&symLink)

			errCollection.AddWrapperPtr(errWrap)
		}

		return errCollection.GetAsErrorWrapperPtr()
	}

	for _, symLink := range symLinks.SymbolicLinks {
		errWrap := symlink.CreateUsingSymbolicLink(&symLink)
		if errWrap.HasError() {
			return errWrap
		}
	}

	return nil
}
