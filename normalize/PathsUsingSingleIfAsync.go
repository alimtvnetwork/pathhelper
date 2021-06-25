package normalize

import (
	"sync"

	"gitlab.com/evatix-go/core/osconsts"
)

func PathsUsingSingleIfAsync(
	isNormalizeLongPathForce bool,
	locations []string,
) []string {
	length := len(locations)
	if length == 0 {
		return []string{}
	}

	if !isNormalizeLongPathForce || !osconsts.IsWindows {
		return locations
	}

	newItems := make([]string, length)
	wg := &sync.WaitGroup{}
	wg.Add(length)

	processor := func(index int, location string) {
		newItems[index] = PathUsingSeparatorIf(
			isNormalizeLongPathForce,
			isNormalizeLongPathForce,
			isNormalizeLongPathForce,
			osconsts.PathSeparator,
			location)

		wg.Done()
	}

	for i, location := range locations {
		go processor(i, location)
	}

	wg.Wait()

	return newItems
}
