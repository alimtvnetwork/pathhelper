package main

import (
	"fmt"

	"gitlab.com/auk-go/pathhelper/filestate"
	"gitlab.com/auk-go/pathhelper/pathsconst"
)

func fileStateTest01() {
	info, errWrap := filestate.NewInfoDefault(pathsconst.RootDir)

	errWrap.HandleError()

	fmt.Println(info.String())
}
