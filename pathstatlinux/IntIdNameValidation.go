package pathstatlinux

import "gitlab.com/auk-go/core/constants"

type IntIdNameValidation struct {
	Id         int
	Name       string
	HasValidId bool
}

func InvalidIntIdNameValidation() *IntIdNameValidation {
	return &IntIdNameValidation{
		Id:         constants.InvalidValue,
		Name:       "",
		HasValidId: false,
	}
}
