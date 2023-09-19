package knowndirget

import (
	"go/build"
	"os"

	"gitlab.com/auk-go/core/constants"

	"gitlab.com/auk-go/pathhelper/internal/ispathinternal"
)

// Reference: https://stackoverflow.com/a/32650077
// Returns env go path : os.Getenv(constants.GoPath) or build.Default.GOPATH
func GoPath() string {
	goPath := os.Getenv(constants.GoPath)

	if ispathinternal.Exists(goPath) {
		goPath = build.Default.GOPATH
	}

	return goPath
}
