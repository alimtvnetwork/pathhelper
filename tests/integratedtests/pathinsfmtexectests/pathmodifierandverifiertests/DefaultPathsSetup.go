package pathinsfmtexectests

import (
	"gitlab.com/auk-go/pathhelper/createpath"
	"gitlab.com/auk-go/pathhelper/deletepaths"
	"gitlab.com/auk-go/pathhelper/tests/testwrappers/pathinsfmtexectestwrappers"
)

func DefaultPathsSetup() {
	deletepaths.AllOnExist(
		pathinsfmtexectestwrappers.PathOneTextFile,
		pathinsfmtexectestwrappers.PathTwoTextFile,
	).HandleError()

	_, errWrap := createpath.CreateMany(
		true,
		true,
		[]string{
			pathinsfmtexectestwrappers.PathOneTextFile,
			pathinsfmtexectestwrappers.PathTwoTextFile,
		},
	)

	errWrap.HandleError()
}
