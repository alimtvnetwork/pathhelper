package oldtests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/auk-go/core/coreutils/stringutil"
	"gitlab.com/auk-go/core/ostype"
)

type generalizedPathWithoutInputTestCaseDataWrapper struct {
	expected, operatingSystemMessage, funcName string
	osUserName                                 string
	operatingSystem                            ostype.Variation
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

		if len(testCase.osUserName) == 0 {
			// Assert
			Convey(GetAssertMessage(actual, testCase.expected, i), func() {
				So(actual, ShouldNotBeNil)
				So(actual, ShouldEqual, testCase.expected)
			})

			return
		}

		leftRightExpected := stringutil.SplitLeftRightTypeTrimmed(testCase.expected, testCase.osUserName)

		// Assert
		Convey(GetAssertMessage(actual, testCase.expected, i), func() {
			So(actual, ShouldNotBeNil)
			So(actual, ShouldStartWith, leftRightExpected.Left)
			So(actual, ShouldEndWith, leftRightExpected.Right)
		})
	})
}
