package pathhelper

import (
	"regexp"
	"strings"

	"gitlab.com/evatix-go/pathhelper/constants"
)

var r, _ = regexp.Compile(constants.RegExForEnvVar)

func getVariables(stringToCheck string) []string {
	var envVariableKeysForMap []string

	if !r.MatchString(stringToCheck) {
		return nil
	}

	envVariableRawKeys := r.FindAllString(stringToCheck, constants.MinusOne)

	for _, arrayItem := range envVariableRawKeys {
		arrayItem = strings.Replace(arrayItem, constants.Dollar, "", constants.One)

		envVariableKeysForMap = append(envVariableKeysForMap, arrayItem)
	}

	return envVariableKeysForMap
}
