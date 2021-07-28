package pathsconst

import (
	"os"
)

var (
	TempDir            = os.TempDir()
	DefaultTempTestDir = TempDir + "/pkg-testing/"
	TestDirPatternName = "cimux-tests"
	UnixTemp           = "/tmp/"
	RootRelativeDir    = ".."
	RootDir            = getRoot()
	ExecutableDir      = getExecutableDirectory()
)
