package tests

import (
	"fmt"
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errconv"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"

	"gitlab.com/evatix-go/pathhelper"
	"gitlab.com/evatix-go/pathhelper/enums"
	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
)

var (
	expectedNewEmptyDirectoryResult = pathhelpercore.DirectoryResult{
		FileInfoWrapper:   nil,
		Error:             errnew.Empty,
		RawPath:           "",
		FileModeRequested: nil,
		HasIssues:         false,
		IsIgnoredAction:   true,
		Action:            enums.EmptyDirectoryResult,
	}

	errorWrapperType = reflect.TypeOf(errorwrapper.Wrapper{})
)

func TestNewEmptyDirectoryResult(t *testing.T) {
	// Arrange
	testMessage := fmt.Sprint("[NewEmptyDirectoryResult] expects pointer to Directory result struct")

	Convey(testMessage, t, func() {
		// Act
		actual := pathhelpercore.NewEmptyDirectoryResult()
		expectedReflect := reflect.ValueOf(expectedNewEmptyDirectoryResult)

		// Assert
		So(*actual, ShouldHaveSameTypeAs, expectedNewEmptyDirectoryResult)

		actualValueOf := reflect.ValueOf(*actual)
		for i := 0; i < actualValueOf.NumField(); i++ {
			actualFieldValue := GetFieldValue(actualValueOf.Field(i))
			expectedFieldValue := GetFieldValue(expectedReflect.Field(i))

			// https://play.golang.org/p/2fEwolio_lY
			if i == 1 {
				AssertErrorWrapperEqual(actualFieldValue, expectedFieldValue, i)

				continue
			}

			Convey(pathhelper.GetAssertMessage(actualFieldValue, expectedFieldValue, i), func() {
				So(actualFieldValue, ShouldEqual, expectedFieldValue)
			})
		}
	})
}

func AssertErrorWrapperEqual(err1, err2 interface{}, index int) {
	if reflect.TypeOf(err1) != errorWrapperType {
		errtype.
			UnexpectedType.
			PanicNoRefs("error wrapper type is not matching.")
	}

	errW1 := errconv.Get(err1)
	errW2 := errconv.GetPtr(err2)

	Convey(pathhelper.GetAssertMessage(err1, err2, index), func() {
		So(errW1.Wrapper.IsEquals(errW2.Wrapper), ShouldBeTrue)
	})
}
