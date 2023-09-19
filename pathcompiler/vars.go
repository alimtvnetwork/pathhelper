package pathcompiler

import (
	"gitlab.com/auk-go/enum/osmixtype"
	"gitlab.com/auk-go/pathhelper/knowndirget"
	"gitlab.com/auk-go/pathhelper/pathjoin"
	"gitlab.com/auk-go/pathhelper/pathsconst"
)

var (
	TempAppTestRoot                = pathsconst.TempAppTestRoot                // /tmp/{app-name}-test-env/
	TempAppRoot                    = pathsconst.TempAppRoot                    // /var/tmp/{app-name}
	AppWindowsUserSpecificTempRoot = pathsconst.AppWindowsUserSpecificTempRoot // %temp%\{app-name}

	windowsProductionDefault = pathjoin.JoinNormalized(
		knowndirget.ProgramFiles(),
		AppName) // c:\program-files\{app-name}

	CurrentOsType = osmixtype.CurrentOsMixType()

	DefaultApp = Basic{
		AppName:      AppName,
		AppNameLower: AppNameLower,
		ProductionMap: map[osmixtype.Variant]*Specific{
			osmixtype.AnyOs:   &UnixOs,
			osmixtype.Ubuntu:  &UnixOs,
			osmixtype.Windows: &WindowsOs,
		},
		TestMap: map[osmixtype.Variant]*Specific{
			osmixtype.AnyOs: &AnyOsTest,
		},
	}
)
