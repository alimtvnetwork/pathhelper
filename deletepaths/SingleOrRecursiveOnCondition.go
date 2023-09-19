package deletepaths

import (
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
)

func SingleOrRecursiveOnCondition(
	condition Condition,
	location string,
) *errorwrapper.Wrapper {
	if !condition.IsRemove {
		return nil
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

	return errnew.Ref.Messages(
		errtype.InvalidOption,
		"condition",
		condition,
		"None of the condition satisfied for path remove using condition!",
	)
}
