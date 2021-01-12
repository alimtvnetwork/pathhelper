package tests

import (
	"fmt"
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
)

var (
	expectedNewEmptyDirectoryResult = pathhelpercore.DirectoryResult{
		FileInfoWrapper:   nil,
		Error:             nil,
		RawPath:           "",
		FileModeRequested: nil,
		HasIssues:         false,
		IsIgnoredAction:   true,
		Action:            enums.EmptyDirectoryResult,
	}
)

func TestNewEmptyDirectoryResult(t *testing.T) {
	// Arrange
	testMessage := fmt.Sprint("[NewEmptyDirectoryResult] expects pointer to Directory result struct")

	Convey(testMessage, t, func() {
		// Act
		actual := pathhelpercore.NewEmptyDirectoryResult()

		// Assert
		So(*actual, ShouldHaveSameTypeAs, expectedNewEmptyDirectoryResult)

		actualValueOf := reflect.ValueOf(*actual)
		for i := 0; i < actualValueOf.NumField(); i++ {
			actualFieldValue := GetFieldValue(actualValueOf.Field(i))
			expectedFieldValue := GetFieldValue(reflect.ValueOf(expectedNewEmptyDirectoryResult).Field(i))

			Convey(pathhelper.GetAssertMessage(actualFieldValue, expectedFieldValue, i), func() {
				So(actualFieldValue, ShouldEqual, expectedFieldValue)
			})
		}
	})
}
