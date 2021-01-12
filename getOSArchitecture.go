package pathhelper

import (
	"fmt"
	"runtime"

	"gitlab.com/evatix-go/core/constants"
)

var x32Architectures = []string{
	"386", "arm", "armbe", "mips", "amd64p32", "mips64p32", "mips64p32le", "ppc", "riscv", "s390", "sparc",
}
var x64Architectures = []string{
	"amd64", "arm64", "ppc64", "ppc64le", "mips64", "mips64le", "riscv64", "s390x", "wasm", "arm64be", "sparc64",
}

func getOSArchitecture() string {
	arch := runtime.GOARCH

	if isStringsContains(x64Architectures, arch) {
		return constants.Architecture64
	}

	if isStringsContains(x32Architectures, arch) {
		return constants.Architecture32
	}

	message := fmt.Sprintf("Operating System Not supported, Os = %s!", arch)
	panic(message)
}
