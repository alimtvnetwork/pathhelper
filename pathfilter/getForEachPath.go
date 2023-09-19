package pathfilter

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/errorwrapper/errdata/errstr"
	"gitlab.com/auk-go/pathhelper/internal/splitinternal"
	"gitlab.com/auk-go/pathhelper/pathext"
	"gitlab.com/auk-go/pathhelper/pathgetter"
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
		return errstr.
			New.
			Results.
			SpreadValuesOnly(
				eachPath)
	}

	isMatchesWithAnyExtension :=
		isPossibilityOfMatchingExtensionAndFile &&
			eachPathExtWrapper.IsExtensionFiltersMatch(
				filter.Extensions(),
				filter.ExtensionsLength())

	if isMatchesWithAnyExtension {
		return errstr.New.Results.SpreadValuesOnly(eachPath)
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

	collection := corestr.New.Collection.Strings(
		files.Values,
	)

	results := getFilteredFilesByExtensions(
		collection,
		filter)

	return errstr.New.Results.Strings(results)
}
