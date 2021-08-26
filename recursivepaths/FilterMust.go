package recursivepaths

import (
	"gitlab.com/evatix-go/pathhelper/pathfuncs"
)

func FilterMust(
	filter pathfuncs.Filter,
	rootPath string,
) []string {
	results := Filter(filter, rootPath)
	results.ErrorWrapper.HandleError()

	return results.ValueNonPtr()
}
