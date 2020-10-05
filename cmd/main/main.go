package main

import (
	"fmt"

	"gitlab.com/evatix-go/pathhelper"
)

func main() {
	fmt.Println("Hello World from main")
	fmt.Println(pathhelper.GetExecutableDirectory())
}
