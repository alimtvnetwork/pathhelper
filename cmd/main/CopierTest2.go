package main

import (
	"fmt"
	"io/ioutil"

	"gitlab.com/auk-go/pathhelper/copyrecursive"
	"gitlab.com/auk-go/pathhelper/pathsconst"
)

func CopierTest2() {
	tmpDir, _ := ioutil.TempDir("", "ttt")

	fmt.Println(pathsconst.RootDir)
	errWrap := copyrecursive.Do(
		false,
		pathsconst.RootDir,
		tmpDir)

	errWrap.HandleError()
}
