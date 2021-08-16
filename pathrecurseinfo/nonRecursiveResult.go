package pathrecurseinfo

import (
	"io/ioutil"
	"strings"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/pathjoin"
)

func nonRecursiveResult(
	normalizedRoot string,
	instruction *Instruction,
	stat *chmodhelper.PathExistStat,
) *Result {
	fileInfos, err := ioutil.ReadDir(normalizedRoot)
	if err != nil {
		errW := errnew.PathMessages(
			errtype.PathExpand,
			normalizedRoot)

		return InvalidResult(
			normalizedRoot,
			errW,
			stat)
	}

	isExcludeAny := instruction.HasExcludingRootNames()
	nameExcludes := instruction.ExcludingNamesHashset()
	excludingPaths := instruction.ExcludingPathsHashset()
	hasAnyExcludingPaths := excludingPaths.Length() > 0
	isUseLibFunc := !instruction.IsNormalize

	if excludingPaths.Has(normalizedRoot) {
		return &Result{
			Root:            normalizedRoot,
			PathStat:        stat,
			IsInvalidResult: false,
			PathsResult: &PathsResult{
				ExpandingPaths: corestr.EmptySimpleSlice(),
				IsExist:        true,
				IsFile:         false,
				IsDir:          true,
			},
			IsRelative: instruction.IsRelativePath,
			ErrWrapper: errnew.EmptyPtr,
		}
	}

	paths := make(
		[]string,
		0,
		len(fileInfos)+5)

	if !instruction.IsRelativePath && !instruction.IsExcludeRoot {
		paths = append(paths, normalizedRoot)
	}

	for _, info := range fileInfos {
		if isExcludeAny && nameExcludes.Has(info.Name()) {
			continue
		}

		fullPath := pathjoin.JoinSimpleIf(
			isUseLibFunc,
			normalizedRoot,
			info.Name())

		if instruction.IsExcludeRoot && fullPath == normalizedRoot {
			continue
		}

		if hasAnyExcludingPaths && excludingPaths.Has(fullPath) {
			continue
		}

		if instruction.IsRelativePath {
			fullPath = strings.TrimPrefix(fullPath, normalizedRoot)
		}

		if fullPath == "" {
			continue
		}

		if instruction.IsRelativePath &&
			strings.HasPrefix(fullPath, osconsts.PathSeparator) {
			fullPath = fullPath[1:]
		}

		switch {
		case instruction.IsIncludeAll:
			paths = append(paths, fullPath)
		case instruction.IsIncludeDirsOnly && info.IsDir():
			paths = append(paths, fullPath)
		case instruction.IsIncludeFilesOnly && !info.IsDir():
			paths = append(paths, fullPath)
		}
	}

	return &Result{
		Root:            normalizedRoot,
		PathStat:        stat,
		IsInvalidResult: false,
		PathsResult: &PathsResult{
			ExpandingPaths: corestr.NewSimpleSliceUsing(false, paths),
			IsExist:        true,
			IsFile:         false,
			IsDir:          true,
		},
		IsRelative: instruction.IsRelativePath,
		ErrWrapper: errnew.EmptyPtr,
	}
}
