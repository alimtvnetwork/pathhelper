package fs

import "sync"

var (
	readWriteMutex = &sync.Mutex{}
)
