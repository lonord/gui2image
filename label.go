package gui2image

import (
	"image"
	"image/color"

	"github.com/golang/freetype/truetype"
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
	Paper
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
	// TODO
	return l.Paper.Image()
}

func (l *Label) checkDefault() {
	l.Paper.checkDefault()
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
