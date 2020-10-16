package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/enums"
)

// Returns .ssh. If doesn't exist then creates and returns the path as string.
func GetSSHGlobal() string {
	return GetUserPathOf(enums.SSHGlobal.Value())
}
