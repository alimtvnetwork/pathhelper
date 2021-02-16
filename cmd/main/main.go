package main

import (
	"fmt"
	"strings"

	"gitlab.com/evatix-go/pathhelper/internal/recursiveinternal"
)

func main() {
	// eep := pathhelper.GetExecutableEnvironmentPathCollection()
	// fmt.Println(eep)
	//
	// fmt.Println(pathhelper.GetWidowsDirectory())

	// collection := recursiveinternal.GetPaths("D:\\github\\Evatix\\text-replace-automation\\SampleFiles", 500, true)

	collection, ew := recursiveinternal.GetDirectoryPaths(
		"\\",
		"D:\\github\\Evatix\\text-replace-automation\\SampleFiles\\From\\FolderSkip",
		true)

	ew.Handle()

	fmt.Println(strings.Join(*collection, "\n\t"))
}
