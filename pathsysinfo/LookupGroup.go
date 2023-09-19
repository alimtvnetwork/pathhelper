package pathsysinfo

import (
	"os/user"

	"gitlab.com/auk-go/core/codestack"
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errtype"
)

func LookupGroup(groupName string) (*user.Group, *errorwrapper.Wrapper) {
	groupObj, errLookupGroup := user.LookupGroup(groupName)

	if errLookupGroup != nil {
		return nil, errorwrapper.NewRef(
			codestack.SkipNone,
			errtype.SearchFailed,
			errLookupGroup,
			"GroupName",
			groupName)
	}

	return groupObj, nil
}
