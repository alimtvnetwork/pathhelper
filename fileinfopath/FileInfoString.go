package fileinfopath

import (
	"os"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/pathhelper/internal/consts"
)

func FileInfoString(fileInfo os.FileInfo) string {
	if fileInfo == nil {
		return constants.EmptyString
	}

	return consts.IndentFileInfoEachLineJoiner + errcore.VarNameValuesJoiner(
		consts.IndentFileInfoEachLineJoiner,
		errcore.NameVal{
			Name:  "Name",
			Value: fileInfo.Name(),
		},
		errcore.NameVal{
			Name:  "Size",
			Value: fileInfo.Size(),
		},
		errcore.NameVal{
			Name:  "LastModifiedDate",
			Value: fileInfo.ModTime(),
		},
		errcore.NameVal{
			Name:  "IsDir",
			Value: fileInfo.IsDir(),
		})
}
