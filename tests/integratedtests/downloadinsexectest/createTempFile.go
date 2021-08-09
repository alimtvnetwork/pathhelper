package downloadinsexectest

import (
	"io/ioutil"
	"math/rand"
	"testing"
)

func createTempFile(t *testing.T) (string, []byte) {
	dir := t.TempDir()

	tempFile, err := ioutil.TempFile(dir, Test2MB)
	if err != nil {
		t.Fatal("creating temp file:", err)
	}

	buff := make([]byte, Size2MB)
	rand.Read(buff)
	if _, err := tempFile.Write(buff); err != nil {
		t.Fatal("write to temp file:", err)
	}

	defer tempFile.Close()
	// tempFile
	return tempFile.Name(), buff
}
