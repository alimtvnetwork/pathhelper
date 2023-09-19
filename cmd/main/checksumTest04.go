package main

import (
	"fmt"

	"gitlab.com/auk-go/asynchelper/syncparallel"
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/errorwrapper/errwrappers"
	"gitlab.com/auk-go/pathhelper/filestate"
	"gitlab.com/auk-go/pathhelper/hashas"
	"gitlab.com/auk-go/pathhelper/hexchecksum"
)

func checksumTest04() {
	files := append(pkgRootFiles())
	var detailedResult *hexchecksum.DetailedResult
	var fileStateMappedInfoItems *filestate.MappedInfoItems
	errCollection := errwrappers.Empty()
	hashMethod := hashas.DefaultFastHashMethod

	requestSample := hexchecksum.FilesRequest{
		Method:                     hashMethod,
		IsContinueOnError:          true,
		IsSortFileNames:            false,
		IsSortFilesChecksum:        true,
		IsGenerateContentsChecksum: true,
		IsGenerateFileListChecksum: true,
		Files:                      files,
	}

	// requestSample.SortFileNames()
	// fmt.Println(strings.Join(requestSample.Files, constants.NewLineUnix))

	syncparallel.Tasks(
		func() {
			detailedResult = hexchecksum.DetailedResultOfRequestAsync(&requestSample)

			detailedResult.ErrorWrapper.LogWithTraces()
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
				info.FullPath+constants.SpaceHyphenAngelBracketSpace+"doesn't match checksum",
				checkSum,
				info.HexContentChecksum+"-current-"+info.ReadCurrentHexChecksumString())

			panic(err)
		} else {
			fmt.Println(keyFilePath, "- has same checksum!")
		}
	}

	stateMapCheckSum := fileStateMappedInfoItems.CompiledChecksumString(requestSample.IsSortFilesChecksum)
	fmt.Println("detailedResult.HexFilesContentsChecksum == stateMapCheckSum :", detailedResult.HexFilesContentsChecksum == stateMapCheckSum)
	// fmt.Println(detailedResult.JsonPtr().PrettyJsonString())
}
