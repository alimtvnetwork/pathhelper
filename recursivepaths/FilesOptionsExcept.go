package recursivepaths

import (
	"gitlab.com/auk-go/errorwrapper/errdata/errstr"
	"gitlab.com/auk-go/pathhelper/pathrecurseinfo"
)

func FilesOptionsExcept(
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
		IsIncludeFilesOnly:     true,
		IsExpandEnvironmentVar: isExpandEnv,
		IsNormalize:            isNormalize,
		IsExcludeRoot:          true,
		IsRecursive:            true,
	}

	return instruction.StringsResults()
}
