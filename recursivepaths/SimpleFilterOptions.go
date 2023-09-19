package recursivepaths

import (
	"gitlab.com/auk-go/errorwrapper/errdata/errstr"
	"gitlab.com/auk-go/pathhelper/pathfuncs"
	"gitlab.com/auk-go/pathhelper/pathrecurseinfo"
)

func SimpleFilterOptions(
	isNormalize, isExpandEnv bool,
	filter pathfuncs.SimpleFilter,
	rootPath string,
) *errstr.Results {
	instruction := pathrecurseinfo.Instruction{
		Root:                   rootPath,
		IsIncludeFilesOnly:     false,
		IsRelativePath:         false,
		IsIncludeDirsOnly:      false,
		IsIncludeAll:           true,
		IsRecursive:            true,
		IsExpandEnvironmentVar: isExpandEnv,
		IsNormalize:            isNormalize,
	}

	result := instruction.Result()

	return result.SimpleFilterFullPathsAsync(filter)
}

func SimpleFilterFilesOptions(
	isNormalize, isExpandEnv bool,
	filter pathfuncs.SimpleFilter,
	rootPath string,
) *errstr.Results {
	instruction := pathrecurseinfo.Instruction{
		Root:                   rootPath,
		IsIncludeFilesOnly:     false,
		IsRelativePath:         false,
		IsIncludeDirsOnly:      false,
		IsIncludeAll:           true,
		IsRecursive:            true,
		IsExpandEnvironmentVar: isExpandEnv,
		IsNormalize:            isNormalize,
	}

	result := instruction.Result()

	return result.SimpleFilterFullPathsAsync(filter)
}
