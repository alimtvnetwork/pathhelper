package main

import (
	"fmt"

	"gitlab.com/evatix-go/pathhelper/filestate"
	"gitlab.com/evatix-go/pathhelper/pathsconst"
)

func fileStateTest01() {
	info, errWp := filestate.NewInfoDefault(pathsconst.RootDir)

	errWp.HandleError()

	fmt.Println(info.String())
}
