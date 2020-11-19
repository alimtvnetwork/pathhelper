package pathhelper

import "gitlab.com/evatix-go/pathhelper/enums"

// Returns path to hosts on different platforms.
func GetHostFilePath() string {
	return enums.HostFile.CombineWith(GetEtcPath())
}
