package hashas

import "gitlab.com/evatix-go/core/coreimpl/enumimpl"

var (
	ranges = [...]string{
		Undefined: "Undefined",
		Md5:       "Md5",
		Sha1:      "Sha1",
		Sha256:    "Sha256",
		Sha512:    "Sha512",
	}

	BasicEnumImpl = enumimpl.NewBasicByteUsingIndexedSlice(
		ranges[:])
)
