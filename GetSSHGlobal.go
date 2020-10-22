package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/enums"
)

// Returns .ssh path as string.
func GetSSHGlobal() string {
	return enums.SSHGlobal.GetPrefixCombinedWith(GetUserPath())
}
