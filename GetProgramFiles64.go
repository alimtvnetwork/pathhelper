package pathhelper

import "gitlab.com/evatix-go/pathhelper/enums"

// Returns Program Files directory on windows with OS architecture x64
func GetProgramFiles64() string {
	if !IsWindows() {
		return ""
	}

	return enums.ProgramFiles64.CombineWith(GetWindowsRoot())
}
