package tests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
)

func TestGetPathAsUri_Windows(t *testing.T) {
	isWindows := pathhelper.IsWindows()
	Convey("If given OS is windows", t, func() {
		if !isWindows {
			t.Skip("Linux Machine: Skipping windows tests.")
		}

		Convey("if GetPathAsUri is run", func() {

			Convey("it should return", func() {
				// Act && Assert
				So(pathhelper.GetPathAsUri("c:\\windows", true), ShouldEqual, "file:///c:/windows")
				So(pathhelper.GetPathAsUri("c:/windows", true), ShouldEqual, "file:///c:/windows")
			})
		})
	})
}

func TestGetPathAsUri_Unix(t *testing.T) {
	isWindows := pathhelper.IsWindows()

	Convey("if given OS is not windows", t, func() {
		if isWindows {
			t.Skip("Windows Machine: Skipping unix tests.")
		}

		Convey("it should return", func() {
			Convey("it should return", func() {
				// Act && Assert
				So(pathhelper.GetPathAsUri("c:/windows", true), ShouldNotEqual, "file://c:/windows")
			})
		})
	})
}
