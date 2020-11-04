package pathhelper

import "gitlab.com/evatix-go/pathhelper/enums"

// Returns path to System drivers directory on different platforms.
func GetSystemDriversPath() string {
	if IsWindows() {
		return GetCombinePathsWith(GetSystem32(), enums.Drivers.Value())
	}

	return enums.DriversUnix.Value()
}
