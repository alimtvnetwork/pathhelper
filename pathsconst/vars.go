package pathsconst

import (
	"path"

	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/core/osconsts"
)

var (
	DefaultTemp                    = chmodhelper.TempDirDefault                            // eg. unix : /tmp, windows: %temp%
	TempPermanentDir               = path.Clean(chmodhelper.TempDirGetter.TempPermanent()) // /var/tmp/
	DefaultTempTestDir             = TempPermanentDir + "/pkg-testing/"
	UnixTemp                       = "/tmp/"
	RootRelativeDir                = ".."
	RootDir                        = getRoot()
	ExecutableDir                  = getExecutableDirectory()
	AppWindowsUserSpecificTempRoot = DefaultTemp + osconsts.PathSeparator + AppTemplateName               // %temp%\{app-name}
	TempAppRoot                    = TempPermanentDir + osconsts.PathSeparator + AppTemplateName          // /var/tmp/{app-name}
	TempAppTestRoot                = DefaultTemp + osconsts.PathSeparator + AppTemplateName + "-test-env" // /tmp/{app-name}-test-env
)
