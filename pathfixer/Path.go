package pathfixer

import (
	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/pathhelper/expandpath"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

type Location struct {
	PathOptions
	Path         string `json:"Location,omitempty"` // empty path will be ignored from applying.
	compiledPath *string
}

func NewLocation(location string) *Location {
	return &Location{
		PathOptions: PathOptions{},
		Path:        location,
	}
}

func NewLocationUsingOptions(
	location string,
	options PathOptions,
) *Location {
	return &Location{
		PathOptions: options,
		Path:        location,
	}
}

func (it *Location) IsEmptyPath() bool {
	return it == nil || it.Path == constants.EmptyString
}

func (it *Location) HasPath() bool {
	return it != nil && it.Path != constants.EmptyString
}

// IsPathExist returns true if exist on the file system
func (it *Location) IsPathExist() bool {
	return it.HasPath() && chmodhelper.IsPathExists(it.CompiledPath())
}

func (it *Location) CompiledPath() string {
	if it.IsEmptyPath() {
		return constants.EmptyString
	}

	if it.compiledPath != nil {
		return *it.compiledPath
	}

	normalized := normalize.PathUsingSeparatorUsingSingleIf(
		it.IsNormalize,
		osconsts.PathSeparator,
		it.Path)

	expandPath := expandpath.ExpandVariablesIf(
		it.IsExpandEnvVar,
		normalized)

	it.compiledPath = &expandPath

	return *it.compiledPath
}

func (it *Location) ClonePath() *Location {
	if it == nil {
		return nil
	}

	return &Location{
		Path:        it.Path,
		PathOptions: *it.ClonePathOptions(),
	}
}
