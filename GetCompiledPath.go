package pathhelper

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/pathhelper/ispath"
)

func GetCompiledPath(
	pathTemplate string,
	compilingMap *map[string]string,
) string {
	if ispath.Empty(pathTemplate) {
		return pathTemplate
	}

	for key, value := range *compilingMap {
		pathTemplate = strings.Replace(pathTemplate, key, value, constants.MinusOne)
	}

	return pathTemplate
}
