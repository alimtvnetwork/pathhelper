package tests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
)

type removeFromPathTestCaseWrapper struct {
	inputPath, expected, expectedMessage, operatingSystem string
	inputBool bool
}

var removingArray =  []string {"/"}

var removeFromPathTestCaseWrappers = []removeFromPathTestCaseWrapper{
	{
		inputPath: "",
		expected: "",
		expectedMessage: "empty return",
		operatingSystem:  "Any OS",
	},
	{
		inputPath: "c://win/etc/",
		inputBool: false,
		expected: "c:winetc",
		expectedMessage: "c:winetc",
		operatingSystem:  "Any OS",
	},
	{
		inputPath: "c:\\win\\etc",
		inputBool: true,
		expected: "c:\\win\\etc",
		expectedMessage: "c:\\win\\etc",
		operatingSystem:  "OS is windows",
	},
	{
		inputPath: "c:\\\\win\\\\etc",
		inputBool: true,
		expected: "c:/win/etc",
		expectedMessage: "c:/win/etc",
		operatingSystem:  "OS other than windows",
	},
}

func TestRemoveFromPath(t *testing.T){
	for _, testCase := range removeFromPathTestCaseWrappers {
		// Arrange
		testCaseMessage := fmt.Sprintf("(%s) [RemoveFromPath] inputs (%s, %v) expects (%s)", testCase.operatingSystem, testCase.inputPath, testCase.inputBool, testCase.expectedMessage)

		Convey(testCaseMessage, t, func() {
			// Act
			actual := pathhelper.RemoveFromPath(testCase.inputPath, &removingArray, testCase.inputBool)

			// Assert
			if pathhelpercore.IsEmptyPath(testCase.inputPath){
				So(actual, ShouldBeEmpty)
				So(actual[1], ShouldEqual, testCase.expected[1])
			}

			if !pathhelpercore.IsEmptyPath(testCase.inputPath) {
				if pathhelper.IsWindows() {
					So(actual[2], ShouldEqual, testCase.expected[2])
				}

				if !pathhelper.IsWindows() {
					So(actual[3], ShouldEqual, testCase.expected[3])
				}
			}
		})
	}
}
