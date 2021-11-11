package nginxlinuxpath

import (
	"os"
	"path/filepath"

	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/extensionsconst"
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/internal/createdirinternal"
	"gitlab.com/evatix-go/pathhelper/internal/fsinternal"
	"gitlab.com/evatix-go/pathhelper/internal/normalizeinternal"
	"gitlab.com/evatix-go/pathhelper/internal/pathgetterinternal"
	"gitlab.com/evatix-go/pathhelper/knowndirstructure"
	"gitlab.com/evatix-go/pathhelper/pathsconst"
)

type NginxDir struct {
	Root, User            *knowndirstructure.NginxApacheDirectory
	SpecificUserRoot      string // /etc/nginx/conf.d/users/userName
	AllUsersRoot          string // /etc/nginx/conf.d/users
	CurrentUserRootConfig string // /etc/nginx/conf.d/users/userName/username.conf
}

func (it NginxDir) MkDir(
	isCreateRootDir,
	isCreateSpecificUserDirOnly,
	isCreateBackup bool,
	mode os.FileMode,
) *errwrappers.Collection {
	errCollection := errwrappers.Empty()

	if isCreateRootDir {
		errCollection.AddCollections(it.Root.MkDirAll(mode))
	}

	if !isCreateSpecificUserDirOnly {
		errCollection.AddWrapperPtr(createdirinternal.AllRecurse(it.AllUsersRoot, mode))
	}

	errCollection.AddWrapperPtr(createdirinternal.AllRecurse(it.SpecificUserRoot, mode))
	errCollection.AddWrapperPtr(it.User.MkDirSitesAvailable(mode))
	errCollection.AddWrapperPtr(it.User.MkDirSitesEnabled(mode))

	if isCreateBackup {
		errCollection.AddWrapperPtr(it.User.MkDirSitesBackup(mode))
	}

	return errCollection
}

func (it NginxDir) MkDirDefault(
	isCreateRootDir,
	isCreateSpecificUserDirOnly bool,
) *errwrappers.Collection {
	return it.MkDir(
		isCreateRootDir,
		isCreateSpecificUserDirOnly,
		false,
		DefaultDirChmod)
}

func (it NginxDir) MkDirDefaultSpecific() *errwrappers.Collection {
	return it.MkDir(
		false,
		true,
		false,
		DefaultDirChmod)
}

func (it NginxDir) MkDirSpecificUserRoot(mode os.FileMode) *errorwrapper.Wrapper {
	return createdirinternal.AllRecurse(it.SpecificUserRoot, mode)
}

func (it NginxDir) MkDirAllUsersRoot(mode os.FileMode) *errorwrapper.Wrapper {
	return createdirinternal.AllRecurse(it.AllUsersRoot, mode)
}

func (it *NginxDir) AllFilesAtSpecificUserRoot() *errstr.Results {
	return pathgetterinternal.GetAllFiles(
		true,
		osconsts.PathSeparator,
		it.SpecificUserRoot)
}

func (it *NginxDir) CurrentUserRootConfigName() string {
	_, fileName := filepath.Split(it.CurrentUserRootConfig)

	return fileName
}

// UsersConfigGlob gives the /etc/nginx/conf.d/users/*.conf
func (it *NginxDir) UsersConfigGlob() string {
	return normalizeinternal.JoinFixIf(
		true,
		it.AllUsersRoot,
		WildcardDotConfFilter)
}

// UsersEnabledGlobConf
//
// gives something like : /etc/nginx/conf.d/users/username/sites-enabled/*.conf
func (it *NginxDir) UsersEnabledGlobConf() string {
	return normalizeinternal.JoinFixIf(
		true,
		it.UserSitesEnableDir(),
		WildcardDotConfFilter)
}

// UsersAvailableSitesGlobConf
//
// gives something like : /etc/nginx/conf.d/users/username/sites-available/*.conf
func (it *NginxDir) UsersAvailableSitesGlobConf() string {
	return normalizeinternal.JoinFixIf(
		true,
		it.UserSitesAvailableDir(),
		WildcardDotConfFilter)
}

