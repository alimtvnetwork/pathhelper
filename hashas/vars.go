package hashas

import (
	"gitlab.com/evatix-go/core/coredata/coredynamic"
	"gitlab.com/evatix-go/core/coreimpl/enumimpl"
)

var (
	ranges = [...]string{
		Undefined: "Undefined",
		Md5:       "Md5",
		Sha1:      "Sha1",
		Sha256:    "Sha256",
		Sha512:    "Sha512",
	}

	DefaultFastHashMethod   = Md5
	DefaultSecureHashMethod = Sha1

	BasicEnumImpl = enumimpl.NewBasicByteUsingIndexedSlice(
		coredynamic.TypeName(Undefined),
		ranges[:])
)
