package tests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
)


func TestRemovingDoubleSlash(t *testing.T) {
	Convey("[RemovingDoubleSlash] with input (home//user) expects non-empty, non-nil return (home/user)", t, func() {
		So(pathhelper.RemovingDoubleSlash("home//user"), ShouldNotBeEmpty)
		So(pathhelper.RemovingDoubleSlash("home//user"), ShouldNotBeNil)
		So(pathhelper.RemovingDoubleSlash("home//user"), ShouldEqual, "home/user")
	})
}
