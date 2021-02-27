package expandpath

import (
	"os"

	"gitlab.com/evatix-go/core/constants"
)

// expandEnvironmentVariable function takes an array of environment variables (string) as input
// and outputs a map of expanded path of those variables if the paths exist.
func expandEnvironmentVariable(variableForExpanding *[]string) *map[string]string {
	var expandedPath = map[string]string{}

	for _, keyName := range *variableForExpanding {
		_, exists := os.LookupEnv(keyName)

		envVariableKeyName := constants.Dollar + keyName

		if exists {
			expandedPath[envVariableKeyName] = os.Getenv(keyName)
		} else {
			expandedPath[envVariableKeyName] = envVariableKeyName
		}
	}

	return &expandedPath
}
