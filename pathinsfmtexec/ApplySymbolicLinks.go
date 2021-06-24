package pathinsfmtexec

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
	"gitlab.com/evatix-go/pathhelper/symlink"
)

func ApplySymbolicLinks(symLinks *pathinsfmt.SymbolicLinks) *errorwrapper.Wrapper {
	if symLinks == nil || symLinks.IsEmpty() {
		return errnew.EmptyPtr
	}

	if symLinks.IsContinueOnError {
		errCollection := errwrappers.Empty()

		for _, symLink := range symLinks.SymbolicLinks {
			errW := symlink.CreateUsingSymbolicLink(&symLink)

			errCollection.AddWrapperPtr(errW)
		}

		return errCollection.GetAsErrorWrapperPtr()
	}

	for _, symLink := range symLinks.SymbolicLinks {
		errW := symlink.CreateUsingSymbolicLink(&symLink)
		if errW.HasError() {
			return errW
		}
	}

	return errnew.EmptyPtr
}
