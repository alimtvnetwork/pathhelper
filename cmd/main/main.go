package main

import (
	"fmt"

	"gitlab.com/evatix-go/pathhelper"
)

var remove = []string{"/"}

func main() {
	// fmt.Println("IsGoModuleOn")
	// fmt.Println(os.ExpandEnv("%JAVA_HOME%"))
	// fmt.Println(os.ExpandEnv("~/home"))

	fmt.Println(pathhelper.RemoveFromPath("abc/ecma/\\\\sad", &remove, true))
	fmt.Println(pathhelper.RemoveFromPath("abc/ecma/\\\\sad", &remove, false))

	fmt.Println(pathhelper.GetPathFromUri("file:///abc/ecma\\/sad", true))
	fmt.Println(pathhelper.GetPathFromUri("", false))

	if pathhelper.IsWindows(){
		fmt.Println("works")
	}
}
