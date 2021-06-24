package consts

import "gitlab.com/evatix-go/core/filemode"

const (
	FilePathEmpty                    = "File path was empty(\"\")."
	BrokenLongPathUncPrefix          = `\?\UNC\`
	BrokenLongPathQuestionMarkPrefix = `\?\`
	DoubleStars                      = "**"
	Export                           = "export"
	SystemCTL                        = "systemctl"
	Service                          = "service"
	Apt                              = "apt"
	AptGet                           = "apt-get"
	Upgrade                          = "upgrade"
	HyphenC                          = "-c"
	Start                            = "start"
	Restart                          = "restart"
	Stop                             = "stop"
	Reload                           = "reload"
	Enable                           = "enable"
	Ln                               = "ln"
	HyphenS                          = "-s"
	Install                          = "install"
	AptUpdate                        = "apt update -y"
	AptGetUpdate                     = "apt-get update -y"
	AptUpgrade                       = "apt upgrade -y"
	PPAWithColon                     = "ppa:"
	AddAptRepository                 = "add-apt-repository -y"
	WGet                             = "wget"
	RmRf                             = "rm -rf"
	HyphenY                          = "-y"
	DefaultFileMode                  = filemode.X644
	DefaultDirectoryFileMode         = filemode.X666
)
