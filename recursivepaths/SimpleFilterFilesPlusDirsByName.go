package recursivepaths

import (
	"gitlab.com/auk-go/errorwrapper/errdata/errstr"
	"gitlab.com/auk-go/pathhelper/pathfuncs"
	"gitlab.com/auk-go/pathhelper/pathrecurseinfo"
)

func SimpleFilterFilesPlusDirsByName(
	isRecursive bool,
	excludingRootName []string,
	excludingPaths []string,
	filter pathfuncs.SimpleFilter,
	rootPath string,
) *errstr.Results {
	instruction := pathrecurseinfo.Instruction{
		Root:                   rootPath,
		ExcludingRootNames:     excludingRootName,
		ExcludingPaths:         excludingPaths,
		IsIncludeFilesOnly:     true,
		IsRelativePath:         false,
		IsIncludeDirsOnly:      true,
		IsIncludeAll:           false,
		IsExcludeRoot:          true,
		IsRecursive:            isRecursive,
		IsExpandEnvironmentVar: false,
		IsNormalize:            false,
	}

	return instruction.Result().SimpleFilterFullPathsAsync(filter)
}
