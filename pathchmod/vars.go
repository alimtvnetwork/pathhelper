package pathchmod

import "gitlab.com/auk-go/core/chmodhelper"

var (
	FriendlyChmod = friendlyChmod{}
	rwxCreator    = chmodhelper.New.RwxWrapper.UsingFileMode
)
