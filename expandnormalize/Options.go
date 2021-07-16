package expandnormalize

import (
	"path/filepath"

	"gitlab.com/evatix-go/core/coredata/corestr"
)

type Options struct {
	IsNormalize,
	IsExpandEnvVar bool
}

func (it Options) Fix(location string) string {
	return FixIf(
		it.IsNormalize,
		it.IsExpandEnvVar,
		location)
}

func (it Options) Join(
	baseDir string,
	locations ...string,
) string {
	if len(locations) == 0 {
		return it.Fix(baseDir)
	}

	combinedLocations := filepath.Join(locations...)
	finalJoin := filepath.Join(baseDir, combinedLocations)

	return it.Fix(finalJoin)
}

func (it Options) FixPaths(locations ...string) *corestr.SimpleSlice {
	if len(locations) == 0 {
		return corestr.EmptySimpleSlice()
	}

	fixSlice := corestr.NewSimpleSlice(
		len(locations))

	for i, location := range locations {
		fixSlice.Items[i] = it.Fix(location)
	}

	return fixSlice
}
