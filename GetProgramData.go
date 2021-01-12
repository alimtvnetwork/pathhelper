package pathhelper

import "gitlab.com/evatix-go/pathhelper/enums"

func GetProgramData() string {
	if !IsWindows() {
		return ""
	}

	return enums.ProgramData.CombineWith(GetWindowsRoot())
}
