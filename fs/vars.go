package fs

import "sync"

var (
	writerMutex = &sync.Mutex{}
)
