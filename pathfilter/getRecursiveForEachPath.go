package pathfilter

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/pathhelper/recursivepaths"

	"gitlab.com/evatix-go/pathhelper/internal/splitinternal"
	"gitlab.com/evatix-go/pathhelper/pathext"
	"gitlab.com/evatix-go/pathhelper/pathgetter"
)

func getRecursiveForEachPath(
	separator string,
	eachPath string,
	filter *Query,
) *errstr.Results {
	eachPathExtWrapper := pathext.New(eachPath)
	isPossibilityOfMatchingExtensionAndFile :=
		eachPathExtWrapper.HasExtension() &&
			eachPathExtWrapper.IsFile()

	if isPossibilityOfMatchingExtensionAndFile {
		return errstr.EmptyErrorResults(eachPath)
	}

	isMatchesWithAnyExtension :=
		isPossibilityOfMatchingExtensionAndFile &&
			eachPathExtWrapper.IsExtensionFiltersMatch(
				filter.Extensions(),
				filter.ExtensionsLength())

	if isMatchesWithAnyExtension {
		return errstr.EmptyErrorResults(eachPath)
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

	files := recursivepaths.Files(
		eachPath,
	)

	if files.HasIssuesOrEmpty() {
		return files
	}

	collection := corestr.NewCollectionUsingStrings(
		files.Values,
		false)

	results := getFilteredFilesByExtensions(
		collection,
		filter)

	return &errstr.Results{
		Values:       *results,
		ErrorWrapper: files.ErrorWrapper,
	}
}
