package pathhelper

import "gitlab.com/evatix-go/pathhelper/enums"

// Returns path to Fonts directory on different platforms.
func GetFontsPath() string {
	if IsWindows() {
		return enums.Fonts.GetPrefixCombinedWith(GetWidowsDirectory())
	}

	return enums.FontsUnix.Value()
}
