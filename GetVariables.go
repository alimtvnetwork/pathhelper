package pathhelper

import (
	"regexp"
	"strings"
)

func GetVariables(stringToCheck string) []string {
	var envVariableKeysForMap []string

	r, _ := regexp.Compile("\\$(\\w+)(\\d*)")

	if !r.MatchString(stringToCheck) {
		return nil
	}

	envVariableRawKeys := r.FindAllString(stringToCheck, -1)

	for _, arrayItem := range envVariableRawKeys {
		arrayItem = strings.Replace(arrayItem, "$", "", 1)

		envVariableKeysForMap = append(envVariableKeysForMap, arrayItem)
	}

	return envVariableKeysForMap
}
