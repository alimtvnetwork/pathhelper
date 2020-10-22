package pathhelper

import "gitlab.com/evatix-go/pathhelper/enums"

// Returns path to SysWOW64 path in windows system.
func GetSystem64() string {
	return GetCombinePathWith(GetWidowsDirectory(), enums.System64.Value())
}
