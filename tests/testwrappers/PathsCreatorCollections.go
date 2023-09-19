package testwrappers

import (
	"gitlab.com/auk-go/pathhelper/pathinsfmt"
)

var PathsCreateInstructionsUnix = []*pathinsfmt.PathsCreatorCollection{
	{
		PathsCreatorItems: []pathinsfmt.PathsCreator{
			{
				BasePathsCreator: pathinsfmt.BasePathsCreator{
					RootDir:     RootPath1,
					Files:       FilesCollection1,
					IsNormalize: true,
				},
				ApplyRwx:       DefaultRwxOwnerGroupOther,
				ApplyUserGroup: DefaultUserNameGroupName,
			},
			{
				BasePathsCreator: pathinsfmt.BasePathsCreator{
					RootDir:     RootPath2,
					Files:       FilesCollection1,
					IsNormalize: true,
				},
				ApplyRwx:       DefaultRwxOwnerGroupOther,
				ApplyUserGroup: DefaultUserNameGroupName,
			},
			{
				BasePathsCreator: pathinsfmt.BasePathsCreator{
					RootDir:     RootPath3,
					Files:       FilesCollection1,
					IsNormalize: true,
				},
				ApplyRwx:       DefaultRwxOwnerGroupOther,
				ApplyUserGroup: DefaultUserNameGroupName,
			},
		},
		IsIgnoreOnExist:         false,
		IsDeleteAllBeforeCreate: true,
	},
}
