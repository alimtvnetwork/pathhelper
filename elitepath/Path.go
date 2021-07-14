package elitepath

import (
	"os"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errdata/errbyte"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/fs"
	"gitlab.com/evatix-go/pathhelper/pathchmod"
	"gitlab.com/evatix-go/pathhelper/pathext"
	"gitlab.com/evatix-go/pathhelper/pathfixer"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
	"gitlab.com/evatix-go/pathhelper/pathinsfmtexec/namegroup"
	"gitlab.com/evatix-go/pathhelper/pathinsfmtexec/pathmodifierverify"
	"gitlab.com/evatix-go/pathhelper/pathjoin"
	"gitlab.com/evatix-go/pathhelper/pathstatlinux"
	"gitlab.com/evatix-go/pathhelper/pathwrapper"
)

type Path struct {
	pathfixer.Location
	fixedPathSlice []string
}

func (it *Path) CombineToEnhance(
	relativePaths ...string,
) *Path {
	combinedPath := it.Combine(relativePaths...)

	return it.ClonePathUsingNew(
		combinedPath)
}

func (it *Path) Combine(relativePaths ...string) string {
	if it == nil {
		return pathjoin.JoinBaseDirWithSep(
			true,
			false,
			false,
			osconsts.PathSeparator,
			"",
			relativePaths...)
	}

	return pathjoin.JoinBaseDirWithSep(
		true,
		it.IsExpandEnvVar,
		it.IsNormalize,
		osconsts.PathSeparator,
		it.Location.Path,
		relativePaths...)
}

func (it *Path) enhancePathsToSliceStrings(enhancePaths ...*Path) []string {
	slice := make([]string, 0, len(enhancePaths))

	if len(enhancePaths) == 0 {
		return slice
	}

	for _, enhancePath := range enhancePaths {
		if enhancePath == nil {
			continue
		}

		slice = append(slice, enhancePath.CompiledPath())
	}

	return slice
}

func (it *Path) CombineWithEnhancePaths(
	enhancePaths ...*Path,
) string {
	relativePaths := it.enhancePathsToSliceStrings(
		enhancePaths...)

	return it.Combine(relativePaths...)
}

func (it *Path) CombineWithEnhancePathsToEnhancePath(
	enhancePaths ...*Path,
) *Path {
	finalPath := it.CombineWithEnhancePaths(
		enhancePaths...)

	return it.ClonePathUsingNew(finalPath)
}

func (it *Path) FileInfo() os.FileInfo {
	return it.ExistStat().FileInfo
}

func (it *Path) FileMode() os.FileMode {
	return it.ExistStat().FileInfo.Mode()
}

func (it *Path) IsDir() bool {
	existStat := it.ExistStat()

	return existStat.IsDir()
}

func (it *Path) IsFile() bool {
	existStat := it.ExistStat()

	return existStat.IsFile()
}

func (it *Path) IsInvalid() bool {
	existStat := it.ExistStat()

	return !existStat.HasFileInfo()
}

func (it *Path) RwxWrapper() *pathchmod.RwxWrapperWithError {
	return pathchmod.ExistingRwxWrapperWithError(it.CompiledPath())
}

func (it *Path) ExistStat() *chmodhelper.PathExistStat {
	return chmodhelper.GetPathExistStat(it.CompiledPath())
}

func (it *Path) SimpleStat() *pathchmod.SimpleStat {
	if it.IsEmptyPath() {
		return nil
	}

	return pathchmod.GetSimpleStat(it.CompiledPath())
}

func (it *Path) PathLinuxStat() *pathstatlinux.Info {
	if it.IsEmptyPath() {
		return pathstatlinux.Get(
			constants.EmptyString)
	}

	return pathstatlinux.Get(it.CompiledPath())
}

func (it *Path) PathWrapper() pathwrapper.Wrapper {
	if it.IsEmptyPath() {
		return constants.EmptyString
	}

	return pathwrapper.Wrapper(it.CompiledPath())
}

