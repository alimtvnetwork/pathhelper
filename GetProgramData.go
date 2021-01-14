package pathhelper

import (
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/knowndir"
)

func GetProgramData() string {
	if !osconsts.IsWindows {
		return ""
	}

	return knowndir.ProgramData.CombineWith(GetWindowsRoot())
}
