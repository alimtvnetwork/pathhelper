![](https://gitlab.com/evatix-go/common-assets/-/blob/pathhelper-logos/Assets/logos/pathhelper/60.png)
# pathhelper

path helper utility tool

## Git Clone

`git clone https://gitlab.com/evatix-go/pathhelper.git`

### Prerequisites

- Either add your ssh key to your gitlab account
- Or, use your access token to clone it.

## Installation

`go get gitlab.com/evatix-go/pathhelper`

## Why *pathhelper*?

Package pathhelper provides an easy and fast way to get your desired OS(operating system) functionality 
without the hassle of considering the OS you are on, what GO packages you may need to decode a path or 
simply find out if the path or file exists etc. We have brought features of different packages and 
injected some new features to make this package a complete solution for obtaining information regarding 
filepath independent of platform.

## Examples

```go
package main

import (
	"fmt"
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
)

func main() {
	// Checking if path is empty
	pathhelpercore.IsEmptyPath("") // returns true

	samplePath := "C:\\users\\"

	// Checking if path exists
	exists :=  pathhelper.IsPathExist(samplePath)
	fmt.Println(exists) // returns true if directory or file exist on that path

	// Getting path as URI
	fmt.Println(pathhelper.GetPathAsUri(samplePath, true)) // file:///c:/users

	// Normalize path
	pathToNormalize := "file:///C:/something/otherthing"
	fmt.Println(pathhelper.NormalizePath(pathToNormalize)) // C:\something\otherthing if OS is windows; C:/something/otherthing if OS is Unix
}
```

## Acknowledgement

For this package we have mainly used OS and filepath packages of GO. For testing the package we have used 
the very convenient package *[Go Convey](http://goconvey.co/)*.

## Links

- [What are conventions for filenames in Go? - Stack Overflow](https://stackoverflow.com/questions/25161774/what-are-conventions-for-filenames-in-go)
- [go - Pass method argument to function - Stack Overflow](https://stackoverflow.com/questions/38897529/pass-method-argument-to-function)
- [exec.Command() in Go with environment variable - Stack Overflow](https://stackoverflow.com/questions/51015569/exec-command-in-go-with-environment-variable)
- [go - Pointers vs. values in parameters and return values - Stack Overflow](https://stackoverflow.com/questions/23542989/pointers-vs-values-in-parameters-and-return-values?rq=1)
- [jmhodges/copyfighter: Statically analyzes Go code and reports functions that are passing large structs by value](https://github.com/jmhodges/copyfighter)
- [CodeReviewComments · golang/go Wiki](https://github.com/golang/go/wiki/CodeReviewComments#pass-values)
- [Difference between := and = operators in Go - Stack Overflow](https://stackoverflow.com/questions/17891226/difference-between-and-operators-in-go?rq=1)

## Notes

## Contributors

## License

[Evatix MIT License](/LICENSE)