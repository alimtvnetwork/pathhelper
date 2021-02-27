package pathhelper

import (
	"sync"

	"gitlab.com/evatix-go/core"

	"gitlab.com/evatix-go/pathhelper/pathfuncs"
)

// Don't modify existing paths and creates new one.
func GetAsyncProcessed(
	fullPaths *[]string,
	processor pathfuncs.Processor,
) *[]string {
	if fullPaths == nil {
		return core.EmptyStringsPtr()
	}

	length := len(*fullPaths)
	list := make([]string, length)

	if length == 0 {
		return &list
	}

	wg := &sync.WaitGroup{}
	wg.Add(length)

	inPlaceProcessor := func(index int, fullPath string) {
		defer wg.Done()

		list[index] = processor(index, fullPath)
	}

	for i, fullPath := range *fullPaths {
		go inPlaceProcessor(i, fullPath)
	}

	wg.Wait()

	return &list
}
