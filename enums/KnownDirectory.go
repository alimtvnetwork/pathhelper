package enums

import (
	"gitlab.com/evatix-go/pathhelper/constants"
	"path"
	"strings"
)

type KnownDirectory string

const (
	ApacheLinuxPath  KnownDirectory = "/etc/apache/"
	AppData          KnownDirectory = "AppData"
	AppDataUnix      KnownDirectory = "/usr/share"
	Bin              KnownDirectory = "bin"
	BinUnix          KnownDirectory = "/usr/bin"
	Documents        KnownDirectory = "Documents"
	Downloads        KnownDirectory = "Downloads"
	Drivers          KnownDirectory = "drivers"
	DriversUnix      KnownDirectory = "/lib/modules/$(uname -r)/kernel/drivers/"
	Etc              KnownDirectory = "etc"
	Fonts            KnownDirectory = "Fonts"
	FontsUnix        KnownDirectory = "usr/share/fonts"
	GitGlobalWin     KnownDirectory = ".gitconfig"
	GitGlobalUnix    KnownDirectory = "/etc/gitconfig"
	GitGlobalUnixXdg KnownDirectory = "XDG_CONFIG_HOME/git/config"
	HostFile         KnownDirectory = "hosts"
	Local            KnownDirectory = "Local"
	LocalTempWin     KnownDirectory = "local\\temp"
	LocalTempUnix    KnownDirectory = "tmp"
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
	WindowsCDrive    KnownDirectory = "C:\\"

	// for paths of nginx and  apache
	Conf             KnownDirectory = "conf.d"
	ConfAvailable    KnownDirectory = "conf-available"
	ConfEnabled      KnownDirectory = "conf-enabled"
	ModsAvailable    KnownDirectory = "mods-available"
	ModsEnabled      KnownDirectory = "mods-enabled"
	ModulesAvailable KnownDirectory = "modules-available"
	ModulesEnabled   KnownDirectory = "modules-enabled"
	SitesAvailable   KnownDirectory = "sites-available"
	SitesEnabled     KnownDirectory = "sites-enabled"
	MimeTypes        KnownDirectory = "mime.types"
)

func (directory KnownDirectory) Value() string {
	return string(directory)
}

// directory.Value() + constants.PathSeparator + knownDirectories.join(constants.PathSeparator)
// Warning: It doesn't perform complex tasks like long path normalize, long path (windows) fix, double separator to single and so on.
func (directory KnownDirectory) CombineWithKnownDirs(knownDirectories ...KnownDirectory) string {
	paths := make([]string, 0, len(knownDirectories)+2)
	paths = append(paths, directory.Value())

	for _, knownDirectory := range knownDirectories {
		paths = append(paths, knownDirectory.Value())
	}

	return path.Clean(strings.Join(paths, constants.PathSeparator))
}

func (directory KnownDirectory) CombineWith(paths ...string) string {
	paths = append(paths, directory.Value())

	return path.Clean(strings.Join(paths, constants.PathSeparator))
}

// KnownDirectory.Value() + constants.PathSeparator + paths with separator
// Warning: It doesn't perform complex tasks like long path normalize, long path (windows) fix, double separator to single and so on.
func (directory KnownDirectory) GetPrefixCombinedWith(paths ...string) string {
	paths = append([]string{directory.Value()}, paths...)

	return path.Clean(strings.Join(paths, constants.PathSeparator))
}

func (directory KnownDirectory) ValuePtr() *string {
	value := directory.Value()

	return &value
}
