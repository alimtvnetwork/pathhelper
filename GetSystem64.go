package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/knowndir"
)

// Returns path to SysWOW64 path in windows system.
func GetSystem64() string {
	return GetCombinePathsWith(GetWidowsDirectory(), knowndir.System64.Value())
}
