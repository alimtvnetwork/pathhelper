package copyrecursive

import "gitlab.com/evatix-go/pathhelper/expandnormalize"

type Options struct {
	IsSkipOnExist      bool
	IsRecursive        bool
	IsClearDestination bool
	IsUseShellOrCmd    bool
	IsNormalize        bool
	IsExpandVar        bool
}

func (it Options) FixedPath(location string) string {
	return expandnormalize.FixIf(
		it.IsNormalize,
		it.IsExpandVar,
		location)
}
