package pathhelper

import "gitlab.com/evatix-go/pathhelper/enums"

// Returns path to Services directory on different platforms.
func GetServicesPath() string {
	if IsWindows() {
		return enums.Services.GetPrefixCombinedWith(GetEtcPath())
	}

	return GetSystemPath()
}
