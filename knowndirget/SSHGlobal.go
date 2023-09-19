package knowndirget

import (
	"gitlab.com/auk-go/pathhelper/knowndir"
)

// Returns .ssh path as string.
func SSHGlobal() string {
	return knowndir.SSHGlobal.CombineWith(UserPath())
}
