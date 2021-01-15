package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/knowndir"
)

// Returns .ssh path as string.
func GetSSHGlobal() string {
	return knowndir.SSHGlobal.CombineWith(GetUserPath())
}
