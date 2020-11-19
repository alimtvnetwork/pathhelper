package tests

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/pathhelper/pathhelpercore"
	"testing"
)

func TestFileInfoWrapper(t *testing.T) {
	// Arrange
	testMessage := fmt.Sprint("[FileInfoWrapper] methods expect boolean return")

	Convey(testMessage, t, func() {
		// Act
		actualNew := pathhelpercore.NewFileWrapperInfo("") //todo
		actualHasError := actualNew.HasError()
		actualPathExists := actualNew.IsPathExists()

		// Assert
		So(actualHasError, ShouldBeTrue)
		So(actualPathExists, ShouldBeFalse)
	})
}
