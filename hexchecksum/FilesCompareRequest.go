package hexchecksum

import "gitlab.com/auk-go/pathhelper/hashas"

type FilesCompareRequest struct {
	Method     hashas.Variant
	LeftFiles  []string
	RightFiles []string
}
