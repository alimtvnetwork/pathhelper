package pathcompilertests

import (
	"strings"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/corevalidator"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/enum/osmixtype"
	"gitlab.com/evatix-go/pathhelper/pathcompiler"
)

func Test_UnixOsSpecificPathSelect(t *testing.T) {
	// Arrange
	// TODO improve expected lines
	expectedPrettyLines := []string{
		"{",
		"\t\"Name\": \"Unix\",",
		"\t\"Description\": \"all unix(ubuntu, debian, linux, darwin ...) related os paths\",",
		"\t\"SpecificPathFileLocation\": \"/var/opt/cimux/defined-paths/paths.json\",",
		"\t\"VarAppRoot\": \"/var/opt/cimux\",",
		"\t\"EtcAppRoot\": \"/etc/cimux\",",
		"\t\"EtcAppConfigRoot\": \"/etc/cimux/config\",",
		"\t\"AppDbRoot\": \"/var/opt/cimux/databases/\",",
		"\t\"TempRoot\": \"C:\\\\Users\\\\ADMINI~1\\\\AppData\\\\Local\\\\Temp\\\\cimux\",",
		"\t\"UserTempRoot\": \"C:\\\\Users\\\\ADMINI~1\\\\AppData\\\\Local\\\\Temp\\\\cimux/users/\",",
		"\t\"CacheTempRoot\": \"C:\\\\Users\\\\ADMINI~1\\\\AppData\\\\Local\\\\Temp\\\\cimux/cache/\",",
		"\t\"InstructionTempRoot\": \"C:\\\\Users\\\\ADMINI~1\\\\AppData\\\\Local\\\\Temp\\\\cimux/instructions/\",",
		"\t\"MigrationCacheRoot\": \"C:\\\\Users\\\\ADMINI~1\\\\AppData\\\\Local\\\\Temp\\\\cimux/migration-cache/\",",
		"\t\"PackageTempRoot\": \"C:\\\\Users\\\\ADMINI~1\\\\AppData\\\\Local\\\\Temp\\\\cimux/packages/\",",
		"\t\"LogAppRoot\": \"/var/log/cimux\",",
		"\t\"VarCacheRoot\": \"/var/opt/cimux/cache/\",",
		"\t\"DownloadsRoot\": \"/var/opt/cimux/downloads/\",",
		"\t\"ScriptsRoot\": \"/var/opt/cimux/scripts/\",",
		"\t\"DecompressRoot\": \"C:\\\\Users\\\\ADMINI~1\\\\AppData\\\\Local\\\\Temp\\\\cimux/decompress/\",",
		"\t\"PackagesRoot\": \"/etc/cimux/packages/\",",
		"\t\"PackagesDownloadRoot\": \"/var/opt/cimux/packages-downloaded/\",",
		"\t\"DefaultInstructionsRoot\": \"/var/opt/cimux/instructions/\",",
		"\t\"DefaultEnvRoot\": \"/var/opt/cimux/env/\",",
		"\t\"DefaultEnvPathRoot\": \"/var/opt/cimux/env-paths/\",",
		"\t\"BackupRoot\": \"/var/opt/cimux/backups/\",",
		"\t\"ArchiveRoot\": \"/var/opt/cimux/archived/\",",
		"\t\"ZipsRoot\": \"/var/opt/cimux/compressed/\",",
		"\t\"DefaultConfigFilePath\": \"/etc/cimux/config/config/default-config.json\"",
		"}",
	}

	// Act
	currentOs := pathcompiler.DefaultApp.By(osmixtype.Unix)
	lines := strings.Split(
		currentOs.PrettyJsonString(),
		constants.DefaultLine)

	sliceValidator := corevalidator.SliceValidator{
		ActualLines:   lines,
		ExpectedLines: expectedPrettyLines,
	}

	validatorParamBase := corevalidator.ValidatorParamsBase{
		CaseIndex:          0,
		IsAttachUserInputs: true,
		IsCaseSensitive:    true,
	}

	validationFinalError := sliceValidator.AllVerifyError(
		&validatorParamBase)
	isValid := validationFinalError == nil

	// Assert
	convey.Convey("unix os with common values", t, func() {
		convey.So(currentOs, convey.ShouldNotBeNil)
		convey.So(lines, convey.ShouldNotBeNil)
		errcore.ErrPrintWithTestIndex(0, validationFinalError)
		convey.So(isValid, convey.ShouldBeTrue)
	})
}
