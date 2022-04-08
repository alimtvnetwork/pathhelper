package pathcompilertests

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/core/coreimpl/enumimpl"
	"gitlab.com/evatix-go/core/coretests"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/enum/osmixtype"
	"gitlab.com/evatix-go/pathhelper/pathcompiler"
)

func Test_UnixOsSpecificPathSelect_OnWindows(t *testing.T) {
	coretests.SkipOnUnix(t)

	// Arrange
	expectedMap := map[string]interface{}{
		"AppDbRoot":                "/var/opt/cimux/databases/",
		"ArchiveRoot":              "/var/opt/cimux/archived/",
		"BackupRoot":               "/var/opt/cimux/backups/",
		"CacheTempRoot":            `c:\Windows\Temp\cimux/cache/`,
		"DecompressRoot":           `c:\Windows\Temp\cimux/decompress/`,
		"DefaultConfigFilePath":    "/etc/cimux/config/default-config.json",
		"DefaultEnvPathRoot":       "/var/opt/cimux/env-paths/",
		"DefaultEnvRoot":           "/var/opt/cimux/env/",
		"DefaultInstructionsRoot":  "/var/opt/cimux/instructions/",
		"Description":              "all unix(ubuntu, debian, linux, darwin ...) related os paths",
		"DownloadsRoot":            "/var/opt/cimux/downloads/",
		"EtcAppConfigRoot":         "/etc/cimux/config",
		"EtcAppRoot":               "/etc/cimux",
		"InstructionTempRoot":      `c:\Windows\Temp\cimux/instructions/`,
		"LogAppRoot":               "/var/log/cimux",
		"MigrationCacheRoot":       `c:\Windows\Temp\cimux/migration-cache/`,
		"Name":                     "Unix",
		"PackageTempRoot":          `c:\Windows\Temp\cimux/packages/`,
		"PackagesDownloadRoot":     `/var/opt/cimux/packages-downloaded/`,
		"PackagesRoot":             "/etc/cimux/packages/",
		"PublicRoot":               "/var/www/",
		"ScriptsRoot":              "/var/opt/cimux/scripts/",
		"SnapshotsRoot":            "/var/opt/cimux-snapshots/",
		"SpecificPathFileLocation": "/var/opt/cimux/defined-paths/paths.json",
		"SslRoot":                  "/var/opt/cimux-ssl/",
		"TempRoot":                 `c:\Windows\Temp\cimux`,
		"UserTempRoot":             `c:\Windows\Temp\cimux/users/`,
		"VarAppRoot":               "/var/opt/cimux",
		"VarCacheRoot":             "/var/opt/cimux/cache/",
		"ZipsRoot":                 `/var/opt/cimux/compressed/`,
	}

	// Act
	currentOs := pathcompiler.DefaultApp.By(osmixtype.Unix)
	jsonResult := currentOs.JsonPtr()
	fieldsMap, parsingErr := DeserializedFieldsToMap(jsonResult)
	errcore.MustBeEmpty(parsingErr)

	var dynamicMap enumimpl.DynamicMap = fieldsMap
	diffMessage := dynamicMap.LogShouldDiffMessage(
		true,
		"Unix Specific fields should match exact",
		expectedMap)

	isValid := diffMessage == ""

	// Assert
	convey.Convey("unix os with common values", t, func() {
		convey.So(currentOs, convey.ShouldNotBeNil)
		convey.So(diffMessage, convey.ShouldBeEmpty)
		convey.So(isValid, convey.ShouldBeTrue)
	})
}
