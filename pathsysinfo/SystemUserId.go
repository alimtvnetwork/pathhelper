package pathsysinfo

import (
	"os/user"
	"strconv"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
)

func SystemUserId(userResult *user.User) (int, *errorwrapper.Wrapper) {
	uid, errUidConvert := strconv.Atoi(userResult.Uid)

	if errUidConvert != nil {
		return constants.InvalidValue, errnew.Type.Error(errtype.ConversionValueToInteger, errUidConvert)
	}

	return uid, nil
}
