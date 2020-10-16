package pathhelper

import "gitlab.com/evatix-go/pathhelper/enums"

func GetProgramData() string {
	return GetCombinePathWith(GetWindowsRoot(), enums.ProgramData.Value())
}
