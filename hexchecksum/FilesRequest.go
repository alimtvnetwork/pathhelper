package hexchecksum

import "gitlab.com/evatix-go/pathhelper/hashas"

type FilesRequest struct {
	Method                     hashas.Variant
	IsGenerateContentsChecksum bool
	Files                      []string
}
