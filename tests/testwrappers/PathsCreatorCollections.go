package testwrappers

import (
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

var PathsCreateInstructionsUnix = []*pathinsfmt.PathsCreatorCollection{
	{
		PathsCreateInstructions: []pathinsfmt.BasePathsCreator{
			{
				RootDir:     RootPath1,
				Files:       FilesCollection1,
				IsNormalize: true,
			},
			{
				RootDir:     RootPath2,
				Files:       FilesCollection1,
				IsNormalize: true,
			},
			{
				RootDir:     RootPath3,
				Files:       FilesCollection1,
				IsNormalize: true,
			},
		},
		IsIgnoreOnExist:         false,
		IsDeleteAllBeforeCreate: true,
		ApplyRwx:                DefaultRwxOwnerGroupOther,
		ApplyUserGroup:          DefaultUserNameGroupName,
	},
}
