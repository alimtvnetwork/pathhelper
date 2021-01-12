package pathhelper

import "gitlab.com/evatix-go/core/constants"

func GetProgramFiles() string {
	if getOSArchitecture() == constants.Architecture32 {
		return GetProgramFiles32()
	}

	return GetProgramFiles64()
}