// UserSitesAvailableDir
//
// Example : /etc/nginx/conf.d/users/username/sites-available
func (it *NginxDir) UserSitesAvailableDir() string {
	return it.User.SitesAvailable
}

// UserSitesEnableDir
//
// Example : /etc/nginx/conf.d/users/username/sites-enabled
func (it *NginxDir) UserSitesEnableDir() string {
	return it.User.SitesEnabled
}

func (it *NginxDir) SiteNameAddConfExt(siteName string) string {
	return siteName + extensionsconst.DotConf
}

func (it *NginxDir) AbsPathOfAvailableSiteAddConfExt(siteName string) string {
	return normalizeinternal.JoinFixIf(
		true,
		it.UserSitesAvailableDir(),
		it.SiteNameAddConfExt(siteName))
}

func (it *NginxDir) CopyCurrentUserRootConfigTo(newLocation string) *errorwrapper.Wrapper {
	return fsinternal.CopyFile(
		it.CurrentUserRootConfig,
		newLocation,
		DefaultDirChmod)
}

// CopyCurrentUserRootConfigToTempRel
//
// Returns final copied path
func (it *NginxDir) CopyCurrentUserRootConfigToTempRel(tempRelativePath string) *errstr.Result {
	finalPath := normalizeinternal.JoinFixIf(
		true,
		pathsconst.TempDir,
		tempRelativePath+osconsts.PathSeparator+it.CurrentUserRootConfigName())

	return errstr.New.Result.Create(
		finalPath,
		it.CopyCurrentUserRootConfigTo(finalPath))
}

// AbsPathOfAvailableSite
//
// siteName : site-name
func (it *NginxDir) AbsPathOfAvailableSite(siteName string) string {
	return normalizeinternal.JoinFixIf(
		true,
		it.UserSitesAvailableDir(),
		siteName)
}

// AbsPathOfEnabledSite
//
// site : site-name
func (it *NginxDir) AbsPathOfEnabledSite(siteName string) string {
	return normalizeinternal.JoinFixIf(
		true,
		it.UserSitesEnableDir(),
		siteName)
}

func (it *NginxDir) AbsPathOfEnabledSiteAddConfExt(siteName string) string {
	return normalizeinternal.JoinFixIf(
		true,
		it.UserSitesEnableDir(),
		it.SiteNameAddConfExt(siteName))
}

// DisableSite disable a specific file, by removing symbolic link
//  TODO: Disabling server will be complicated, if multiple servers
//   present in same site file
func (it *NginxDir) DisableSite(siteName string) *errorwrapper.Wrapper {
	absPathOfEnabledSite := it.AbsPathOfEnabledSite(siteName)

	return fsinternal.SafeRemove(absPathOfEnabledSite)
}

func (it *NginxDir) DisableAllSites() *errwrappers.Collection {
	errC := errwrappers.Empty()
	for _, enabledSitePath := range it.AllUserEnabledSitesNoError() {
		if errW := fsinternal.SafeRemove(enabledSitePath); errW.HasError() {
			errC.AddWrapperPtr(errW)
		}
	}

	return errC
}

func (it *NginxDir) EnableSiteAddConfExt(siteName string) *errorwrapper.Wrapper {
	return it.EnableSite(it.SiteNameAddConfExt(siteName))
}

func (it *NginxDir) EnableSite(siteName string) *errorwrapper.Wrapper {
	destinationSiteFilePath := it.AbsPathOfEnabledSite(siteName)
	errW := fsinternal.CreateDirectoryAllUptoParent(
		destinationSiteFilePath,
		DefaultDirChmod)

	if errW.HasError() {
		return errW
	}

	sourceSitePath := it.AbsPathOfAvailableSite(siteName)
	err := os.Symlink(sourceSitePath, destinationSiteFilePath)

	return errnew.SrcDst.Error(
		errtype.SymbolicLink,
		err,
		sourceSitePath,
		destinationSiteFilePath)
}

func (it *NginxDir) RemoveAllUserSiteFiles() *errorwrapper.Wrapper {
	return fsinternal.
		SafeRemove(it.UserSitesAvailableDir()).
		ConcatNew().
		Wrapper(
			fsinternal.SafeRemove(
				it.UserSitesEnableDir()))
}

