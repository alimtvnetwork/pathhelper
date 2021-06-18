package deferrwrappers

import (
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

var (
	InvalidSystemUser                   = errnew.MessagesPtr(errtype.NotFound, "System user not found or id has issues.")
	InvalidSystemGroup                  = errnew.MessagesPtr(errtype.NotFound, "System group not found or id has issues.")
	CannotApplyChmodWithSingleParameter = errnew.MessagesPtr(
		errtype.SysGroupInvalid,
		"Cannot process chmod using single user or group. "+
			"Both (user + group) needs to be present. "+
			"Btw, we can provide only group name then it"+
			" will switch command to chgrp to apply group on paths...")
)
