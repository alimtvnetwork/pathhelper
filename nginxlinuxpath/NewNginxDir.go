package nginxlinuxpath

import (
	"gitlab.com/evatix-go/asynchelper/syncparallel"
	"gitlab.com/evatix-go/core/coreinstruction"
	"gitlab.com/evatix-go/core/extensionsconst"
	"gitlab.com/evatix-go/pathhelper/internal/normalizeinternal"
	"gitlab.com/evatix-go/pathhelper/knowndirstructure"
)

func NewNginxDir(
	isNormalize bool,
	currentNginxRoot,
	username string,
) *NginxDir {
	var nginxRoot, userDir *knowndirstructure.NginxApacheDirectory
	var specificUserRoot, specificUserRootConfig string
	allUsersRoot := normalizeinternal.JoinPathsFixIf(
		isNormalize,
		currentNginxRoot,
		ExtraConfName,
		Users,
	)

	syncparallel.Tasks(
		func() {
			nginxRoot = GetFullDirStructure(
				isNormalize,
				DefaultDirChmod,
				currentNginxRoot)
		},
		func() {
			specificUserRoot = normalizeinternal.JoinPathsFixIf(
				isNormalize,
				allUsersRoot,
				username)

			specificUserRootConfig = normalizeinternal.JoinPathsFixIf(
				isNormalize,
				specificUserRoot,
				username+extensionsconst.DotConf,
			)

			userDir = GetFullDirStructure(
				isNormalize,
				DefaultDirChmod,
				specificUserRoot)
		})

	return &NginxDir{
		BaseUsername:          *coreinstruction.NewUsername(username),
		Root:                  nginxRoot,
		User:                  userDir,
		AllUsersRoot:          allUsersRoot,
		SpecificUserRoot:      specificUserRoot,
		CurrentUserRootConfig: specificUserRootConfig,
	}
}
