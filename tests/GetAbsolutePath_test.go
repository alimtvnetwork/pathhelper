package tests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
)

type getAbsolutePathTestCaseWrapper struct{
	inputBasepath, inputRelativePath, expected, expectedMessage, operatingSystem  string
}

var getAbsoltePathTestCaseWrappers = []getAbsolutePathTestCaseWrapper{
	// {
	// 	inputBasepath: "",
	// 	inputRelativePath: "",
	// 	expected: "",
	// 	expectedMessage: "empty return",
	// 	operatingSystem:  "For any OS",
	// },
	{
		inputBasepath: "c:\\Windows\\",
		inputRelativePath: "\\whatever",
		expected: "c:\\Windows\\whatever",
		expectedMessage: "non-empty return of (c:\\Windows\\whatever)",
		operatingSystem:  "OS is windows",
	},
	{
		inputBasepath: "c:\\Windows\\",
		inputRelativePath: "whatever",
		expected: "c:\\Windows\\whatever",
		expectedMessage: "non-empty return of (c:\\Windows\\whatever)",
		operatingSystem:  "OS is windows",
	},
	{
		inputBasepath: "/home/your_user_name/my_script/",
		inputRelativePath: "/whatever",
		expected: "/home/your_user_name/my_script/whatever",
		expectedMessage: "non-empty return of (/home/your_user_name/my_script/whatever)",
		operatingSystem:  "OS other than windows",
	},
	{
		inputBasepath: "/home/your_user_name/my_script",
		inputRelativePath: "/whatever",
		expected: "/home/your_user_name/my_script/whatever",
		expectedMessage: "non-empty return of (/home/your_user_name/my_script/whatever)",
		operatingSystem:  "OS other than windows",
	},
}

func TestGetAbsolutePath(t *testing.T) {
	for _, testCase := range getAbsoltePathTestCaseWrappers{
		// Arrange
		testCaseMessage := fmt.Sprintf("(%s) [GetAbsolutePath] inputs (%s, %s) expects (%s)", testCase.operatingSystem, testCase.inputBasepath, testCase.inputRelativePath, testCase.expectedMessage)

		Convey(testCaseMessage, t, func(){
			// Act
			actual := pathhelper.GetAbsolutePath(testCase.inputBasepath, testCase.inputRelativePath)

			// Assert
			if pathhelper.IsWindows() {
				So(actual[1], ShouldEqual, testCase.expected[1])
				So(actual[2], ShouldEqual, testCase.expected[2])
			}

			if !pathhelper.IsWindows() {
				So(actual[3], ShouldEqual, testCase.expected[3])
				So(actual[4], ShouldEqual, testCase.expected[4])
			}
		})
	}
	// Convey("If function is run", t, func() {
	// 	Convey("if not windows", func() {
	// 		So(pathhelper.IsWindows(), ShouldBeTrue)
	// 		Convey("it should return absolute path", func() {
	// 			So(pathhelper.GetAbsolutePath("c:\\Windows\\", "\\whatever"), ShouldEqual, "c:\\Windows\\whatever")
	// 			So(pathhelper.GetAbsolutePath("c:\\Windows\\", "whatever"), ShouldEqual, "c:\\Windows\\whatever")
	// 			So(pathhelper.GetAbsolutePath("c:\\Windows", "\\whatever"), ShouldEqual, "c:\\Windows\\whatever")
	// 		})
	// 	})
	// })
	//
	// Convey("If function is run", t, func() {
	// 	Convey("if not windows", func() {
	// 		So(pathhelper.IsWindows(), ShouldBeTrue)
	// 		Convey("it should return absolute path", func() {
	// 			So(pathhelper.GetAbsolutePath("/home/your_user_name/my_script/", "/whatever"), ShouldNotEqual, "/home/your_user_name/my_script/whatever")
	// 			So(pathhelper.GetAbsolutePath("/home/your_user_name/my_script/", "whatever"), ShouldNotEqual, "/home/your_user_name/my_script/whatever")
	// 			So(pathhelper.GetAbsolutePath("/home/your_user_name/my_script", "/whatever"), ShouldNotEqual, "/home/your_user_name/my_script/whatever")
	// 		})
	//
	// 	})
	// })

}
