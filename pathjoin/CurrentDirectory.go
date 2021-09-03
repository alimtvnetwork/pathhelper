package pathjoin

import (
	"path/filepath"
	"runtime"
)

func CurrentDirectory() string {
	_, b, _, _ := runtime.Caller(1)

	return filepath.Dir(b)
}
