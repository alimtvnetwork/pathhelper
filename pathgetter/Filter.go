package pathgetter

import (
	"io/ioutil"

	"gitlab.com/evatix-go/pathhelper/normalize"
	"gitlab.com/evatix-go/pathhelper/pathfuncs"
)

func Filter(
	separator, rootPath string,
	isNormalize bool,
	isIgnoreOnError bool,
	filter pathfuncs.Filter,
) *[]*pathfuncs.FilterResult {
	rootPath2 := normalize.PathUsingSeparatorUsingSingleIf(
		isNormalize,
		separator,
		rootPath)

	allPaths, err := ioutil.ReadDir(rootPath2)

	if err != nil {
		empty := make([]*pathfuncs.FilterResult, 0)

		return &empty
	}

	results := make([]*pathfuncs.FilterResult, 0, len(allPaths))

	for _, fileInfo := range allPaths {
		if fileInfo.IsDir() {
			continue
		}

		isDir := fileInfo.IsDir()
		name := fileInfo.Name()
		combinedPath := rootPath2 +
			separator +
			name

		arg := &pathfuncs.FilterArg{
			RootPath:   rootPath,
			FileName:   name,
			FullPath:   combinedPath,
			Separator:  separator,
			IsFile:     !isDir,
			IsDirector: isDir,
			FileInfo:   fileInfo,
		}

		result := filter(arg)

		if result.Wrapper != nil {
			if result.HasError() && isIgnoreOnError {
				continue
			}

			result.HandleError()
		}

		if result.IsKeep {
			results = append(results, result)
		}

		if result.IsBreak {
			return &results
		}
	}

	return &results
}
