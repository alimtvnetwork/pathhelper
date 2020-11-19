package tests

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type pathTestCaseDataWrapper struct {
	OSName, funcName, expected string
}

func pathTestCaseInternal_linux(t *testing.T, testData pathTestCaseDataWrapper, actualFuncCall func() string) {
	// Arrange
	SkipOnWindows(t)

	testMessage := fmt.Sprintf("(%s) [%s] inputs() expects string output", testData.OSName, testData.funcName)

	Convey(testMessage, t, func() {
		// Act
		actual := actualFuncCall()

		// Assert
		So(actual, ShouldEqual, testData.expected)
	})
}

func pathTestCaseInternalFromWrappers(t *testing.T, testData []pathTestCaseDataWrapper, actualFuncCall func() string) {
	for _, testCase := range testData {
		fmt.Println(testCase)
		pathTestCaseInternal_linux(t, testCase, actualFuncCall)
	}
}
