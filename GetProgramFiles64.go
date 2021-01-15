package pathhelper

import (
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/knowndir"
)

// Returns Program Files directory on windows with OS architecture x64
func GetProgramFiles64() string {
	if !osconsts.IsWindows {
		return ""
	}

	return knowndir.ProgramFiles64.CombineWith(GetWindowsRoot())
}
