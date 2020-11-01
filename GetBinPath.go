package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/enums"
)

// Returns path to bin directory as a string.
func GetBinPath() string {
	if !IsWindows() {
		return enums.BinUnix.Value()
	}

	binPath := enums.Bin.GetPrefixCombinedWith(GetUserPath())

	return binPath
}
