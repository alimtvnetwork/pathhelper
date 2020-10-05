package main

import (
	"fmt"

	"gitlab.com/evatix-go/pathhelper"
)

func main() {
	fmt.Println("GetExecutablePath")
	fmt.Println(pathhelper.GetExecutableDirectory())
}
