package knowndirget

import (
	"gitlab.com/auk-go/core/osconsts"

	"gitlab.com/auk-go/pathhelper/internal/ispathinternal"
	"gitlab.com/auk-go/pathhelper/knowndir"
)

// Returns path to .git. Checks for it on all possible locations. If .git doesn't exist creates it.
// todo should we  use createDirectory to make .git . is it a directory or a file
func GitGlobal() string {
	homePath := UserPath()
	var outputPath, outputPathAlternate string

	if osconsts.IsWindows {
		outputPath = knowndir.GitGlobalWin.CombineWith(homePath)
	} else {
		outputPath = knowndir.GitGlobalUnix.CombineWith(homePath)
		outputPathAlternate = knowndir.GitGlobalUnixXdg.Value()
	}

	if !ispathinternal.Exists(outputPath) {
		outputPath = outputPathAlternate
	}

	return outputPath
}
