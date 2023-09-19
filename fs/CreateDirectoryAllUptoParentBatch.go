package fs

import (
	"gitlab.com/auk-go/errorwrapper"
)

func CreateDirectoryAllUptoParentMany(paths ...string) *errorwrapper.Wrapper {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	for _, path := range paths {
		if errWrap := CreateDirectoryAllUptoParent(path); errWrap.HasError() {
			return errWrap
		}
	}

	return nil
}
