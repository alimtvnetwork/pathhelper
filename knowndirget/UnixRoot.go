package knowndirget

import "gitlab.com/auk-go/pathhelper/knowndir"

// Returns unix system root as a string
func UnixRoot() string {
	return knowndir.UnixRoot.Value()
}
