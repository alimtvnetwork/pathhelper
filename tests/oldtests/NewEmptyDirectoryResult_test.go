package oldtests

import (
	"fmt"
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errconv"
	"gitlab.com/auk-go/errorwrapper/errtype"

	"gitlab.com/auk-go/pathhelper/dirinfo"
)

var (
	expectedNewEmptyDirectoryResult = dirinfo.Empty()
	errorWrapperType                = reflect.TypeOf(errorwrapper.Wrapper{})
)

func TestNewEmptyDirectoryResult(t *testing.T) {
	// Arrange
	testMessage := fmt.Sprint("[Empty] expects pointer to Directory result struct")

	Convey(testMessage, t, func() {
		// Act
		actual := dirinfo.Empty()
		expectedReflect := reflect.ValueOf(*expectedNewEmptyDirectoryResult)

		// Assert
		So(*actual, ShouldHaveSameTypeAs, *expectedNewEmptyDirectoryResult)

		actualValueOf := reflect.ValueOf(*actual)
		for i := 0; i < actualValueOf.NumField(); i++ {
			actualFieldValue := GetFieldValue(actualValueOf.Field(i))
			expectedFieldValue := GetFieldValue(expectedReflect.Field(i))

			Convey(GetAssertMessage(actualFieldValue, expectedFieldValue, i), func() {
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

	Convey(GetAssertMessage(err1, err2, index), func() {
		So(errW1.Wrapper.IsEquals(errW2.Wrapper), ShouldBeTrue)
	})
}
