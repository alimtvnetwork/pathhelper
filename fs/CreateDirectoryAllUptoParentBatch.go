package fs

import (
	"gitlab.com/evatix-go/errorwrapper"
)

func CreateDirectoryAllUptoParentMany(paths ...string) *errorwrapper.Wrapper {
	readWriteMutex.Lock()
	defer readWriteMutex.Unlock()

	for _, path := range paths {
		if errW := CreateDirectoryAllUptoParent(path); errW.HasError() {
			return errW
		}
	}

	return nil
}
