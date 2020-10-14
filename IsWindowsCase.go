package pathhelper

import "gitlab.com/evatix-go/pathhelper/enums"

func IsWindowsCase(os enums.OperatingSystem) bool {
	return os == enums.Windows
}
