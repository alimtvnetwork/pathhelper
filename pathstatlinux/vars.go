package pathstatlinux

import (
	"gitlab.com/auk-go/pathhelper/internal/deferrwrappers"
	"gitlab.com/auk-go/pathhelper/pathsysinfo"
)

var (
	invalidGroupInfo = pathsysinfo.InvalidGroupInfo(deferrwrappers.InvalidSystemGroup)
	invalidUserInfo  = pathsysinfo.InvalidUserInfo(deferrwrappers.InvalidSystemUser)
)
