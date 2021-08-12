package deletepaths

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func SingleOrRecursiveOnCondition(
	condition Condition,
	location string,
) *errorwrapper.Wrapper {
	if !condition.IsRemove {
		return errnew.EmptyPtr
	}

	if condition.IsRecursive && condition.IsExistBeforeClear {
		return RecursiveOnExist(location)
	} else if condition.IsRecursive && !condition.IsExistBeforeClear {
		return Recursive(location)
	}

	if !condition.IsRecursive && condition.IsExistBeforeClear {
		return SingleOnExist(location)
	} else if condition.IsRecursive && !condition.IsExistBeforeClear {
		return Single(location)
	}

	return errnew.Ref1Msg(
		errtype.InvalidOption,
		"None of the condition satisfied for path remove using condition!",
		"condition",
		condition)
}
