package pathhelper

import "gitlab.com/evatix-go/pathhelper/knowndir"

// Returns path to hosts on different platforms.
func GetHostFilePath() string {
	return knowndir.HostFile.CombineWith(GetEtcPath())
}
