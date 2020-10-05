package tests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
)

func TestRemovingDoubleBackSlash(t *testing.T) {
	Convey("If given argument is  \"c:\\win\"", t, func() {
		Convey("it should return \"c:\\win\"", func() {
			So(pathhelper.RemovingDoubleBackSlash("c:\\\\win\\\\"), ShouldEqual, "c:\\win\\")
		})
	})
}
