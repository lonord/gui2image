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
	parent := &Paper{Control: Control{Background: color.RGBA{255, 255, 255, 255}, Bounds: image.Rect(0, 0, 200, 200)}}
	child := &Paper{Control: Control{Background: color.RGBA{0, 0, 0, 255}, Bounds: image.Rect(10, 10, 100, 100)}}
	parent.AddSub(child)
	img := parent.Image()
	var b bytes.Buffer
	png.Encode(&b, img)
	matchFile(t, "TestRenderPaper.png", b.Bytes())
}

func TestRenderPaperMulti(t *testing.T) {
	parent := &Paper{Control: Control{Background: color.RGBA{255, 255, 255, 255}, Bounds: image.Rect(0, 0, 200, 200)}}
	child := &Paper{Control: Control{Background: color.RGBA{0, 0, 0, 255}, Bounds: image.Rect(50, 50, 150, 150)}}
	child2 := &Paper{Control: Control{Background: color.RGBA{255, 0, 0, 255}, Bounds: image.Rect(25, 75, 125, 175)}}
	child3 := &Paper{Control: Control{Background: color.RGBA{0, 255, 0, 255}, Bounds: image.Rect(75, 25, 175, 125)}}
	parent.AddSub(child)
	parent.AddSub(child2)
	parent.AddSub(child3)
	img := parent.Image()
	var b bytes.Buffer
	png.Encode(&b, img)
	matchFile(t, "TestRenderPaperMulti.png", b.Bytes())
}

func TestRenderPaperTree(t *testing.T) {
	parent := &Paper{Control: Control{Background: color.RGBA{255, 255, 255, 255}, Bounds: image.Rect(0, 0, 200, 200)}}
	child := &Paper{Control: Control{Background: color.RGBA{0, 0, 0, 255}, Bounds: image.Rect(50, 50, 150, 150)}}
	child2 := &Paper{Control: Control{Background: color.RGBA{255, 0, 0, 255}, Bounds: image.Rect(5, 5, 80, 80)}}
	child3 := &Paper{Control: Control{Background: color.RGBA{0, 255, 0, 255}, Bounds: image.Rect(5, 5, 55, 55)}}
	parent.AddSub(child)
	child.AddSub(child2)
	child2.AddSub(child3)
	img := parent.Image()
	var b bytes.Buffer
	png.Encode(&b, img)
	matchFile(t, "TestRenderPaperTree.png", b.Bytes())
}

func TestRenderLabelDefault(t *testing.T) {
	label := &Label{
		Control:   Control{Background: color.RGBA{255, 255, 255, 255}, Bounds: image.Rect(0, 0, 200, 200)},
		Text:      "hello world",
		FontSize:  12,
		TextColor: color.Black,
	}
	img := label.Image()
	var b bytes.Buffer
	png.Encode(&b, img)
	matchFile(t, "TestRenderLabelDefault.png", b.Bytes())
}

func TestRenderLabelCenter(t *testing.T) {
	label := &Label{
		Control:   Control{Background: color.RGBA{255, 255, 255, 255}, Bounds: image.Rect(0, 0, 200, 200)},
		Text:      "hello world",
		FontSize:  12,
		TextColor: color.Black,
		HAlign:    AlignCenter,
		VAlign:    AlignCenter,
	}
	img := label.Image()
	var b bytes.Buffer
	png.Encode(&b, img)
	matchFile(t, "TestRenderLabelCenter.png", b.Bytes())
}

func TestRenderLabelEnd(t *testing.T) {
	label := &Label{
		Control:   Control{Background: color.RGBA{255, 255, 255, 255}, Bounds: image.Rect(0, 0, 200, 200)},
		Text:      "hello world",
		FontSize:  12,
		TextColor: color.Black,
		HAlign:    AlignEnd,
		VAlign:    AlignEnd,
	}
	img := label.Image()
	var b bytes.Buffer
	png.Encode(&b, img)
	matchFile(t, "TestRenderLabelEnd.png", b.Bytes())
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
