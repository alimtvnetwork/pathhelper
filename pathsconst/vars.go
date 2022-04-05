package pathsconst

import (
	"os"

	"gitlab.com/evatix-go/core/osconsts"
)

var (
	TempDir            = os.TempDir()
	DefaultTempTestDir = TempDir + "/pkg-testing/"
	UnixTemp           = "/tmp/"
	RootRelativeDir    = ".."
	RootDir            = getRoot()
	ExecutableDir      = getExecutableDirectory()
	TempAppRoot        = TempDir + osconsts.PathSeparator + AppNameLower
	TempAppTestRoot    = TempDir + osconsts.PathSeparator + AppNameLower + "-test-env" // /tmp/{app-name}-test-env
)
