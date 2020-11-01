package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/enums"
)

// "/etc/apache/"
func GetApacheLinuxPath() string {
	return enums.ApacheLinuxPath.Value()
}
