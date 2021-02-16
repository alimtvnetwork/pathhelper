package knowndirget

import (
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/knowndir"
)

// Returns path to SysWOW64 path in windows system.
func GetSystem64() string {
	return pathhelper.GetCombinePathsWith(WidowsDirectory(), knowndir.System64.Value())
}
