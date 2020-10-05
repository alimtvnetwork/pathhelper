package tests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
)

func TestNormalizePath(t *testing.T) {

	Convey("If function is run", t, func() {

		Convey("it should return absolute path", func() {
			So(pathhelper.NormalizePath("c:\\windows//system32\\//etc"), ShouldEqual, "c:\\windows\\system32\\etc")
		})

	})

}
