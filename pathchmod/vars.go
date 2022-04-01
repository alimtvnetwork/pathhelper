package pathchmod

import "gitlab.com/evatix-go/core/chmodhelper"

var (
	rwxCreator = chmodhelper.New.RwxWrapper.UsingFileMode
)
