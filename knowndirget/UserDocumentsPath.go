package knowndirget

import (
	"gitlab.com/auk-go/pathhelper/knowndir"
)

// Returns documents directory path as a string.
func UserDocumentsPath() string {
	return knowndir.Documents.CombineWith(UserPath())
}
