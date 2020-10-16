package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/enums"
)

// Returns path to etc directory on different platforms.
func GetEtcPath() string {
	if IsWindows() {
		return GetCombinePathWith(GetWidowsDirectory(), GetSystem32(), enums.Etc.Value())
	}

	return GetCombinePathWith(enums.Etc.Value())
}
