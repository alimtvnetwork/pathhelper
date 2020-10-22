package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/enums"
)

// Returns path to etc directory on different platforms.
func GetEtcPath() string {
	if IsWindows() {
		return enums.Etc.GetPrefixCombinedWith(GetSystemDriversPath())
	}

	return enums.Etc.Value()
}
