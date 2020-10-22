package main

import (
	"fmt"

	"gitlab.com/evatix-go/pathhelper"
)

func main() {
	eep := pathhelper.GetExecutableEnvironmentPaths()
	fmt.Println(eep)
}
