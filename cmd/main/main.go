package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("IsGoModuleOn")
	fmt.Println(os.ExpandEnv("%JAVA_HOME%"))
	fmt.Println(os.ExpandEnv("~/home"))
}
