package testwrappers

import (
	"gitlab.com/evatix-go/core/coredata/corestr"
)

func GetSetupPaths() []string {
	linkedCollection := corestr.NewLinkedCollections()

	for _, ins := range PathsCreateInstructionsUnix {
		linkedCollection.AddStrings(
			ins.LazyFlatPaths()...,
		)
	}

	return linkedCollection.
		ToCollection(0).
		ListStrings()
}
