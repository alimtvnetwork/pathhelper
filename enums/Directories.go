package enums

import (
	"strings"

	"gitlab.com/evatix-go/pathhelper/constants"
)

type KnownDirectory string

const (
	AppData          KnownDirectory = "AppData"
	Bin              KnownDirectory = "bin"
	BinUnix          KnownDirectory = "/usr/bin"
	Documents        KnownDirectory = "Documents"
	Downloads        KnownDirectory = "Downloads"
	Etc              KnownDirectory = "etc"
	GitGlobalWin     KnownDirectory = ".gitconfig"
	GitGlobalUnix    KnownDirectory = "/etc/gitconfig"
	GitGlobalUnixXdg KnownDirectory = "XDG_CONFIG_HOME/git/config"
	Local            KnownDirectory = "Local"
	LocalTempWin     KnownDirectory = "local\\temp"
	LocalTempUnix    KnownDirectory = "temp"
	Music            KnownDirectory = "Music"
	Pictures         KnownDirectory = "Pictures"
	ProgramFiles32   KnownDirectory = "Program Files"
	ProgramFiles64   KnownDirectory = "Program Files (x86)"
	ProgramData      KnownDirectory = "Program Data"
	Roaming          KnownDirectory = "Roaming"
	SSHGlobal        KnownDirectory = ".ssh"
	System           KnownDirectory = "System"
	System32         KnownDirectory = "System32"
	SystemUnix       KnownDirectory = "/etc/systemd/system"
	Temp             KnownDirectory = "Temp"
	TempDir          KnownDirectory = "TMPDIR"
	User             KnownDirectory = "User"
	UserBin          KnownDirectory = "UserBin"
	Users            KnownDirectory = "Users"
	Videos           KnownDirectory = "Videos"
	WindowsDirectory KnownDirectory = "windir"
	WindowsCDrive    KnownDirectory = "c:\\"
	UnixRoot         KnownDirectory = "/"
	NginxLinuxPath   KnownDirectory = "/etc/nginx/"
)

func (directory KnownDirectory) Value() string {
	return string(directory)
}

func (directory KnownDirectory) CombineWith(paths ...string) string {
	paths = append(paths, directory.Value())

	return strings.Join(paths, constants.PathSeparator)
}

func (directory KnownDirectory) ValuePtr() *string {
	value := directory.Value()

	return &value
}
