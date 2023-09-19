package pathcompilertests

import (
	"gitlab.com/auk-go/core/coreimpl/enumimpl"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/enum/osmixtype"
)

type TestWrapper struct {
	coretests.BaseTestCase
	SelectBy                      osmixtype.Variant
	NameAssert, DescriptionAssert string
	IsTestEnv                     bool
	RunsIn                        []osmixtype.Variant
}

func (it *TestWrapper) ExpectedAsDynamicMap() enumimpl.DynamicMap {
	conv, isSuccess := it.Expected().(enumimpl.DynamicMap)

	if !isSuccess {
		panic("expected type doesn't meet")
	}

	return conv
}
