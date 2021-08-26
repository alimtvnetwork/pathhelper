package pathfilter

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/pathhelper/internal/splitinternal"
	"gitlab.com/evatix-go/pathhelper/pathext"
	"gitlab.com/evatix-go/pathhelper/pathgetter"
)

func getFilesForEachPath(
	separator string,
	eachPath string,
	filter *Query,
) *errstr.Results {
	eachPathExtWrapper := pathext.New(eachPath)
	isPossibilityOfMatchingExtensionAndFile :=
		eachPathExtWrapper.HasExtension() &&
			eachPathExtWrapper.IsFile()

	if isPossibilityOfMatchingExtensionAndFile {
		return &errstr.Results{
			Values:       &[]string{eachPath},
			ErrorWrapper: errnew.EmptyPtr,
		}
	}

	isMatchesWithAnyExtension :=
		isPossibilityOfMatchingExtensionAndFile &&
			eachPathExtWrapper.IsExtensionFiltersMatch(
				filter.Extensions(),
				filter.ExtensionsLength())

	if isMatchesWithAnyExtension {
		return &errstr.Results{
			Values:       &[]string{eachPath},
			ErrorWrapper: errnew.EmptyPtr,
		}
	}

	// get all files in the dir.
	isOnlyDot := *eachPathExtWrapper.DotExtension() ==
		constants.Dot

	if isOnlyDot {
		// get all files in the dir.
		dir, _ := splitinternal.GetWithoutSlash(
			eachPath)

		return pathgetter.Files(
			false,
			separator,
			dir,
		)
	}

	files := pathgetter.Files(
		false,
		separator,
		eachPath,
	)

	if files.HasError() {
		return files
	}

	collection := corestr.NewCollectionUsingStrings(
		files.ValueMust(),
		false)

	results := getFilteredFilesByExtensions(
		collection,
		filter)

	return &errstr.Results{
		Values:       results,
		ErrorWrapper: errnew.EmptyPtr,
	}
}
