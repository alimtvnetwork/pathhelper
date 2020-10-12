package pathhelper

import "gitlab.com/evatix-go/pathhelper/enums"

// it returns if not windows
func IsUnixCase(os enums.OperatingSystem) bool {
	// return os != enums.Windows && !IsWindows()
	return !IsWindowsCase(os)
}
