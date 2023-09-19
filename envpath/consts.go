package envpath

import (
	"gitlab.com/auk-go/core/constants"
)

const (
	unixEnvPathSplitter    = constants.Colon
	windowsEnvPathSplitter = constants.SemiColon
	pathEqual              = "PATH="
)
