package pathhelper

import (
	"runtime"

	"gitlab.com/evatix-go/core/constants"
)

func IsWindows() bool {
	return runtime.GOOS == constants.WindowsOS
}
