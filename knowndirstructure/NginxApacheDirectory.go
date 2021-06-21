package knowndirstructure

import (
	"os"

	"gitlab.com/evatix-go/core/filemode"
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/internal/createdirinternal"
	"gitlab.com/evatix-go/pathhelper/internal/fsinternal"
	"gitlab.com/evatix-go/pathhelper/internal/normalizeinternal"
	"gitlab.com/evatix-go/pathhelper/internal/pathgetterinternal"
)

type NginxApacheDirectory struct {
	Root             string `json:"Root,omitempty"`
	RootConfigFile   string `json:"RootConfigFile,omitempty"`
	ConfigAvailable  string `json:"ConfigAvailable,omitempty"`
	ConfigEnabled    string `json:"ConfigEnabled,omitempty"`
	SitesAvailable   string `json:"SitesAvailable,omitempty"`
	SitesEnabled     string `json:"SitesEnabled,omitempty"`
	ExtraConfig      string `json:"ExtraConfig,omitempty"`
	ModulesAvailable string `json:"ModulesAvailable,omitempty"`
	ModulesEnabled   string `json:"ModulesEnabled,omitempty"`
}

func (receiver *NginxApacheDirectory) IsRootExist() bool {
	return fsinternal.IsPathExists(receiver.Root)
}

func (receiver *NginxApacheDirectory) IsRootConfigFile() bool {
	return fsinternal.IsPathExists(receiver.RootConfigFile)
}

func (receiver *NginxApacheDirectory) IsConfigAvailable() bool {
	return fsinternal.IsPathExists(receiver.ConfigAvailable)
}

func (receiver *NginxApacheDirectory) IsConfigEnabled() bool {
	return fsinternal.IsPathExists(receiver.ConfigEnabled)
}

func (receiver *NginxApacheDirectory) IsSitesAvailable() bool {
	return fsinternal.IsPathExists(receiver.SitesAvailable)
}

func (receiver *NginxApacheDirectory) IsSitesEnabled() bool {
	return fsinternal.IsPathExists(receiver.SitesEnabled)
}

func (receiver *NginxApacheDirectory) IsExtraConfig() bool {
	return fsinternal.IsPathExists(receiver.ExtraConfig)
}

func (receiver *NginxApacheDirectory) IsModulesAvailable() bool {
	return fsinternal.IsPathExists(receiver.ModulesAvailable)
}

func (receiver *NginxApacheDirectory) IsModulesEnabled() bool {
	return fsinternal.IsPathExists(receiver.ModulesEnabled)
}

func (receiver *NginxApacheDirectory) MkDirRoot(mode os.FileMode) *errorwrapper.Wrapper {
	return createdirinternal.AllRecurse(receiver.Root, mode)
}

func (receiver *NginxApacheDirectory) MkDirConfigAvailable(mode os.FileMode) *errorwrapper.Wrapper {
	return createdirinternal.AllRecurse(receiver.ConfigAvailable, mode)
}

func (receiver *NginxApacheDirectory) MkDirConfigEnabled(mode os.FileMode) *errorwrapper.Wrapper {
	return createdirinternal.AllRecurse(receiver.ConfigEnabled, mode)
}

func (receiver *NginxApacheDirectory) MkDirSitesAvailable(mode os.FileMode) *errorwrapper.Wrapper {
	return createdirinternal.AllRecurse(receiver.SitesAvailable, mode)
}

func (receiver *NginxApacheDirectory) MkDirSitesEnabled(mode os.FileMode) *errorwrapper.Wrapper {
	return createdirinternal.AllRecurse(receiver.SitesEnabled, mode)
}

func (receiver *NginxApacheDirectory) MkDirExtraConfig(mode os.FileMode) *errorwrapper.Wrapper {
	return createdirinternal.AllRecurse(receiver.ExtraConfig, mode)
}

func (receiver *NginxApacheDirectory) MkDirModulesAvailable(mode os.FileMode) *errorwrapper.Wrapper {
	return createdirinternal.AllRecurse(receiver.ModulesAvailable, mode)
}

func (receiver *NginxApacheDirectory) MkDirModulesEnabled(mode os.FileMode) *errorwrapper.Wrapper {
	return createdirinternal.AllRecurse(receiver.ModulesEnabled, mode)
}

func (receiver *NginxApacheDirectory) MkDirAll(mode os.FileMode) *errorwrapper.Wrapper {
	errCollection := errwrappers.Empty()

	errCollection.AddWrapperPtr(receiver.MkDirRoot(mode))
	errCollection.AddWrapperPtr(receiver.MkDirConfigAvailable(mode))
	errCollection.AddWrapperPtr(receiver.MkDirConfigEnabled(mode))
	errCollection.AddWrapperPtr(receiver.MkDirSitesAvailable(mode))
	errCollection.AddWrapperPtr(receiver.MkDirSitesEnabled(mode))
	errCollection.AddWrapperPtr(receiver.MkDirExtraConfig(mode))
	errCollection.AddWrapperPtr(receiver.MkDirModulesAvailable(mode))
	errCollection.AddWrapperPtr(receiver.MkDirModulesEnabled(mode))

	return errCollection.GetAsErrorWrapperPtr()
}

func (receiver *NginxApacheDirectory) MkDirAllDefault() *errorwrapper.Wrapper {
	return receiver.MkDirAll(filemode.X644)
}

func (receiver *NginxApacheDirectory) CombinedSitesAvailable(
	combinedPaths ...string,
) (
	first string,
	allCombinedPaths []string,
) {
	return normalizeinternal.PathsCombine(
		receiver.SitesAvailable,
		combinedPaths)
}

func (receiver *NginxApacheDirectory) CombinedSitesEnabled(
	combinedPaths ...string,
) (
	first string,
	allCombinedPaths []string,
) {
	return normalizeinternal.PathsCombine(
		receiver.SitesEnabled,
		combinedPaths)
}

func (receiver *NginxApacheDirectory) CombinedRoot(
	combinedPaths ...string,
) (
	first string,
	allCombinedPaths []string,
) {
	return normalizeinternal.PathsCombine(
		receiver.Root,
		combinedPaths)
}

func (receiver *NginxApacheDirectory) AllFilesAtSitesAvailable() *errstr.Results {
	return pathgetterinternal.GetAllFiles(
		true,
		osconsts.PathSeparator,
		receiver.SitesAvailable)
}

func (receiver *NginxApacheDirectory) AllFilesAtSitesEnabled() *errstr.Results {
	return pathgetterinternal.GetAllFiles(
		true,
		osconsts.PathSeparator,
		receiver.Root)
}

func (receiver *NginxApacheDirectory) AllPathsAtRoot() *errstr.Results {
	return pathgetterinternal.GetAllPaths(
		true,
		osconsts.PathSeparator,
		receiver.Root)
}

func (receiver *NginxApacheDirectory) AllFilesAtRoot() *errstr.Results {
	return pathgetterinternal.GetAllFiles(
		true,
		osconsts.PathSeparator,
		receiver.Root)
}

func (receiver *NginxApacheDirectory) IsAllExist() bool {
	return receiver.IsRootExist() &&
		receiver.IsRootConfigFile() &&
		receiver.IsConfigAvailable() &&
		receiver.IsConfigEnabled() &&
		receiver.IsSitesAvailable() &&
		receiver.IsSitesEnabled() &&
		receiver.IsExtraConfig() &&
		receiver.IsModulesAvailable()
}
