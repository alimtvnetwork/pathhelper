package checksummer

import (
	"gitlab.com/auk-go/errorwrapper/errdata/errbyte"

	"gitlab.com/auk-go/pathhelper/hashas"
)

func FileRaw(filePath string, v hashas.Variant) *errbyte.Results {
	return v.SumOfFile(filePath)
}
