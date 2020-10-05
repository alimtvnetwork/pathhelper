package tests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
)

func TestRemovingDoubleSlash(t *testing.T) {
	Convey("If given argument is  \"home//user\"", t, func() {
		Convey("it should return \"home/user\"", func() {
			So(pathhelper.RemovingDoubleSlash("home//user"), ShouldEqual, "home/user")
		})
	})
}
