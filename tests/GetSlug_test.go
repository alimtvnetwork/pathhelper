package tests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
)

func TestGetSlug(t *testing.T) {

	Convey("If given isSlug(string) is true", t, func() {
		Convey("it should return nil", func() {
			So(pathhelper.GetSlug("_2093073070271-b21-2987_", "_"), ShouldEqual, "")
		})
	})

	Convey("if given isSlug(string) is false", t, func() {

		Convey("it should return string without forbidden array", func() {
			So(pathhelper.GetSlug("%&^2093073070271 b21 2987$#&^^&$(*&$(", "_"), ShouldEqual, "_2093073070271_b21_2987_")
			So(pathhelper.GetSlug("%&^2093073070271 b21 2987$#&^^&$(*&$(", "-"), ShouldEqual, "-2093073070271-b21-2987-")
			So(pathhelper.GetSlug("%&^209307307%271 b21 2^87$#&^^&$(*&$(", "_"), ShouldEqual, "_209307307_271_b21_2_87_")
			So(pathhelper.GetSlug("%&^2093*73070271 b21 2987$#&^^&$(*&$(", "-"), ShouldEqual, "-2093-73070271-b21-2987-")
			So(pathhelper.GetSlug("%&^20930__070271 b21 2987$#&^^&$(*&$(", "_"), ShouldEqual, "_20930_070271_b21_2987_")
		})

	})

}
