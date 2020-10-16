package main

import (
	"fmt"

	"gitlab.com/evatix-go/pathhelper"
)

func main() {
	fmt.Println(pathhelper.IsAllPathNotExist("C:\\", "windows\\"))
}
