package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/enums"
)

// Returns path to .git. Checks for it on all possible locations. If .git doesn't exist creates it.
// todo should we  use createDirectory to make .git . is it a directory or a file
func GetGitGlobal() string {
	homePath := GetUserPath()
	var outputPath, outputPathAlternate string

	if IsWindows() {
		outputPath = enums.GitGlobalWin.CombineWith(homePath)
	} else {
		outputPath = enums.GitGlobalUnix.CombineWith(homePath)
		outputPathAlternate = enums.GitGlobalUnixXdg.Value()
	}

	if !IsPathExist(outputPath) {
		outputPath = outputPathAlternate
	}

	return outputPath
}
