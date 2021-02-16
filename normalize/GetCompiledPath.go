package normalize

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
)

func GetCompiledPath(
	pathTemplate string,
	compilingMap *map[string]string,
) string {
	if pathhelpercore.IsEmptyPath(pathTemplate) {
		return pathTemplate
	}

	for key, value := range *compilingMap {
		pathTemplate = strings.Replace(pathTemplate, key, value, constants.MinusOne)
	}

	return pathTemplate
}
