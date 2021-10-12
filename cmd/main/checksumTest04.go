package main

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreasync"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/filestate"
	"gitlab.com/evatix-go/pathhelper/hashas"
	"gitlab.com/evatix-go/pathhelper/hexchecksum"
)

func checksumTest04() {
	files := pkgRootFiles()
	var detailedResult *hexchecksum.DetailedResult
	var fileStateMappedInfoItems *filestate.MappedInfoItems
	errCollection := errwrappers.Empty()
	hashMethod := hashas.DefaultFastHashMethod

	requestSample := hexchecksum.FilesRequest{
		Method:                     hashMethod,
		IsSortFileNames:            false,
		IsSortFilesChecksum:        false,
		IsGenerateContentsChecksum: true,
		IsGenerateFileListChecksum: true,
		Files:                      files,
	}

	// requestSample.SortFileNames()
	// fmt.Println(strings.Join(requestSample.Files, constants.NewLineUnix))

	coreasync.Waited.ParallelVoidTasks(
		func() {
			detailedResult = hexchecksum.DetailedResultOfRequestAsync(&requestSample)
		},
		func() {
			fileStateMappedInfoItems, errCollection = filestate.NewMappedInfoItemsUsingFilePaths(
				hashMethod,
				false,
				files...)

			fmt.Println(errCollection.String())
		},
	)

	for keyFilePath, checkSum := range detailedResult.Hashmap.Items() {
		info := fileStateMappedInfoItems.GetInfoByFilePath(keyFilePath)

		if info.HexContentChecksum != checkSum {
			err := errcore.ExpectingSimpleNoType(
				info.FullPath+constants.SpaceHypheAngelBracketSpace+"doesn't match checksum",
				checkSum,
				info.HexContentChecksum+"-current-"+info.ReadCurrentHexChecksumString())

			panic(err)
		} else {
			fmt.Println(keyFilePath, "- has same checksum!")
		}
	}
}
