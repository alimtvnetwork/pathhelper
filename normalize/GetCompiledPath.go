package normalize

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

func GetCompiledPath(
	pathTemplate string,
	compilingMap map[string]string,
) string {
	if pathTemplate == constants.EmptyString {
		return pathTemplate
	}

	if compilingMap == nil || len(compilingMap) == 0 {
		return pathTemplate
	}

	for key, value := range compilingMap {
		pathTemplate = strings.Replace(
			pathTemplate,
			key,
			value,
			constants.MinusOne)
	}

	return pathTemplate
}
