package pathfilter

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/normalize"
)

// rootPath should not ends with separator
func GetFiltersArray(
	rootPath string,
	additionalFilters *[]string,
	extensions *[]string,
	isRecursive bool,
	isNormalizePath bool,
) *[]string {
	separator := osconsts.PathSeparator
	rootPath2 := normalize.PathUsingSeparatorIf(
		isNormalizePath,
		isNormalizePath,
		isNormalizePath,
		separator,
		rootPath,
	)

	additionalFilterLength :=
		corestr.LengthOfStrings(additionalFilters)

	extensionsLength :=
		corestr.LengthOfStrings(extensions)

	newExtensionsPtr := extensions

	if extensionsLength == 0 {
		newExtensionsPtr = &(constants.EmptyStrings)
	}

	arg := &recursiveFilterGetterParam{
		separator:               separator,
		rootPath:                rootPath2,
		eachFilterPath:          "", // will be generated in each case
		rootPathPlusSeparator:   rootPath2 + separator,
		extensionsLength:        extensionsLength,
		additionalFiltersLength: additionalFilterLength,
		additionalFilters:       additionalFilters,
		extensions:              newExtensionsPtr,
	}

	if additionalFilterLength == 0 && isRecursive {
		arg.eachFilterPath = rootPath2

		return getRecursiveFilterForEachFilterPath(arg)
	}

	if additionalFilterLength == 0 && !isRecursive {
		arg.eachFilterPath = rootPath2

		return getFilterForEachFilterPath(arg)
	}

	if isRecursive {
		return getRecursiveFilterDefinitions(arg)
	}

	return getFilterDefinitions(arg)
}
