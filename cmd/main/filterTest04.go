package main

import (
	"fmt"

	"gitlab.com/auk-go/errorwrapper/errtype"
	"gitlab.com/auk-go/pathhelper/pathsconst"
	"gitlab.com/auk-go/pathhelper/recursivepaths"
)

func filterTest04() {
	files := recursivepaths.SimpleFilterFilesPlusDirsByName(
		false,
		nil,
		nil,
		func(fullPath string) (isTake bool, err error) {
			isMatch := fullPath == "D:\\others-git\\gitlabs\\pathhelper\\tests"

			if isMatch {
				err = errtype.AlreadyDefined.ReferencesLinesError("", "D:\\others-git\\gitlabs\\pathhelper\\tests")
			}

			return fullPath == "D:\\others-git\\gitlabs\\pathhelper\\tests", err
		},
		pathsconst.RootDir,
	)

	files.ErrorWrapper.HandleError()

	fmt.Println(files.String())
}
