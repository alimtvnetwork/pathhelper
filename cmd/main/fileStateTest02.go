package main

import (
	"fmt"

	"gitlab.com/auk-go/pathhelper/filestate"
	"gitlab.com/auk-go/pathhelper/hashas"
	"gitlab.com/auk-go/pathhelper/pathsconst"
	"gitlab.com/auk-go/pathhelper/recursivepaths"
)

func fileStateTest02() {
	files := recursivepaths.FilesOptionsExcept(
		false,
		false,
		[]string{".git"},
		nil,
		pathsconst.RootDir)

	info, errWrap := filestate.NewInfoCollectionUsingFilePathsAsync(
		hashas.DefaultFastHashMethod,
		true,
		files.Values...)

	errWrap.HandleError()

	fmt.Println(info.String())
}
