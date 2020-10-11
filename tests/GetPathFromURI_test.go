package tests

import (
	"testing"
	"fmt"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
)

type getPathFromUriTestCaseWrapper struct {
	input, expected, expectedMessage, operatingSystem string
	inputBool bool
}

var getPathFromUriTestCaseWrappers = []getPathFromUriTestCaseWrapper{
	{
		input: "file://c:/windows/users/etc/more",
		inputBool: true,
		expected: "c:\\windows\\users\\etc\\more",
		expectedMessage: "c:\\windows\\users\\etc\\more",
		operatingSystem: "OS is windows",
	},
	{
		input: "c:\\windows\\users\\etc\\more",
		inputBool: false,
		expected: "c:\\windows\\users\\etc\\more",
		expectedMessage: "c:\\windows\\users\\etc\\more",
		operatingSystem: "OS other than windows",
	},
}

func TestGetPathFromUri(t *testing.T) {
	// Arrange
	for _, testCase := range getPathFromUriTestCaseWrappers{
		testCaseMessage := fmt.Sprintf("[getPathFromURI] inputs (%s, %v) expects (%s)", testCase.input, testCase.inputBool, testCase.expectedMessage)

		Convey(testCaseMessage, t, func(){
			// Act
			actual := pathhelper.GetPathFromUri(testCase.input, testCase.inputBool)

			// Assert
			So(actual, ShouldEqual, testCase.expected)
		})
	}
	//
	// Convey("If given OS is windows", t, func() {
	// 	So(pathhelper.IsWindows(), ShouldBeTrue)
	// 	Convey("if GetPathFromUri is run", func() {
	//
	// 		Convey("it should return", func() {
	// 			So(pathhelper.GetPathFromUri("file://c:/windows/users/etc/more", true), ShouldEqual, "c:\\windows\\users\\etc\\more")
	// 		})
	//
	// 	})
	//
	// })
	//
	// Convey("if given OS is not windows", t, func() {
	//
	// 	Convey("it should return", func() {
	//
	// 		Convey("it should return", func() {
	// 			So(pathhelper.GetPathFromUri("file://c:/windows/users/etc/more",  true), ShouldNotEqual, "c:/windows/users/etc/more")
	// 		})
	//
	// 	})
	//
	// })
}
