package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/enums"
)

// "/etc/apache/"
func GetApacheLinuxPath() string {
	if !IsUnix() {
		panic("Path only available for Unix OS")
	}

	return enums.ApacheLinuxPath.Value()
}
