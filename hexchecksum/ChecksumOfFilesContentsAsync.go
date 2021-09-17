package hexchecksum

import (
	"sync"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/msgtype"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/hashas"
)

func ChecksumOfFilesContentsAsync(
	hashMethod hashas.Variant,
	files []string,
) *errstr.Result {
	if len(files) == 0 {
		return errstr.Empty()
	}

	locker := sync.Mutex{}
	wg := &sync.WaitGroup{}
	var sliceErr []string
	checkSumSlice := make([]string,
		len(files))

	hexChecksum := func(index int, source string) bool {
		hexFileChecksumResult := hashMethod.HexSumOfFile(source)

		if hexFileChecksumResult.IsSuccess() {
			wg.Done()
			checkSumSlice[index] = hexFileChecksumResult.Value

			return true
		}

		// failed
		locker.Lock()
		defer locker.Unlock()
		sliceErr = append(
			sliceErr,
			hexFileChecksumResult.ErrorWrapper.String())
		wg.Done()

		return false
	}

	for i, filePath := range files {
		wg.Add(constants.One)
		go hexChecksum(i, filePath)
	}

	wg.Wait()

	err := msgtype.SliceToError(sliceErr)

	if err == nil {
		// success
		return hashMethod.HexSumOfAny(checkSumSlice)
	}

	return errstr.Error(
		errtype.CheckSumCorrupted,
		err)
}
