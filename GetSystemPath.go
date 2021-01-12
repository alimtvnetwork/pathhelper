package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/enums"
)

// Returns path to System directory on different platforms.
func GetSystemPath() string {
	if IsWindows() {
		return GetWidowsDirectory()
	}

	return enums.SystemUnix.Value()
}
