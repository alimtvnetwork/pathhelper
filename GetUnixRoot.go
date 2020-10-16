package pathhelper

import "gitlab.com/evatix-go/pathhelper/enums"

// Returns unix system root as a string
func GetUnixRoot() string {
	return enums.UnixRoot.Value()
}
