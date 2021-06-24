package pathsysinfo

import (
	"fmt"
	"os/user"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper"
)

type GroupInfo struct {
	*user.Group
	Id           int
	IsValidGroup bool
	HasValidId   bool
	ErrorWrapper *errorwrapper.Wrapper
}

func InvalidGroupInfo(errorWrapper *errorwrapper.Wrapper) *GroupInfo {
	return &GroupInfo{
		Id:           constants.InvalidValue,
		IsValidGroup: false,
		HasValidId:   false,
		ErrorWrapper: errorWrapper,
	}
}

func (receiver *GroupInfo) String() string {
	return fmt.Sprintf(printFormat,
		receiver.Name,
		receiver.Id,
		receiver.IsValidGroup,
		receiver.HasValidId)
}
