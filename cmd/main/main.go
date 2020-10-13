package main

import (
	"fmt"
	"strings"

	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
)

func getCombinedPathUsingConfigInternal(
	pathConfig *pathhelpercore.PathConfig,
	paths []string,
) string {
	if pathhelpercore.IsEmptyArray(paths) {
		panic("Empty paths given.")
	}

	if pathConfig == nil {
		pathConfig = pathhelpercore.NewDefaultPathConfig()
	}
	fmt.Println(pathConfig)
	var combinedPath string

	if !pathConfig.IsIgnoreEmptyPath {
		combinedPath = strings.Join(paths, pathConfig.Separator)
	} else {
		combinedPath = pathhelper.GetCombinedOfNonEmptyPaths(pathConfig.Separator, paths)
	}
	fmt.Println("combinedPath", combinedPath)
	if pathConfig.IsNormalize {
		combinedPath = pathhelper.NormalizePath(combinedPath)
	}

	return combinedPath
}
func main() {
	// fmt.Println("IsGoModuleOn")
	// fmt.Println(os.ExpandEnv("%JAVA_HOME%"))
	// fmt.Println(os.ExpandEnv("~/home"))

	fmt.Println(getCombinedPathUsingConfigInternal(&pathhelpercore.PathConfig{IsNormalize: true}, []string{"C:\\", "\\somethingelse\\", "\\etc"}))
	fmt.Println(pathhelper.NormalizePath("c:\\\\windows\\\\users\\\\etc\\\\more"))
	// fmt.Println(pathhelper.GetPathFromUri("file:\\\\c:\\windows\\users\\etc\\more", true))
}
