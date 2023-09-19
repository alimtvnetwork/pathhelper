package pathhelpercore

import "gitlab.com/auk-go/errorwrapper"

type BaseErrorWrapper struct {
	ErrorWrapper *errorwrapper.Wrapper `json:"ErrorWrapper,omitempty"`
}
