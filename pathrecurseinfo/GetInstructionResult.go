package pathrecurseinfo

import (
	"io/fs"
	"path/filepath"
	"strings"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/msgtype"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

// GetInstructionResult TODO implementation not complete, exclude names requires work
func GetInstructionResult(instruction *Instruction) *Result {
	if instruction == nil || instruction.Root == "" {
		return InvalidResult(
			constants.EmptyString,
			errnew.EmptyFilePath,
			nil)
	}

	root := normalize.PathUsingSingleIf(
		instruction.IsNormalize,
		instruction.Root)

	pathStat := chmodhelper.GetPathExistStat(
		root)
	if pathStat.HasError() {
		errW := errnew.Path(
			errtype.MissingPathsOrInvalidPaths,
			pathStat.Error,
			root)

		return InvalidResult(
			root,
			errW,
			pathStat)
	}

	if !pathStat.IsExist {
		// not exist
		errW := errnew.PathMessages(
			errtype.MissingPathsOrInvalidPaths,
			root)

		return InvalidResult(
			root,
			errW,
			pathStat)
	}

	if pathStat.IsFile() {
		return &Result{
			Root:            instruction.Root,
			PathStat:        pathStat,
			IsInvalidResult: false,
			PathsResult: &PathsResult{
				ExpandingPaths: corestr.NewSimpleSliceUsing(false, []string{root}),
				IsExist:        true,
				IsFile:         true,
				IsDir:          false,
			},
			IsRelative: instruction.IsRelativePath,
			ErrWrapper: errnew.EmptyPtr,
		}
	}

	if !instruction.IsRecursive {
		return nonRecursiveResult(
			root,
			instruction,
			pathStat)
	}

	paths := make(
		[]string,
		0,
		constants.ArbitraryCapacity32)

	var sliceErr []string
	isExcludeAny := instruction.HasAnyExcludingCondition()
	nameExcludes := instruction.ExcludingNamesHashset()
	isExcludeRoot := instruction.IsExcludeRoot
	isRelativePath := instruction.IsRelativePath

	filepath.Walk(
		root,
		func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				sliceErr = append(sliceErr, err.Error()+" - "+path)

				return err
			}

			if info == nil {
				sliceErr = append(sliceErr, "Nil file info - "+path)

				return err
			}

			if isExcludeAny && nameExcludes.Has(info.Name()) {
				return nil
			}

			if isExcludeRoot && root == path {
				return nil
			}

			finalizedPath := path
			if isRelativePath {
				finalizedPath = strings.Replace(
					finalizedPath,
					root,
					constants.EmptyString,
					1)
			}

			if finalizedPath == "" {
				return nil
			}

			switch {
			case instruction.IsIncludeAll:
				paths = append(paths, finalizedPath)
			case instruction.IsIncludeDirsOnly && info.IsDir():
				paths = append(paths, finalizedPath)
			case instruction.IsIncludeFilesOnly && !info.IsDir():
				paths = append(paths, finalizedPath)
			}

			return nil
		},
	)

	compiledErr := msgtype.SliceToError(sliceErr)

	return &Result{
		Root:            root,
		PathStat:        pathStat,
		IsInvalidResult: false,
		PathsResult: &PathsResult{
			ExpandingPaths: corestr.NewSimpleSliceUsing(false, paths),
			IsExist:        true,
			IsFile:         false,
			IsDir:          true,
		},
		IsRelative: instruction.IsRelativePath,
		ErrWrapper: errnew.Path(errtype.PathExpand, compiledErr, root),
	}
}
