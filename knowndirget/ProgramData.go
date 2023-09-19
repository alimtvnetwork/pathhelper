package knowndirget

import (
	"gitlab.com/auk-go/core/osconsts"

	"gitlab.com/auk-go/pathhelper/knowndir"
)

func ProgramData() string {
	if !osconsts.IsWindows {
		return ""
	}

	return knowndir.ProgramData.CombineWith(WindowsRoot())
}
