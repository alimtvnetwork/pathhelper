package constants

import "os"

const (
	OtherPathSeparator   = "/"
	WindowsPathSeparator = "\\"
	WindowsOS            = "windows"
	// filePrefix              = "file:"
	DoubleBackSlash = "\\\\"
	BackSlash       = "\\"
	// "//"
	DoubleForwardSlash = "//"
	// "/"
	ForwardSlash              = "/"
	UriSchemePrefixStandard   = "file:///"
	UriSchemePrefixTwoSlashes = "file://"
	Underscore                = "_"
	Dash                      = "-"
	DoubleDash                = "--"
	DoubleUnderscore          = "__"
	EmptyString               = ""
	MinusOne                  = -1
	GoPath                    = "GOPATH"
	GoBinPath                 = "GOBIN"
	Go111ModuleEnvironment    = "GO111MODULE"
	// "on"
	On            = "on"
	PathSeparator = string(os.PathSeparator)
	Dollar        = "$"
	One           = 1

	InvalidEmptyPathErrorMessage = "Invalid : Empty path given, cannot process it."
)
