package pathhelper

import "gitlab.com/evatix-go/pathhelper/constants"

func getEnvironmentVarRawKeysAndRegularExpressionSymbol(stringToCheck string) ([]string, string) {
	var regularExpressionSymbol string
	var envVariableRawKeys []string

	if regularExpressionForEachWordsWithDollarSymbol.MatchString(stringToCheck) {
		envVariableRawKeys = regularExpressionForEachWordsWithDollarSymbol.FindAllString(stringToCheck, constants.MinusOne)
		regularExpressionSymbol = constants.Dollar
	}

	envVariableRawKeys = regularExpressionForEachWordsWithinPercentSymbol.FindAllString(stringToCheck, constants.MinusOne)
	regularExpressionSymbol = constants.Percent

	return envVariableRawKeys, regularExpressionSymbol
}
