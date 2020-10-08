package pathhelper

import (
	"os"
	"strings"
)

func expandEnvironmentVariable(variableForExapanding []string) map[string]string {
	var expandedPath = map[string]string{}

	for _, keyName := range variableForExapanding {
		_, exists := os.LookupEnv(keyName)

		if !exists {
			expandedPath["$"+keyName] = ""
		}

		if exists {
			expandedPath["$"+keyName] = os.Getenv(strings.ToLower(keyName))
		}
	}

	return expandedPath
}
