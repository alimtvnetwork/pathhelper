package pathhelper

import "gitlab.com/evatix-go/pathhelper/enums"

// Returns path to System drivers directory on different platforms.
func GetSystemDriversPath() string {
	if IsWindows() {
		return enums.Drivers.CombineWith(GetSystem32())
	}

	return enums.DriversUnix.Value()
}
