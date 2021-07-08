package expandpath

import (
	"gitlab.com/evatix-go/core/constants"
)

func getDollarOrPercentSymbolIdentifierEnvInfos(
	stringToCheck string,
) ([]EnvKeyInfo, string) {
	envVariableRawKeys := regexEachWordDollar.FindAllString(
		stringToCheck,
		constants.MinusOne)

	var envInfos []EnvKeyInfo

	if len(envVariableRawKeys) > 0 {
		envInfos = curlyBraceRemoveFromKeys(envVariableRawKeys)
	}

	envVariableRawKeys2 := regexEachWordPercent.FindAllString(
		stringToCheck, constants.MinusOne)

	if len(envVariableRawKeys2) > 0 {
		envInfos2 := curlyBraceRemoveFromKeys(envVariableRawKeys2)

		for _, envInfo := range envInfos2 {
			envInfos = append(
				envInfos,
				envInfo)
		}
	}

	return envInfos, constants.Percent
}
