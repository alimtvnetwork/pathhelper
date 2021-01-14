package pathhelper

import "gitlab.com/evatix-go/pathhelper/knowndir"

// Returns unix system root as a string
func GetUnixRoot() string {
	return knowndir.UnixRoot.Value()
}
