package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/enums"
)

// Returns path to System32 path in windows system.
func GetSystem32() string {
	return enums.System32.CombineWith(GetWidowsDirectory())
}
