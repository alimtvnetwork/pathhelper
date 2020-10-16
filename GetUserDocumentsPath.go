package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/enums"
)

// Returns documents directory path. If directory doesn't exist, then function creates the directory and returns the path as a string.
func GetUserDocumentsPath() string {
	return GetUserPathOf(enums.Documents.Value())
}
