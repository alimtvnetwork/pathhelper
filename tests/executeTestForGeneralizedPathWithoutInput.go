package tests

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
	"testing"
)

type generalizedPathWithoutInputTestCaseDataWrapper struct {
	expected, operatingSystemMessage, funcName string
	operatingSystem                            enums.OperatingSystem
}

func executeTestForGeneralizedPathWithoutInput(
	t *testing.T,
	testCase generalizedPathWithoutInputTestCaseDataWrapper,
	funcCall func() string,
	i int,
) {
	testCaseMessage := fmt.Sprintf("(%s) [%s] expects (%s)", testCase.operatingSystemMessage, testCase.funcName, testCase.expected)

	Convey(testCaseMessage, t, func() {
		// Act
		actual := funcCall()

		// Assert
		Convey(pathhelper.GetAssertMessage(actual, testCase.expected, i), func() {
			So(actual, ShouldNotBeNil)
			So(actual, ShouldEqual, testCase.expected)
		})
	})
}
