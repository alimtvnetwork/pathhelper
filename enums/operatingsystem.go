package enums

type OperatingSystem int

const (
	Debian OperatingSystem = iota
	Linux
	// Darwin is Mac OS or iOS
	DarwinOrMacOrIOS
	Ubuntu
	Android
	Dragonfly
	FreeBSD
	Nacl
	OpenBSD
	Js
	NetBSD
	Solaris
	Windows
)
