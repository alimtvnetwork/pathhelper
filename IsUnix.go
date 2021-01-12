package pathhelper

import (
	"runtime"

	"gitlab.com/evatix-go/core/constants"
)

func IsUnix() bool {
	return runtime.GOOS != constants.WindowsOS
}