func (it *NginxDir) AllSiteFiles() ([]string, error) {
	return filepath.Glob(it.UsersAvailableSitesGlobConf())
}

func (it *NginxDir) AllSiteFilesNoError() []string {
	sites, _ := filepath.Glob(it.UsersAvailableSitesGlobConf())

	return sites
}

func (it *NginxDir) AllUserEnabledSites() ([]string, error) {
	return filepath.Glob(it.UsersEnabledGlobConf())
}

func (it *NginxDir) AllUserEnabledSitesNoError() []string {
	sites, _ := filepath.Glob(it.UsersEnabledGlobConf())

	return sites
}

func (it *NginxDir) HasSiteFile(site string) bool {
	return fsinternal.IsPathExists(it.AbsPathOfAvailableSite(site))
}

func (it *NginxDir) HasSiteFileAddConfExt(site string) bool {
	return it.HasSiteFile(it.SiteNameAddConfExt(site))
}

func (it *NginxDir) CombinedSpecificUserRoot(
	combinedPaths ...string,
) (
	first string,
	allCombinedPaths []string,
) {
	return normalizeinternal.PathsCombine(
		it.SpecificUserRoot,
		combinedPaths)
}

func (it *NginxDir) CombinedAllUsersRoot(
	combinedPaths ...string,
) (
	first string,
	allCombinedPaths []string,
) {
	return normalizeinternal.PathsCombine(
		it.AllUsersRoot,
		combinedPaths)
}

func (it *NginxDir) AllFilesAtAllUsersRoot() *errstr.Results {
	return pathgetterinternal.GetAllFiles(
		true,
		osconsts.PathSeparator,
		it.AllUsersRoot)
}

func (it *NginxDir) UsersAvailableSites() *errstr.Results {
	return pathgetterinternal.GetAllFiles(
		true,
		osconsts.PathSeparator,
		it.UserSitesAvailableDir())
}

func (it *NginxDir) AllUsersDirs() *errstr.Results {
	return pathgetterinternal.GetAllDirectories(
		true,
		osconsts.PathSeparator,
		it.AllUsersRoot)
}

func (it *NginxDir) AllUsersNames() *errstr.Results {
	results := pathgetterinternal.GetAllDirectories(
		true,
		osconsts.PathSeparator,
		it.AllUsersRoot)

	if results.HasIssuesOrEmpty() {
		return errstr.New.Results.ErrorWrapper(results.ErrorWrapper)
	}

	fileNames := make([]string, results.Length())

	for i, filePath := range results.SafeValues() {
		fileInfo, err := os.Stat(filePath)

		if err != nil {
			return errstr.New.Results.ErrorWrapper(
				errnew.
					Path.
					Error(errtype.InvalidPath, err, filePath),
			)
		}

		if fileInfo == nil {
			return errstr.New.Results.ErrorWrapper(
				errnew.
					Path.
					Messages(errtype.InvalidPath, filePath, "fileinfo nil"),
			)
		}

		fileNames[i] = fileInfo.Name()
	}

	return errstr.New.Results.Strings(fileNames)
}

func (it *NginxDir) UsersEnabledSites() *errstr.Results {
	return pathgetterinternal.GetAllFiles(
		true,
		osconsts.PathSeparator,
		it.UserSitesEnableDir())
}

func (it NginxDir) Json() corejson.Result {
	return corejson.NewFromAny(it)
}

func (it NginxDir) JsonPtr() *corejson.Result {
	return corejson.NewFromAnyPtr(it)
}

func (it NginxDir) JsonString() string {
	return corejson.NewFromAnyPtr(it).JsonString()
}

func (it NginxDir) JsonModelAny() interface{} {
	return it
}

func (it *NginxDir) JsonParseSelfInject(jsonResult *corejson.Result) error {
	err := jsonResult.Unmarshal(it)

	return err
}

func (it *NginxDir) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}

func (it *NginxDir) AsJsoner() corejson.Jsoner {
	return it
}
