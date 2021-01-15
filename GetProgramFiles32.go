package pathhelper

import (
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/knowndir"
)

// Returns Program Files directory on windows with OS architecture x32
func GetProgramFiles32() string {
	if !osconsts.IsWindows {
		return ""
	}

	return knowndir.ProgramFiles32.CombineWith(GetWindowsRoot())
}
