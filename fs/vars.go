package fs

import "sync"

var (
	globalMutex = &sync.Mutex{}
)
