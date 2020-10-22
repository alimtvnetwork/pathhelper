package pathhelper

import (
	"os"
	"strings"

	"gitlab.com/evatix-go/pathhelper/constants"
)

// todo change filename accordingly. issue #31. https://gitlab.com/evatix-go/pathhelper/-/issues/31
func GetRawExecutableEnvironmentPathCollection() []string {
	var paths []string

	pathString := os.Getenv(constants.Path)
	paths = strings.Split(pathString, constants.SemiColon)

	return paths
}
