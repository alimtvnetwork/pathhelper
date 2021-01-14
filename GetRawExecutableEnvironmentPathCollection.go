package pathhelper

import (
	"os"
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/osconsts"
)

func GetRawExecutableEnvironmentPathCollection() []string {
	pathString := os.Getenv(constants.Path)

	if osconsts.IsWindows {
		return strings.Split(pathString, constants.SemiColon)
	}

	return strings.Split(pathString, constants.Colon)
}
