package pathinsfmt

import "gitlab.com/evatix-go/core/chmodhelper/chmodins"

type PathVerifier struct {
	BaseUserNamePlusGroupName
	chmodins.BaseRwxInstructions
	Path string `json:"Path"`
}
