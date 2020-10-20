package constants

import (
	"os"
)

const (
	OtherPathSeparator   = "/"
	WindowsPathSeparator = "\\"
	WindowsOS            = "windows"
	// filePrefix              = "file:"
	DoubleBackSlash = "\\\\"
	TripleBackSlash = "\\\\\\"
	BackSlash       = "\\"
	// "//"
	DoubleForwardSlash        = "//"
	TripleForwardSlash        = "///"
	BackwardAndForwardSlashes = "\\//"
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

	Perm           = 0777
	Dollar         = "$"
	One            = 1
	SemiColon      = ";"
	Path           = "PATH"
	Architecture64 = "X64"
	Architecture32 = "X32"

	// `\\?\UNC\`
	LongPathUncPrefix = `\\?\UNC\`

	// `\\?\`
	LongPathQuestionMarkPrefix = `\\?\`
)
