package pathgetter

import (
	"gitlab.com/evatix-go/pathhelper/fileinfo"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

func AllOfSinglePathInfoWrappers(
	separator string,
	isNormalize bool,
	exploringPath string,
) *fileinfo.Wrappers {
	rootPath2 := normalize.PathUsingSeparatorUsingSingleIf(
		isNormalize,
		separator,
		exploringPath)

	return fileinfo.NewWrappersPtr(rootPath2)
}
