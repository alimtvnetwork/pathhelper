package pathhelper

import "gitlab.com/evatix-go/pathhelper/enums"

func GetProgramData() string {
	return enums.ProgramData.GetPrefixCombinedWith(GetWindowsRoot())
}
