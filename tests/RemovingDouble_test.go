package tests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
)

func TestRemovingDouble(t *testing.T) {

	Convey("If given argument is \"home//user\" or \"c:\\win\"", t, func() {

		Convey("it should return \"home/user\" or \"c:\\win\"", func() {
			So(pathhelper.RemovingDouble("home//user"), ShouldEqual, "home/user")
			So(pathhelper.RemovingDouble("c:\\\\win\\\\"), ShouldEqual, "c:\\win\\")
		})

	})

}
