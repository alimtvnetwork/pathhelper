package fstestwrapper

import (
	"gitlab.com/auk-go/core/chmodhelper/chmodins"
	"gitlab.com/auk-go/pathhelper/pathinsfmt"
	"gitlab.com/auk-go/pathhelper/pathsconst"
)

var (
	SetupFiles = pathinsfmt.PathsCreator{
		BasePathsCreator: pathinsfmt.BasePathsCreator{
			RootDir: pathsconst.DefaultTempTestDir,
			Files: []string{
				setupFilePath,
			},
			IsNormalize: true,
		},
		ApplyRwx: &chmodins.RwxOwnerGroupOther{
			Owner: "rwx",
			Group: "rwx",
			Other: "rwx",
		},
		ApplyUserGroup: nil,
	}
)
