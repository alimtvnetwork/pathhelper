package fs

import "os"

func IsPathExistsUsingLock(location string) bool {
	readWriteMutex.Lock()
	defer readWriteMutex.Unlock()

	_, err := os.Stat(location)

	return !os.IsNotExist(err)
}
