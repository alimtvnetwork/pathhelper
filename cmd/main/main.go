package main

import (
	"fmt"

	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
)

func main() {
	fmt.Println("IsGoModuleOn")
	var x string = "      "
	fmt.Println(pathhelpercore.IsEmptyPath(x))
}
