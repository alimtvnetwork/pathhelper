package main

import (
	"fmt"

	"gitlab.com/evatix-go/pathhelper"
)

func main() {
	eep := pathhelper.GetExecutableEnvironmentPathCollection()
	fmt.Println(eep)

	fmt.Println(pathhelper.GetWidowsDirectory())
}
