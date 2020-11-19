package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/enums"
)

// Returns documents directory path as a string.
func GetUserDocumentsPath() string {
	return enums.Documents.CombineWith(GetUserPath())
}
