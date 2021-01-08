package pathhelper

import (
	"regexp"
	"strings"

	"gitlab.com/evatix-go/pathhelper/constants"
)

var (
	regularExpressionForEachWordsWithDollarSymbol, _    = regexp.Compile(constants.RegExForEachWordsWithDollarSymbol)
	regularExpressionForEachWordsWithinPercentSymbol, _ = regexp.Compile(constants.RegExForEachWordsWithinPercentSymbol)
)

// getVariables function takes a string input and identifies every word that begins with "$" or every word within two "%"
// in that input string then returns an array of those words. If input is empty or has no such word then returns nil.
func GetVariables(stringToCheck string) []string {
	var regularExpressionSymbol string
	var envVariableKeysForMap, envVariableRawKeys []string

	if len(stringToCheck) == 0 {
		return nil
	}

	// Check which regular expression case is true
	isNotRegularExpressionCase := !regularExpressionForEachWordsWithDollarSymbol.MatchString(stringToCheck) &&
		!regularExpressionForEachWordsWithinPercentSymbol.MatchString(stringToCheck)

	if isNotRegularExpressionCase {
		return nil
	}

	envVariableRawKeys, regularExpressionSymbol = getEnvironmentVarRawKeysAndRegularExpressionSymbol(stringToCheck)

	for _, rawEnvKeys := range envVariableRawKeys {
		rawEnvKeys = strings.Replace(rawEnvKeys, regularExpressionSymbol, constants.EmptyString, constants.MinusOne)

		envVariableKeysForMap = append(envVariableKeysForMap, rawEnvKeys)
	}

	return envVariableKeysForMap
}
