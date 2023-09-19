package pathfuncs

import (
	"fmt"
	"strings"
	"sync"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/errorwrapper/errdata/errstr"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
	"gitlab.com/auk-go/pathhelper/internal/consts"
)

func SimpleFilterFullPathsAsync(
	isContinueOnError bool,
	filter SimpleFilter,
	fullPaths ...string,
) *errstr.Results {
	length := len(fullPaths)

	if filter == nil || length == 0 {
		return errstr.Empty.Results()
	}

	if length <= consts.NonAsyncSafeRange {
		return SimpleFilterFullPaths(
			isContinueOnError,
			filter,
			fullPaths...)
	}

	wg := sync.WaitGroup{}
	firstFoundItems := make([]*string, length)
	locker := sync.Mutex{}
	isErrorFound := false
	var errSlice []string
	foundItems := 0

	adderFunc := func(index int, fullPath string) {
		defer wg.Done()

		isTake, err := filter(fullPath)

		if isTake {
			firstFoundItems[index] = &fullPath
			foundItems++
		}

		if err != nil {
			locker.Lock()
			isErrorFound = true
			errString := err.Error()

			if strings.Contains(errString, fullPath) {
				errSlice = append(
					errSlice,
					errString)
			} else {
				errSlice = append(
					errSlice,
					fmt.Sprintf(
						constants.MessageWrapMessageFormat,
						errString,
						fullPath))
			}

			locker.Unlock()
		}
	}

	for i, fullPath := range fullPaths {
		if !isContinueOnError && isErrorFound {
			break
		}

		wg.Add(1)
		go adderFunc(i, fullPath)
	}

	wg.Wait()

	finalItems := make([]string, 0, foundItems+constants.Capacity4)

	for i, item := range firstFoundItems {
		if item == nil {
			continue
		}

		finalItems = append(
			finalItems,
			*item)

		firstFoundItems[i] = nil
	}

	err := errcore.SliceToError(
		errSlice)

	if err != nil {
		errWrap := errnew.Type.Error(
			errtype.PathIssue,
			err,
		)

		return errstr.New.Results.Create(
			errWrap,
			finalItems)
	}

	firstFoundItems = nil

	return errstr.New.Results.Strings(finalItems)
}
