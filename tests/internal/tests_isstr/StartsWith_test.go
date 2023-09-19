package tests_isstr

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/coreutils/stringutil"
)

func TestIsStartsWith(t *testing.T) {
	for i, testCase := range startsWithWrappersCases {
		// Arrange
		testCaseMessenger := testCase.AsTestCaseMessenger()
		testHeader := coretests.GetTestHeader(testCaseMessenger)

		// Act
		actual := stringutil.IsStartsWith(
			testCase.baseDir,
			testCase.search,
			testCase.isIgnoreCase)

		testCase.SetActual(actual)

		Convey(testHeader, t, func() {
			// Assert
			Convey(coretests.GetAssertMessage(testCaseMessenger, i), func() {
				So(actual, ShouldEqual, testCase.Expected())
			})
		})
	}
}
