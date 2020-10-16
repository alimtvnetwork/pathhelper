package enums

type KnownDirectory string

const (
	AppData          KnownDirectory = KnownDirectory("AppData")
	Bin              KnownDirectory = KnownDirectory("bin")
	BinUnix          KnownDirectory = KnownDirectory("/usr/bin")
	Documents        KnownDirectory = KnownDirectory("Documents")
	Downloads        KnownDirectory = KnownDirectory("Downloads")
	Etc              KnownDirectory = KnownDirectory("etc")
	GitGlobalWin     KnownDirectory = KnownDirectory(".gitconfig")
	GitGlobalUnix    KnownDirectory = KnownDirectory("/etc/gitconfig")
	GitGlobalUnixXdg KnownDirectory = KnownDirectory("XDG_CONFIG_HOME/git/config")
	Local            KnownDirectory = KnownDirectory("Local")
	LocalTempWin     KnownDirectory = KnownDirectory("local\\temp")
	LocalTempUnix    KnownDirectory = KnownDirectory("temp")
	Music            KnownDirectory = KnownDirectory("Music")
	Pictures         KnownDirectory = KnownDirectory("Pictures")
	ProgramFiles32   KnownDirectory = KnownDirectory("Program Files")
	ProgramFiles64   KnownDirectory = KnownDirectory("Program Files (x86)")
	ProgramData      KnownDirectory = KnownDirectory("Program Data")
	Roaming          KnownDirectory = KnownDirectory("Roaming")
	SSHGlobal        KnownDirectory = KnownDirectory(".ssh")
	System           KnownDirectory = KnownDirectory("System")
	System32         KnownDirectory = KnownDirectory("System32")
	SystemUnix       KnownDirectory = KnownDirectory("/etc/systemd/system")
	Temp             KnownDirectory = KnownDirectory("Temp")
	TempDir          KnownDirectory = KnownDirectory("TMPDIR")
	User             KnownDirectory = KnownDirectory("User")
	UserBin          KnownDirectory = KnownDirectory("UserBin")
	Users            KnownDirectory = KnownDirectory("Users")
	Videos           KnownDirectory = KnownDirectory("Videos")
	WindowsDirectory KnownDirectory = KnownDirectory("windir")
	WindowsCDrive    KnownDirectory = KnownDirectory("c:\\")
	UnixRoot         KnownDirectory = KnownDirectory("/")
)

func (directory KnownDirectory) Value() string {
	return string(directory)
}

func (directory KnownDirectory) ValuePtr() *string {
	value := directory.Value()
	return &value
}
