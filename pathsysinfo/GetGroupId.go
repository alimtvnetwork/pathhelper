package pathsysinfo

import (
	"os/user"
	"strconv"

	"gitlab.com/auk-go/core/codestack"
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errtype"
)

func GetGroupId(groupObj *user.Group) (int, *errorwrapper.Wrapper) {
	gid, errGidConvert := strconv.Atoi(groupObj.Gid)

	if errGidConvert != nil {
		return constants.InvalidValue, errorwrapper.NewRef(
			codestack.SkipNone,
			errtype.SearchFailed,
			errGidConvert,
			"GroupNameId",
			groupObj.Gid)
	}

	return gid, nil
}
