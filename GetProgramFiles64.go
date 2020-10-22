package pathhelper

import "gitlab.com/evatix-go/pathhelper/enums"

// Returns Program Files directory on windows with OS architecture x64
func GetProgramFiles64() string {
	return enums.ProgramFiles64.GetPrefixCombinedWith(GetWindowsRoot())
}
