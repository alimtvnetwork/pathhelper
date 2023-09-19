package hashas

import (
	"strconv"
	"sync"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/errorwrapper/errdata/errstr"
	"gitlab.com/auk-go/errorwrapper/errtype"
)

func HexChecksumOfAnyItems(
	isSkipNil bool,
	method Variant,
	items ...interface{},
) *errstr.Results {
	if len(items) == 0 {
		return errstr.Empty.Results()
	}

	locker := sync.Mutex{}
	wg := &sync.WaitGroup{}
	var sliceErr []string
	checkSumSlice := make(
		[]string,
		len(items))

	hexChecksum := func(index int, source interface{}) bool {
		defer wg.Done()
		hexFileChecksumResult := method.HexSumOfAny(source)
		checkSumSlice[index] = hexFileChecksumResult.Value

		if hexFileChecksumResult.IsSuccess() {
			return true
		}

		// failed
		locker.Lock()
		defer locker.Unlock()

		message := "Failed Index : " +
			strconv.Itoa(index) +
			constants.Comma +
			coredynamic.TypeName(source) +
			constants.Comma +
			hexFileChecksumResult.
				ErrorWrapper.
				String()

		sliceErr = append(
			sliceErr,
			message)

		return false
	}

	for i, item := range items {
		if isSkipNil && item == nil {
			continue
		}

		wg.Add(constants.One)
		go hexChecksum(i, item)
	}

	wg.Wait()

	err := errcore.SliceToError(
		sliceErr)

	if err == nil {
		return errstr.New.Results.Strings(
			checkSumSlice)
	}

	// Failed
	return errstr.New.Results.Error(
		errtype.CheckSumCorrupted,
		err)
}
