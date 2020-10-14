package enums

type OperatingSystem int

const (
	Any OperatingSystem = iota
	Debian
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
