package tests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
)

func TestHasPathIssues(t *testing.T) {

	Convey("If given string has prefix", t, func() {

		Convey("it should return true", func() {
			So(pathhelper.HasPathIssues("file:///C:/"), ShouldBeTrue)
		})

	})

	Convey("if given string has slash and backslash at the same time", t, func() {

		Convey("it should return true", func() {
			So(pathhelper.HasPathIssues("//C:\\win"), ShouldBeTrue)
		})

	})

	Convey("if given string has prefix and/or slash and backslash at the same time", t, func() {

		Convey("it should return true", func() {
			So(pathhelper.HasPathIssues("file:///C:\\win\\users"), ShouldBeTrue)
		})

	})

	Convey("if given string does not have prefix and/or slash and backslash at the same time", t, func() {

		Convey("it should return false", func() {
			So(pathhelper.HasPathIssues("C://windows/"), ShouldBeFalse)
		})

	})

}
