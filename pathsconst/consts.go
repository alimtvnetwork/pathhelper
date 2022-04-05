package pathsconst

const (
	AppName                 = "Cimux"
	AppNameLower            = "cimux"
	VarOpt                  = "/var/opt/"
	Etc                     = "/etc/"
	VarLog                  = "/var/log/"
	TestDirPatternName      = AppNameLower + "-tests"
	DefaultConfigRootSuffix = "/config"
	PackagesDirName         = "/packages/"
	InstructionDirName      = "/instructions/"
	UnixLogAppRoot          = VarLog + AppNameLower
	UnixVarAppRoot          = VarOpt + AppNameLower
	EtcApp                  = Etc + AppNameLower
	UnixConfigRoot          = EtcApp + DefaultConfigRootSuffix
)
