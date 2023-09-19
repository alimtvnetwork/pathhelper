package pathfuncs

import (
	"os"

	"gitlab.com/auk-go/core/conditional"
	"gitlab.com/auk-go/core/defaultcapacity"
	"gitlab.com/auk-go/errorwrapper/errwrappers"
)

func FilterResults(
	isContinueOnError bool,
	errCollection *errwrappers.Collection,
	filter Filter,
	rootPath string,
	fullPaths ...string,
) []*FilterResult {
	length := len(fullPaths)

	if filter == nil || length == 0 {
		return []*FilterResult{}
	}

	foundItems := make([]*FilterResult, 0, defaultcapacity.OfSearch(length))
	for _, fullPath := range fullPaths {
		fileInfo, err := os.Stat(fullPath)
		hasFileInfo := fileInfo != nil
		fileName := conditional.StringTrueFunc(
			hasFileInfo,
			fileInfo.Name)

		arg := &FilterArg{
			RootPath:    rootPath,
			FileName:    fileName,
			FullPath:    fullPath,
			IsFile:      hasFileInfo && !fileInfo.IsDir(),
			IsDirectory: hasFileInfo && fileInfo.IsDir(),
			FileInfo:    fileInfo,
			InputError:  err,
		}

		filterResult := filter(arg)
		errCollection.AddWrapperPtr(filterResult.ErrorWrapper)

		if filterResult.IsTake {
			foundItems = append(
				foundItems,
				filterResult)
		}

		if filterResult.IsBreak {
			break
		}

		hasErr := filterResult.ErrorWrapper.HasError()
		if !isContinueOnError && hasErr {
			break
		}
	}

	return foundItems
}
