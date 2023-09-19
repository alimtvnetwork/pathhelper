package knowndirget

import "gitlab.com/auk-go/pathhelper/knowndir"

// Returns path to hosts on different platforms.
func HostFilePath() string {
	return knowndir.HostFile.CombineWith(EtcPath())
}
