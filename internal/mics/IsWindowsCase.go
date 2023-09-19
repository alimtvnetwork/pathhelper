package mics

import "gitlab.com/auk-go/core/ostype"

func IsWindowsCase(os ostype.Variation) bool {
	return os == ostype.Windows
}
