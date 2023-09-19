package main

import (
	"fmt"

	"gitlab.com/auk-go/pathhelper/hashas"
	"gitlab.com/auk-go/pathhelper/hexchecksum"
)

func checksumTest01() {
	rs := hexchecksum.DetailedResultOfRequestAsync(&hexchecksum.FilesRequest{
		Method:                     hashas.DefaultFastHashMethod,
		IsGenerateContentsChecksum: true,
		IsGenerateFileListChecksum: true,
		Files:                      pkgRootFiles(),
	})

	fmt.Println(rs.JsonString())
}
