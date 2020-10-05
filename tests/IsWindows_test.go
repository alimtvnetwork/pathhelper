package tests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
)

func TestIsWindows(t *testing.T) {

	Convey("given OS is windows", t, func() {

		Convey("it should return true", func() {
			So(pathhelper.IsWindows(), ShouldBeTrue)
		})

	})

	Convey("if OS is not windows", t, func() {

		Convey("it should return false", func() {
			So(!pathhelper.IsWindows(), ShouldBeFalse)
		})

	})

}
