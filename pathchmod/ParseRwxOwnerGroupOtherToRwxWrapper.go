package pathchmod

import (
	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
	"gitlab.com/evatix-go/core/codestack"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func ParseRwxOwnerGroupOtherToRwxWrapper(
	rwxOwnerGroupOther *chmodins.RwxOwnerGroupOther,
) (*chmodhelper.RwxWrapper, *errorwrapper.Wrapper) {
	if rwxOwnerGroupOther == nil {
		return nil, errnew.Messages.Many(
			errtype.EmptyPointerOrNullPointer,
			"cannot process empty or nil pointer of chmodins.RwxOwnerGroupOther")
	}

	varWrapper, err := chmodhelper.ParseRwxOwnerGroupOtherToRwxVariableWrapper(
		rwxOwnerGroupOther)

	if err != nil {
		return nil, errorwrapper.NewRef(
			codestack.SkipNone,
			errtype.ChmodInvalid,
			err,
			"RwxOwnerGroupOther",
			rwxOwnerGroupOther.String())
	}

	return varWrapper.ToCompileFixedPtr(), nil
}
