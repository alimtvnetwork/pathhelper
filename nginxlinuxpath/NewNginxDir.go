package nginxlinuxpath

import (
	"gitlab.com/evatix-go/core/coreasync"
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

	coreasync.Waited.ParallelVoidTasks(
		func() {
			nginxRoot = GetFullDirStructure(
				isNormalize,
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
				specificUserRoot)
		})

	return &NginxDir{
		Root:                  nginxRoot,
		User:                  userDir,
		AllUsersRoot:          allUsersRoot,
		SpecificUserRoot:      specificUserRoot,
		CurrentUserRootConfig: specificUserRootConfig,
	}
}
