package main

import (
	"fmt"

	"gitlab.com/evatix-go/pathhelper/nginxlinuxpath"
)

func main() {
	// eep := pathhelper.GetExecutableEnvironmentPaths()
	// fmt.Println(eep)

	fmt.Println(nginxlinuxpath.GetConf())
}
