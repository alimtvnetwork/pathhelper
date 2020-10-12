package pathhelper

import "gitlab.com/evatix-go/pathhelper/enums"

func IsBothCase(os enums.OperatingSystem) bool {
	return IsWindowsCase(os) || IsUnixCase(os)
}
