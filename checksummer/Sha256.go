package checksummer

import (
	"gitlab.com/auk-go/errorwrapper/errdata/errbyte"

	"gitlab.com/auk-go/pathhelper/hashas"
)

func Sha256(filePath string) *errbyte.Results {
	return FileRaw(filePath, hashas.Sha256)
}
