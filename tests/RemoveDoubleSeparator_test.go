package tests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
)

func TestRemoveDoubleSeparator(t *testing.T) {
	Convey("If given argument is \"__hello___\"", t, func() {
		Convey("it should return \"_hello_", func() {
			So(pathhelper.RemoveDoubleSeparator("__hello___", "_"), ShouldEqual, "_hello_")
			So(pathhelper.RemoveDoubleSeparator("__hello___", "-"), ShouldEqual, "-hello-_")
			So(pathhelper.RemoveDoubleSeparator("--hello---", "_"), ShouldEqual, "_hello_-")
			So(pathhelper.RemoveDoubleSeparator("--hello---", "-"), ShouldEqual, "-hello-")
			So(pathhelper.RemoveDoubleSeparator("-hello--", "-"), ShouldEqual, "-hello-")
		})
	})
}
