package pathsysinfo

import (
	"os/user"

	"gitlab.com/auk-go/core/codestack"
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errtype"
)

func LookupUser(userName string) (userResult *user.User, errorWrapper *errorwrapper.Wrapper) {
	userObj, errLookup := user.Lookup(userName)

	if errLookup != nil {
		return nil, errorwrapper.NewRef(
			codestack.SkipNone,
			errtype.SearchFailed,
			errLookup,
			"UserName",
			userName)
	}

	return userObj, nil
}
