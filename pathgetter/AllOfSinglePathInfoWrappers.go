package pathgetter

import (
	"gitlab.com/auk-go/pathhelper/fileinfo"
	"gitlab.com/auk-go/pathhelper/normalize"
)

func AllOfSinglePathInfoWrappers(
	separator string,
	isNormalize bool,
	exploringPath string,
) *fileinfo.Wrappers {
	normalizedPath := normalize.PathUsingSeparatorUsingSingleIf(
		isNormalize,
		separator,
		exploringPath)

	return fileinfo.NewWrappersPtr(
		normalizedPath,
		separator,
		false)
}
