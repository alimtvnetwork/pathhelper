package pathhelper

import (
	"os"

	"gitlab.com/evatix-go/pathhelper/constants"
)

// expandEnvironmentVariable function takes an array of environment variables (string) as input
// and outputs a map of expanded path of those variables if the paths exist.
func expandEnvironmentVariable(variableForExapanding []string) map[string]string {
	var expandedPath = map[string]string{}

	for _, keyName := range variableForExapanding {
		_, exists := os.LookupEnv(keyName)

		var envVariableKeyName string = constants.Dollar + keyName

		if exists {
			expandedPath[envVariableKeyName] = os.Getenv(keyName)
		} else {
			expandedPath[envVariableKeyName] = envVariableKeyName
		}
	}

	return expandedPath
}