func (it *Path) PathExtWrapper() *pathext.Wrapper {
	if it.IsEmptyPath() {
		return pathext.NewPtr(
			constants.EmptyString)
	}

	return pathext.NewPtr(it.CompiledPath())
}

func (it *Path) LocationInfo() *pathhelper.LocationInfo {
	if it.IsEmptyPath() {
		return &pathhelper.LocationInfo{
			RawLocation:           "",
			FileNameWithExtension: "",
			BaseDir:               "",
			FileName:              "",
			DotExtension:          "",
			Extension:             "",
		}
	}

	return pathhelper.GetLocationInfo(
		it.CompiledPath())
}

func (it *Path) ReadFileBytes() *errbyte.Results {
	return fs.ReadFileUsingLock(it.CompiledPath())
}

func (it *Path) ReadFileString() *errstr.Result {
	return fs.ReadFileStringUsingLock(it.CompiledPath())
}

func (it *Path) ReadFileUnmarshal(
	unmarshallingObjectRef interface{},
) *errorwrapper.Wrapper {
	return fs.
		JsonReadUnmarshal(
			it.CompiledPath(),
			unmarshallingObjectRef)
}

func (it *Path) ReadJsonParseSelfInjector(
	injector corejson.JsonParseSelfInjector,
) *errorwrapper.Wrapper {
	return fs.
		ReadJsonParseSelfInjector(
			it.CompiledPath(),
			injector)
}

func (it *Path) ApplyRwxInstruction(
	rwx *chmodins.RwxInstruction,
) *errorwrapper.Wrapper {
	if it.IsEmptyPath() || rwx == nil {
		return errnew.EmptyPtr
	}

	return pathchmod.ApplyChmodRwxOwnerGroupOther(
		rwx.IsRecursive,
		rwx.IsSkipOnInvalid,
		rwx.IsContinueOnError,
		&rwx.RwxOwnerGroupOther,
		it.FixedPathAsSlice())
}

func (it *Path) ApplyPathVerifiers(
	errorCollection *errwrappers.Collection,
	pathVerifiers *pathinsfmt.PathVerifiers,
) (isSuccess bool) {
	if it.IsEmptyPath() || pathVerifiers == nil {
		return true
	}

	return pathmodifierverify.ApplyUsingFlatPaths(
		false,
		pathVerifiers,
		errorCollection,
		it.FixedPathAsSlice())
}

func (it *Path) ApplyPathVerifier(
	isSkipOnInvalid bool,
	errorCollection *errwrappers.Collection,
	pathVerifier *pathinsfmt.PathVerifier,
) (isSuccess bool) {
	if it.IsEmptyPath() || pathVerifier == nil {
		return true
	}

	return pathmodifierverify.ApplyVerifierDirect(
		false,
		false,
		isSkipOnInvalid,
		pathVerifier,
		errorCollection,
		it.FixedPathAsSlice()...)
}

func (it *Path) ApplyChown(
	chown *pathinsfmt.Chown,
) *errorwrapper.Wrapper {
	if it.IsEmptyPath() || chown == nil {
		return errnew.EmptyPtr
	}

	return namegroup.Apply(
		chown.IsRecursive,
		false,
		&chown.UserGroupName,
		it.CompiledPath())
}

func (it *Path) FixedPathAsSlice() []string {
	if it.IsEmptyPath() {
		return []string{}
	}

	if it.fixedPathSlice != nil {
		return it.fixedPathSlice
	}

	it.fixedPathSlice = []string{
		it.CompiledPath(),
	}

	return it.fixedPathSlice
}

func (it *Path) ClonePath() *Path {
	if it == nil {
		return nil
	}

	return &Path{
		Location: *it.Location.ClonePath(),
	}
}

func (it *Path) ClonePathUsingNew(newLocation string) *Path {
	if it == nil {
		return nil
	}

	newPath := Path{
		Location: *it.Location.ClonePath(),
	}

	newPath.Path = newLocation

	return &newPath
}
