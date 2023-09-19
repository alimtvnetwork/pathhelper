package normalize

import (
	"sync"

	"gitlab.com/auk-go/core/coredata/stringslice"
	"gitlab.com/auk-go/core/osconsts"
)

func TrimPrefixUncPaths(
	isSkipOnUnix bool,
	locations ...string,
) []string {
	if isSkipOnUnix && osconsts.IsUnixGroup {
		return locations
	}

	length := len(locations)
	if length == 0 {
		return locations
	}

	wg := &sync.WaitGroup{}
	slice := stringslice.MakeLen(length)

	prefixReplacerFunc := func(index int, source string) {
		slice[index] = TrimPrefixUncPath(source)
	}

	wg.Add(length)
	for i, location := range locations {
		go prefixReplacerFunc(i, location)
	}

	return slice
}
