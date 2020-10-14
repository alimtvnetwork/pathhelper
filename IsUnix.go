package pathhelper

import (
	"runtime"

	"gitlab.com/evatix-go/pathhelper/constants"
)

func IsUnix() bool {
	return runtime.GOOS != constants.WindowsOS
}
