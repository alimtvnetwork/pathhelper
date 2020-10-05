package tests

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
)

type pathExistsTestCaseWrapper struct {
	OperatingSystem string
	Input           string
	Expected        bool
}

var isPathExistsTestCases = []pathExistsTestCaseWrapper{
	{
		OperatingSystem: "Windows",
		Input:           "c:\\sampleSomething",
		Expected:        false,
	},
	{
		OperatingSystem: "Windows",
		Input:           "c:\\windows",
		Expected:        true,
	},
	{
		OperatingSystem: "Windows",
		Input:           "c:\\windows\\etc",
		Expected:        false,
	},
	{
		OperatingSystem: "Unix",
		Input:           "c:\\windows\\etc",
		Expected:        false,
	},
}

func TestIsPathExists(t *testing.T) {
	for _, testCase := range isPathExistsTestCases {
		convey.Convey("If (IsPathExists) function is run", t, func() {
			testMethodName := fmt.Sprintf("%s should return %s", testCase.Input, strconv.FormatBool(testCase.Expected))

			convey.Convey(testMethodName, func() {
				// Act
				actual := pathhelper.IsPathExist(testCase.Input)

				// Asset
				convey.So(actual, convey.ShouldEqual, testCase.Expected)
			})
		})
	}
}
