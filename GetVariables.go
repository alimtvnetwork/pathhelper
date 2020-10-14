package pathhelper

import (
	"regexp"
	"strings"

	"gitlab.com/evatix-go/pathhelper/constants"
)

var regularExpressionForEachWordsWithDollarSymbol, _ = regexp.Compile(constants.RegExForEachWordsWithDollarSymbol)

// getVariables function takes a string input and identifies every word that begins with "$" in that input
// string then returns an array of those words. If input has no word starting with "$" then returns nil.
func GetVariables(stringToCheck string) []string {
	var envVariableKeysForMap []string

	if !regularExpressionForEachWordsWithDollarSymbol.MatchString(stringToCheck) {
		return nil
	}

	envVariableRawKeys := regularExpressionForEachWordsWithDollarSymbol.FindAllString(stringToCheck, constants.MinusOne)

	for _, rawEnvKeys := range envVariableRawKeys {
		rawEnvKeys = strings.Replace(rawEnvKeys, constants.Dollar, "", constants.One)

		envVariableKeysForMap = append(envVariableKeysForMap, rawEnvKeys)
	}

	return envVariableKeysForMap
}
