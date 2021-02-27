package expandpath

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

// getVariables function takes a string input and identifies every word
// that begins with "$" or every word within two "%"
// in that input string then returns an array of those words.
// If input is empty or has no such word then returns nil.
func GetEnvironmentVariables(
	pathContainsEnvVarStartingDollarSymbol string,
) *[]string {
	var regularExpressionSymbol string
	var envVariableKeysForMap, envVariableRawKeys []string

	if len(pathContainsEnvVarStartingDollarSymbol) == 0 {
		return nil
	}

	// Check which regular expression case is true
	isNotRegularExpressionCase :=
		!regularExpressionForEachWordsWithDollarSymbol.
			MatchString(pathContainsEnvVarStartingDollarSymbol) &&
			!regularExpressionForEachWordsWithinPercentSymbol.
				MatchString(pathContainsEnvVarStartingDollarSymbol)

	if isNotRegularExpressionCase {
		return nil
	}

	envVariableRawKeys, regularExpressionSymbol =
		getEnvironmentVarRawKeysAndRegularExpressionSymbol(
			pathContainsEnvVarStartingDollarSymbol)

	for _, rawEnvKeys := range envVariableRawKeys {
		rawEnvKeys = strings.Replace(
			rawEnvKeys,
			regularExpressionSymbol,
			constants.EmptyString,
			constants.MinusOne,
		)

		envVariableKeysForMap = append(
			envVariableKeysForMap,
			rawEnvKeys,
		)
	}

	return &envVariableKeysForMap
}
