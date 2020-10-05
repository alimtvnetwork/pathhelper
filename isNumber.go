package pathhelper

import "unicode"

func isNumber(path string) bool {
	for i, r := range path {
		isMinus := i == 0 && r == '-'
		isValid := isMinus || unicode.IsNumber(r)

		if !isValid {
			return false
		}
	}

	return true
}
