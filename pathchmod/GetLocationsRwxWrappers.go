package pathchmod

import (
	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
)

func GetLocationsRwxWrappers(
	isContinueOnError bool,
	locations []string,
) (
	filePathToRwxWrapper map[string]*chmodhelper.RwxWrapper,
	errWrap *errorwrapper.Wrapper,
) {
	resultMap, err := chmodhelper.GetExistingChmodRwxWrappers(
		isContinueOnError, locations...)

	if err != nil {
		return resultMap, errnew.Error.Type(
			errtype.ExistingChmodReadFailed,
			err,
		)
	}

	return resultMap, nil
}
