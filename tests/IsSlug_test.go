package tests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
)

func TestIsSlug(t *testing.T) {

	Convey("given string doesn't contain anything from forbidden array", t, func() {

		Convey("it should return true", func() {
			So(pathhelper.IsSlug("xyz_"), ShouldEqual, true)
			So(pathhelper.IsSlug("xyz-"), ShouldEqual, true)
			So(pathhelper.IsSlug("xyz_"), ShouldEqual, true)
		})

	})

	Convey("if string contains anything from forbidden array", t, func() {

		Convey("it should return false", func() {
			So(pathhelper.IsSlug("xyz_#"), ShouldEqual, false)
			So(pathhelper.IsSlug("%&^2093073070271 b21 2987$#&^^&$(*&$("), ShouldEqual, false)
			So(pathhelper.IsSlug("xyz*(}"), ShouldEqual, false)
		})

	})

}
