package pathhelper

import (
	"os"
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

func GetRawExecutableEnvironmentPathCollection() []string {
	pathString := os.Getenv(constants.Path)

	if IsWindows() {
		return strings.Split(pathString, constants.SemiColon)
	}

	return strings.Split(pathString, constants.Colon)
}
