package pathhelper

import "strings"

// TODO : Discuss this for inside the function
func RemoveDoubleSeparator(path, separator string) string {
	if strings.Contains(path, DoubleDash) {
		for {
			path = strings.ReplaceAll(path, DoubleDash, separator)
			if !strings.Contains(path, DoubleDash) {
				break
			}
		}
	}

	if strings.Contains(path, DoubleUnderscore) {
		for {
			path = strings.ReplaceAll(path, DoubleUnderscore, separator)
			if !strings.Contains(path, DoubleUnderscore) {
				break
			}
		}
	}

	return path
}
