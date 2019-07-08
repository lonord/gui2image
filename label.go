package gui2image

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
)

var defaultFont *truetype.Font

// AlignType enum
const (
	AlignBegin AlignType = iota
	AlignCenter
	AlignEnd
)

// AlignType indicates align type
type AlignType int

// Label is the view display text
type Label struct {
	Control
	Text      string
	TextColor color.Color
	VAlign    AlignType
	HAlign    AlignType
	FontSize  float64
	Font      *truetype.Font
}

func init() {
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		return
	}
	defaultFont = font
}

// Image implements method of View
func (l *Label) Image() image.Image {
	l.checkDefault()
	img := image.NewRGBA(l.Bounds)
	// fill background
	draw.Draw(img, img.Bounds(), image.NewUniform(l.Background), image.ZP, draw.Over)
	// get text width and height
	img0 := image.NewRGBA(image.Rect(0, 0, (l.Bounds.Max.X-l.Bounds.Min.X)*2, (l.Bounds.Max.Y-l.Bounds.Min.Y)*2))
	width, height := l.draw(img0, 0, 0)
	// draw text
	px, py := 0, 0
	imgSize := l.Bounds.Size()
	if l.VAlign == AlignCenter {
		py = (imgSize.Y - height) / 2
	} else if l.VAlign == AlignEnd {
		py = imgSize.Y - height
	}
	if py < 0 {
		py = 0
	}
	if l.HAlign == AlignCenter {
		px = (imgSize.X - width) / 2
	} else if l.HAlign == AlignEnd {
		px = imgSize.X - width
	}
	if px < 0 {
		px = 0
	}
	l.draw(img, px+l.Bounds.Min.X, py+l.Bounds.Min.Y)
	return img
}

func (l *Label) draw(img draw.Image, px, py int) (int, int) {
	c := freetype.NewContext()
	c.SetDPI(72.0)
	c.SetFont(l.Font)
	c.SetFontSize(l.FontSize)
	c.SetClip(l.Bounds)
	c.SetDst(img)
	c.SetSrc(image.NewUniform(l.TextColor))
	c.SetHinting(font.HintingFull)
	fHeight := int(c.PointToFixed(l.FontSize) >> 6)
	pt := freetype.Pt(px, py+fHeight)
	ept, err := c.DrawString(l.Text, pt)
	if err != nil {
		panic(err)
	}
	ew := ept.X - pt.X
	return int(ew >> 6), fHeight
}

func (l *Label) checkDefault() {
	l.Control.checkDefault()
	if l.TextColor == nil {
		l.TextColor = color.Black
	}
	if l.FontSize == 0 {
		l.FontSize = 12
	}
	if l.Font == nil {
		l.Font = defaultFont
	}
}
