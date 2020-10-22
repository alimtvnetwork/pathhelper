package enums

import (
	"gitlab.com/evatix-go/pathhelper"
	"strings"

	"gitlab.com/evatix-go/pathhelper/constants"
)

type KnownDirectory string

const (
	AppData          KnownDirectory = "AppData"
	AppDataUnix          KnownDirectory = "/usr/share"
	Bin              KnownDirectory = "bin"
	BinUnix          KnownDirectory = "/usr/bin"
	Documents        KnownDirectory = "Documents"
	Downloads        KnownDirectory = "Downloads"
	Drivers          KnownDirectory = "drivers"
	DriversUnix      KnownDirectory = "/lib/modules/$(uname -r)/kernel/drivers/"
	Etc              KnownDirectory = "etc"
	Fonts            KnownDirectory = "fonts"
	FontsUnix        KnownDirectory = "usr/share/fonts"
	GitGlobalWin     KnownDirectory = ".gitconfig"
	GitGlobalUnix    KnownDirectory = "/etc/gitconfig"
	GitGlobalUnixXdg KnownDirectory = "XDG_CONFIG_HOME/git/config"
	HostFile         KnownDirectory = "hosts"
	Local            KnownDirectory = "Local"
	LocalTempWin     KnownDirectory = "local\\temp"
	LocalTempUnix    KnownDirectory = "temp"
	Music            KnownDirectory = "Music"
	NginxLinuxPath   KnownDirectory = "/etc/nginx/"
	Pictures         KnownDirectory = "Pictures"
	ProgramFiles32   KnownDirectory = "Program Files"
	ProgramFiles64   KnownDirectory = "Program Files (x86)"
	ProgramData      KnownDirectory = "Program Data"
	Roaming          KnownDirectory = "Roaming"
	Services         KnownDirectory = "services"
	SSHGlobal        KnownDirectory = ".ssh"
	System32         KnownDirectory = "System32"
	System64         KnownDirectory = "SysWOW64"
	SystemUnix       KnownDirectory = "/etc/systemd/system"
	Temp             KnownDirectory = "Temp"
	TempDir          KnownDirectory = "TMPDIR"
	UnixRoot         KnownDirectory = "/"
	User             KnownDirectory = "User"
	UserBin          KnownDirectory = "UserBin"
	Users            KnownDirectory = "Users"
	Videos           KnownDirectory = "Videos"
	WindowsDirectory KnownDirectory = "windir"
	WindowsCDrive    KnownDirectory = "c:\\"
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

func (directory KnownDirectory) GetPrefixCombinedWith(paths ...string) string {
	paths = append(paths, directory.Value())

	return pathhelper.GetCombinedPath(
		constants.PathSeparator,
		true,
		true,
		true,
		strings.Join(paths, constants.PathSeparator))
}

