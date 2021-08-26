package pathfixer

import (
	"os"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/pathhelper/expandpath"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

// Location Don't use Path Directly but use CompiledPath
type Location struct {
	PathOptions
	Path         string `json:"Location,omitempty"` // empty path will be ignored from applying, must use CompiledPath
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

func (it *Location) IsInvalid() bool {
	existStat := it.ExistStat()

	return !existStat.HasFileInfo()
}

// IsPathExist returns true if exist on the file system
func (it *Location) IsPathExist() bool {
	return it.HasPath() && chmodhelper.IsPathExists(it.CompiledPath())
}

func (it *Location) CompiledPath() string {
	if it.compiledPath != nil {
		return *it.compiledPath
	}

	normalizedPath := it.Path

	if normalizedPath != "" {
		expandPath := expandpath.ExpandVariablesIf(
			it.IsExpandEnvVar,
			normalizedPath)

		normalizedPath = normalize.PathUsingSeparatorUsingSingleIf(
			it.IsNormalize,
			osconsts.PathSeparator,
			expandPath)
	}

	it.compiledPath = &normalizedPath

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

func (it *Location) IsEqual(another *Location) bool {
	if !it.IsEqualWithoutOptions(another) {
		return false
	}

	return it.PathOptions.IsEqual(
		&another.PathOptions)
}

func (it *Location) IsEqualWithoutOptions(another *Location) bool {
	if it == nil && another == nil {
		return true
	}

	if it == nil || another == nil {
		return false
	}

	return it.Path == another.Path
}

func (it *Location) ExistStat() *chmodhelper.PathExistStat {
	return chmodhelper.GetPathExistStat(it.CompiledPath())
}

func (it *Location) FileInfo() os.FileInfo {
	fileInfo, _ := os.Stat(it.CompiledPath())

	return fileInfo
}

func (it *Location) FileMode() os.FileMode {
	fileInfo, _ := os.Stat(it.CompiledPath())

	return fileInfo.Mode()
}

func (it *Location) SafeFileMode() os.FileMode {
	fileInfo, _ := os.Stat(it.CompiledPath())

	if fileInfo != nil {
		return fileInfo.Mode()
	}

	return constants.Zero
}

func (it *Location) IsDir() bool {
	existStat := it.ExistStat()

	return existStat.IsDir()
}

func (it *Location) IsInvalidPath() bool {
	existStat := it.ExistStat()

	return !existStat.IsExist
}

func (it *Location) IsExistButDir() bool {
	existStat := it.ExistStat()

	return existStat.IsExist && existStat.IsDir()
}

func (it *Location) IsExistButFile() bool {
	existStat := it.ExistStat()

	return existStat.IsExist && existStat.IsFile()
}

func (it *Location) IsFile() bool {
	existStat := it.ExistStat()

	return existStat.IsFile()
}
