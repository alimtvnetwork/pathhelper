package pathstatlinux

import (
	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
)

type RwxSimple struct {
	HyphenedRwxValue string
	RwxWrapper       *chmodhelper.RwxWrapper
	ErrorWrapper     *errorwrapper.Wrapper
	IsRwxValid       bool
}

func InvalidRwxSimple() *RwxSimple {
	return &RwxSimple{
		HyphenedRwxValue: constants.EmptyString,
		RwxWrapper:       nil,
		ErrorWrapper:     errnew.Type.Create(errtype.ChmodInvalid),
		IsRwxValid:       false,
	}
}

// func (receiver *RwxSimple) ApplyOnLinuxPaths()  {
// 	// return receiver.RwxWrapper.clo
// }
//
// func (receiver *RwxSimple) ApplyOnLinuxPaths()  {
//
// }
