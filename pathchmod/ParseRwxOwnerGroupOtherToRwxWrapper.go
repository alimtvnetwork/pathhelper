package pathchmod

import (
	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/core/chmodhelper/chmodins"
	"gitlab.com/auk-go/core/codestack"
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
)

func ParseRwxOwnerGroupOtherToRwxWrapper(
	rwxOwnerGroupOther *chmodins.RwxOwnerGroupOther,
) (*chmodhelper.RwxWrapper, *errorwrapper.Wrapper) {
	if rwxOwnerGroupOther == nil {
		return nil, errnew.Null.WithMessage(
			"cannot process empty or nil pointer of chmodins.RwxOwnerGroupOther",
			rwxOwnerGroupOther)
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
