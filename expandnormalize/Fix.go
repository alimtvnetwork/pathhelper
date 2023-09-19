package expandnormalize

import (
	"gitlab.com/auk-go/pathhelper/expandpath"
	"gitlab.com/auk-go/pathhelper/normalize"
)

func Fix(
	location string,
) string {
	location = expandpath.ExpandVariables(
		location)

	return normalize.Path(
		location)
}
