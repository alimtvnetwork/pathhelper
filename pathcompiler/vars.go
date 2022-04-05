package pathcompiler

import (
	"gitlab.com/evatix-go/enum/osmixtype"
	"gitlab.com/evatix-go/pathhelper/knowndirget"
	"gitlab.com/evatix-go/pathhelper/pathjoin"
	"gitlab.com/evatix-go/pathhelper/pathsconst"
)

var (
	TempAppTestRoot          = pathsconst.TempAppTestRoot // /tmp/{app-name}-test-env/
	TempAppRoot              = pathsconst.TempAppRoot
	windowsProductionDefault = pathjoin.JoinNormalized(
		knowndirget.ProgramFiles(),
		AppName)

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
