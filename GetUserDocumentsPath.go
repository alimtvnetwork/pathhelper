package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/knowndir"
)

// Returns documents directory path as a string.
func GetUserDocumentsPath() string {
	return knowndir.Documents.CombineWith(GetUserPath())
}
