package recursivepaths

import (
	"gitlab.com/auk-go/errorwrapper/errdata/errstr"
	"gitlab.com/auk-go/pathhelper/pathrecurseinfo"
)

func AllOptionsExcept(
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
		IsRecursive:            true,
		IsIncludeAll:           true,
		IsExpandEnvironmentVar: isExpandEnv,
		IsNormalize:            isNormalize,
	}

	return instruction.StringsResults()
}
