package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/constants"
	"gitlab.com/evatix-go/pathhelper/enums"
)

// Returns path to bin directory. If directory doesn't exist in windows it creates the directory and returns the path as a string.
func GetBinPath() string {
	if !IsWindows() {
		return enums.BinUnix.Value()
	}

	binPath := GetCombinePathWith(GetUserPath(), enums.Bin.Value())
	CreateDirectory(binPath, constants.Perm)

	return binPath
}
