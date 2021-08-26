package recursivepaths

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/pathhelper/pathrecurseinfo"
)

func DirectoriesOptionsExcept(
	isNormalize,
	isExpandEnv bool,
	skipRootNames []string,
	skipPaths []string,
	rootPath string,
) *errstr.Results {
	instruction := pathrecurseinfo.Instruction{
		Root:                   rootPath,
		ExcludingRootNames:     skipRootNames,
		ExcludingPaths:         skipPaths,
		IsIncludeDirsOnly:      true,
		IsRecursive:            true,
		IsExpandEnvironmentVar: isExpandEnv,
		IsNormalize:            isNormalize,
	}

	return instruction.StringsResults()
}
