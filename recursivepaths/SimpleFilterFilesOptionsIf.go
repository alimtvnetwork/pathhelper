package recursivepaths

import (
	"gitlab.com/auk-go/errorwrapper/errdata/errstr"
	"gitlab.com/auk-go/pathhelper/pathfuncs"
	"gitlab.com/auk-go/pathhelper/pathrecurseinfo"
)

func SimpleFilterFilesOptionsIf(
	isRecursive,
	isNormalize, isExpandEnv bool,
	filter pathfuncs.SimpleFilter,
	rootPath string,
) *errstr.Results {
	instruction := pathrecurseinfo.Instruction{
		Root:                   rootPath,
		IsRelativePath:         false,
		IsIncludeFilesOnly:     true,
		IsIncludeDirsOnly:      false,
		IsIncludeAll:           false,
		IsRecursive:            isRecursive,
		IsExpandEnvironmentVar: isExpandEnv,
		IsNormalize:            isNormalize,
	}

	result := instruction.Result()

	return result.SimpleFilterFullPathsAsync(filter)
}
