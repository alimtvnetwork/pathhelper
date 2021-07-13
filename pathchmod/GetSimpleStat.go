package pathchmod

import (
	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func GetSimpleStat(
	location string,
) *SimpleStat {
	info, isExist, err := chmodhelper.GetPathExistStatExpand(
		location)

	if err != nil {
		pathErr := errnew.Path(
			errtype.InvalidPath,
			err,
			location)

		return &SimpleStat{
			Location:        location,
			FileInfo:        info,
			HasFileInfo:     info != nil,
			InvalidFileInfo: info == nil,
			IsNotExist:      true,
			IsExist:         false,
			IsDir:           false,
			IsFile:          false,
			ErrWrapper:      pathErr,
		}
	}

	return &SimpleStat{
		Location:        location,
		FileInfo:        info,
		HasFileInfo:     info != nil,
		InvalidFileInfo: info == nil,
		IsNotExist:      !isExist,
		IsExist:         isExist,
		IsDir:           isExist && info != nil && info.IsDir(),
		IsFile:          isExist && info != nil && !info.IsDir(),
		ErrWrapper:      errnew.EmptyPtr,
	}
}
