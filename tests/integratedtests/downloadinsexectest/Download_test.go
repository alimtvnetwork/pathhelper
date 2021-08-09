package downloadinsexectest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper/fs"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
	"gitlab.com/evatix-go/pathhelper/pathinsfmtexec/downloadinsexec"
)

func Test_Download(t *testing.T) {
	// 0. Setup
	tempFile, buff := createTempFile(t)

	// spin up the server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, tempFile)
	}))

	defer ts.Close()

	download := pathinsfmt.NewDownload(ts.URL, filePath)
	errW := downloadinsexec.Apply(download)

	convey.Convey("Download ErrorWrapper Should Return False", t, func() {
		convey.So(errW.HasError(), convey.ShouldBeFalse)
	})

	convey.Convey("Download Content Should Resemble Temp Content", t, func() {
		errBytesResults := fs.ReadFile(filePath)
		convey.So(errBytesResults.HasError(), convey.ShouldBeFalse)
		convey.So(*errBytesResults.Values, convey.ShouldResemble, buff)
	})
}

