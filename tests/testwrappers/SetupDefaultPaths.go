package testwrappers

import (
	"gitlab.com/auk-go/pathhelper/pathinsfmtexec/pathscreateinsexec"
)

func SetupDefaultPathsUnix() []string {
	errCollection := pathscreateinsexec.ApplyPathsCreatorCollectionsReturnErrorCollection(
		true,
		true,
		true,
		false,
		PathsCreateInstructionsUnix)

	if errCollection.HasError() {
		trace := errCollection.FullStringWithTraces()
		panic("Failed to create default paths." + AllPathsString + trace)
	}

	return DefaultWorkingPaths
}
