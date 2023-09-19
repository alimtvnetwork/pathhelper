package main

import (
	"fmt"

	"gitlab.com/auk-go/pathhelper/hashas"
	"gitlab.com/auk-go/pathhelper/hexchecksum"
)

func checksumTest02() {
	rs := hexchecksum.OfFilesRequest(&hexchecksum.FilesRequest{
		Method:                     hashas.DefaultFastHashMethod,
		IsGenerateContentsChecksum: true,
		IsGenerateFileListChecksum: true,
		Files:                      pkgRootFiles(),
	})

	fmt.Println(rs.JsonString())
}
