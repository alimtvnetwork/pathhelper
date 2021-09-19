package hexchecksum

import (
	"sync"

	"gitlab.com/evatix-go/core/msgtype"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/hashas"
)

func OfFilesContentsAsync(
	hashMethod hashas.Variant,
	filesPaths ...string,
) *errstr.Result {
	if len(filesPaths) == 0 {
		return errstr.Empty()
	}

	locker := sync.Mutex{}
	wg := &sync.WaitGroup{}
	var sliceErr []string
	checkSumSlice := make(
		[]string,
		len(filesPaths))

	hexChecksum := func(index int, source string) bool {
		defer wg.Done()
		hexFileChecksumResult := hashMethod.
			HexSumOfFile(source)

		if hexFileChecksumResult.IsSuccess() {
			checkSumSlice[index] = hexFileChecksumResult.Value

			return true
		}

		// failed
		locker.Lock()
		sliceErr = append(
			sliceErr,
			hexFileChecksumResult.
				ErrorWrapper.
				String())
		locker.Unlock()

		return false
	}

	wg.Add(len(filesPaths))
	for i, filePath := range filesPaths {
		go hexChecksum(i, filePath)
	}

	wg.Wait()

	err := msgtype.SliceToError(
		sliceErr)

	if err == nil {
		// success
		return hashMethod.HexSumOfAny(
			checkSumSlice)
	}

	return errstr.Error(
		errtype.CheckSumCorrupted,
		err)
}
