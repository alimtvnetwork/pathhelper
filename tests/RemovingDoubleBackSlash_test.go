package tests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
)

func TestRemovingDoubleBackSlash(t *testing.T) {
	Convey("[RemovingDoubleBackSlash] with input (c:\\\\win\\\\) expects non-empty, non-nil return of (c:\\win\\)", t, func() {
		So(pathhelper.RemovingDoubleBackSlash("c:\\\\win\\\\"), ShouldNotBeEmpty)
		So(pathhelper.RemovingDoubleBackSlash("c:\\\\win\\\\"), ShouldNotBeNil)
		So(pathhelper.RemovingDoubleBackSlash("c:\\\\win\\\\"), ShouldEqual, "c:\\win\\")
	})
}
