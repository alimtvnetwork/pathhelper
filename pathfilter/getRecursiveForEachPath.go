package pathfilter

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"

	"gitlab.com/evatix-go/pathhelper/internal/pathsplitinternal"
	"gitlab.com/evatix-go/pathhelper/internal/recursiveinternal"
	"gitlab.com/evatix-go/pathhelper/pathext"
	"gitlab.com/evatix-go/pathhelper/pathgetter"
)

func getRecursiveForEachPath(
	separator string,
	eachPath string,
	filter *Query,
) *errstr.ResultsWithErrorCollection {
	eachPathExtWrapper := pathext.New(eachPath)
	isPossibilityOfMatchingExtensionAndFile :=
		eachPathExtWrapper.HasExtension() &&
			eachPathExtWrapper.IsFile()
	isMatchesWithAnyExtension :=
		isPossibilityOfMatchingExtensionAndFile &&
			eachPathExtWrapper.IsExtensionFiltersMatch(
				filter.Extensions(),
				filter.ExtensionsLength())

	if isMatchesWithAnyExtension {
		return errstr.
			EmptyResultsWithErrorCollectionPtr()
	}

	if isPossibilityOfMatchingExtensionAndFile {
		// no need process a file
		return errstr.
			EmptyResultsWithErrorCollectionPtr()
	}

	// get all files in the dir.
	isOnlyDot := *eachPathExtWrapper.DotExtension() ==
		constants.Dot

	if isOnlyDot {
		// get all files in the dir.
		dir, _ := pathsplitinternal.GetWithoutSlash(
			eachPath)

		return pathgetter.Files(
			separator,
			dir,
			false)
	}

	files := recursiveinternal.GetFilesPaths(
		separator,
		eachPath,
		false)

	if files.HasError() {
		return files
	}

	collection := corestr.NewCollectionUsingStrings(
		files.ValueMust(),
		false)

	results := getFilteredFilesByExtensions(
		collection,
		filter)

	return &errstr.ResultsWithErrorCollection{
		Values:        results,
		ErrorWrappers: files.ErrorWrappers,
	}
}
