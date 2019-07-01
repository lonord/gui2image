package gui2image

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"io"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestRenderPaper(t *testing.T) {
	parent := &Paper{Background: color.RGBA{255, 255, 255, 255}, Bounds: image.Rect(0, 0, 200, 200)}
	child := &Paper{Background: color.RGBA{0, 0, 0, 255}, Bounds: image.Rect(10, 10, 100, 100)}
	parent.AddSub(child)
	img := parent.Image()
	var b bytes.Buffer
	png.Encode(&b, img)
	matchFile(t, "TestRenderPaper.png", b.Bytes())
}

func matchFile(t *testing.T, name string, b []byte) {
	path := filepath.Join("testdata", name)
	fileBytes, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if len(b) != len(fileBytes) {
		f := outputMismatchFile(b, name)
		t.Fatal("file length mismatch, actual", f)
	}
	for i := 0; i < len(b); i++ {
		if b[i] != fileBytes[i] {
			f := outputMismatchFile(b, name)
			t.Fatal("file content mismatch at position", i, "actual", f)
		}
	}
}

func outputMismatchFile(b []byte, name string) string {
	tmpfile, _ := ioutil.TempFile("", "*."+name)
	path := tmpfile.Name()
	defer tmpfile.Close()
	io.Copy(tmpfile, bytes.NewBuffer(b))
	return path
}
