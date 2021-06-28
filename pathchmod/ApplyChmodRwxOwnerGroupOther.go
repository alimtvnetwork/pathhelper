package pathchmod

import (
	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
)

func ApplyChmodRwxOwnerGroupOther(
	isRecursive,
	isSkipOnInvalid,
	isContinueOnError bool,
	rwxOwnerGroupOther *chmodins.RwxOwnerGroupOther,
	paths []string,
) *errorwrapper.Wrapper {
	if rwxOwnerGroupOther == nil ||
		len(paths) == 0 {
		return errnew.EmptyPtr
	}

	condition := &chmodins.Condition{
		IsSkipOnInvalid:   isSkipOnInvalid,
		IsContinueOnError: isContinueOnError,
		IsRecursive:       isRecursive,
	}

	fileMode, err := ParseRwxOwnerGroupOtherToFileMode(rwxOwnerGroupOther)

	if err.HasError() {
		return err
	}

	_, errWp := ApplyChmodOnFilesUsingCondition(
		condition,
		fileMode,
		paths...)

	return errWp
}
