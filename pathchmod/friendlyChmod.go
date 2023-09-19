package pathchmod

import (
	"fmt"
	"path/filepath"

	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/osconsts"
	"gitlab.com/auk-go/core/simplewrap"
	"gitlab.com/auk-go/errorwrapper/errcmd"
	"gitlab.com/auk-go/pathhelper/normalize"
)

type friendlyChmod struct{}

func (it friendlyChmod) OfPath(
	location string,
) string {
	fileChmod, isInvalid := chmodhelper.
		GetExistingChmodOfValidFile(location)

	if isInvalid {
		return "path invalid : " + simplewrap.WithDoubleQuote(location)
	}

	return chmodhelper.FileModeFriendlyString(fileChmod) +
		constants.SpaceColonSpace +
		simplewrap.WithDoubleQuote(location)
}

func (it friendlyChmod) OfDirFiles(
	location string,
) string {
	if chmodhelper.IsPathInvalid(location) {
		return "path invalid : " + simplewrap.WithDoubleQuote(location)
	}

	parentDir := it.parentDir(location)

	scriptBuilder := errcmd.
		New.
		ScriptBuilder.
		DefaultDependingOnOs()

	if osconsts.IsWindows {
		scriptBuilder.Args(fmt.Sprintf(
			powershellListWithOwner,
			normalize.PathFixWithoutLongPathIf(
				true,
				parentDir)))
	} else {
		scriptBuilder.Args(simplewrap.WithDoubleQuote(parentDir))
	}

	dirChmodDisplay := it.OfPath(location)

	return dirChmodDisplay +
		constants.DefaultLine +
		scriptBuilder.DetailedOutput()
}

func (it friendlyChmod) LogOfPath(
	location string,
) {
	chmodMsg := it.OfPath(location)

	fmt.Println(chmodMsg)
}

func (it friendlyChmod) LogOfDirFiles(
	location string,
) {
	chmodMsg := it.OfDirFiles(location)

	fmt.Println(chmodMsg)
}

func (it friendlyChmod) parentDir(location string) string {
	if chmodhelper.IsDirectory(location) {
		return location
	}

	return filepath.Dir(location)
}
