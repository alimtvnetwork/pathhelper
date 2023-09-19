package pathfuncs

import (
	"gitlab.com/auk-go/core/coredata/stringslice"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/errorwrapper/errdata/errstr"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
)

func SimpleFilterFullPaths(
	isContinueOnError bool,
	filter SimpleFilter,
	fullPaths ...string,
) *errstr.Results {
	length := len(fullPaths)

	if filter == nil || length == 0 {
		return errstr.Empty.Results()
	}

	foundItems := stringslice.MakeDefault(length)
	var errSlice []string

	for _, fullPath := range fullPaths {
		isTake, err := filter(fullPath)

		if isTake {
			foundItems = append(foundItems, fullPath)
		}

		if err != nil {
			message := errtype.PathIssue.ReferencesCsv(
				err.Error(),
				fullPath)

			errSlice = append(
				errSlice,
				message)
		}

		if !isContinueOnError && err != nil {
			break
		}
	}

	err := errcore.SliceToError(errSlice)

	if err != nil {
		errWrap := errnew.
			Path.
			Error(
				errtype.PathIssue,
				err,
				"")

		return errstr.New.Results.Create(
			errWrap,
			foundItems)
	}

	return errstr.New.Results.Strings(
		foundItems)
}
