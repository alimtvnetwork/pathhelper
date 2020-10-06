package pathhelper

import (
	"runtime"

	"gitlab.com/evatix-go/pathhelper/constants"
)

func IsWindows() bool {
	return runtime.GOOS == constants.WindowsOS
}
