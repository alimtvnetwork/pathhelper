package hexchecksum

import "gitlab.com/evatix-go/pathhelper/hashas"

type FilesChecksumRequest struct {
	Method                     hashas.Variant
	IsGenerateContentsChecksum bool
	Files                      []string
}
