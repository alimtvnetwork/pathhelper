package tests_isstr

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/coreutils/stringutil"
)

func TestIsEndsWith(t *testing.T) {
	for i, testCase := range endsWithWrappersCases {
		// Arrange
		testCaseMessenger := testCase.AsTestCaseMessenger()
		testHeader := coretests.GetTestHeader(testCaseMessenger)

		// Act
		actual := stringutil.IsEndsWith(
			testCase.baseDir,
			testCase.search,
			testCase.isIgnoreCase)

		testCase.SetActual(actual)

		convey.Convey(testHeader, t, func() {
			// Assert
			convey.Convey(coretests.GetAssertMessage(testCaseMessenger, i), func() {
				convey.So(actual, convey.ShouldEqual, testCase.Expected())
			})
		})
	}
}
