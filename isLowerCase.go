package pathhelper

import "gitlab.com/auk-go/core/constants"

func isLowerCase(eachChar rune) bool {
	return eachChar >= constants.LowerCaseA && eachChar <= constants.LowerCaseZ
}
