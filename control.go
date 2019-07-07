package gui2image

import (
	"image"
	"image/color"
	"image/draw"
)

// Control is the base view
type Control struct {
	Background color.Color
	Bounds     image.Rectangle
}

// Image implements method of View
func (p *Control) Image() image.Image {
	p.checkDefault()
	img := image.NewRGBA(p.Bounds)
	// fill background
	draw.Draw(img, img.Bounds(), image.NewUniform(p.Background), image.ZP, draw.Src)
	return img
}

func (p *Control) checkDefault() {
	if p.Background == nil {
		p.Background = color.White
	}
}
