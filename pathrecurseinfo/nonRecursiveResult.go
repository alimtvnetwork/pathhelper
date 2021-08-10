package pathrecurseinfo

import (
	"io/ioutil"
	"strings"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
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

	paths := make(
		[]string,
		0,
		len(fileInfos))

	isExcludeAny := instruction.HasAnyExcludingCondition()
	nameExcludes := instruction.ExcludingNamesHashset()

	for _, info := range fileInfos {
		if isExcludeAny && nameExcludes.Has(info.Name()) {
			continue
		}

		fullPath := pathjoin.JoinSimple(
			normalizedRoot,
			info.Name())

		if instruction.IsExcludeRoot && fullPath == normalizedRoot {
			continue
		}

		if instruction.IsRelativePath {
			fullPath = strings.Replace(
				fullPath,
				normalizedRoot,
				constants.EmptyString,
				1)
		}

		if fullPath == "" {
			continue
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
