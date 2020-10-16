package pathhelper

import "gitlab.com/evatix-go/pathhelper/enums"

// Returns Program Files directory on windows with OS architecture x32
func GetProgramFiles32() string {
	return GetCombinePathWith(GetWindowsRoot(), enums.ProgramFiles32.Value())
}
