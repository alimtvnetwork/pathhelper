package fs

import (
	"gitlab.com/auk-go/errorwrapper/errdata/errbyte"
	"gitlab.com/auk-go/pathhelper/hashas"
)

func CheckSumFileBytesUsingLock(
	hashType hashas.Variant,
	location string,
) *errbyte.Results {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	return CheckSumFileBytes(hashType, location)
}
