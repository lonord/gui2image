package gui2image

import (
	"image"
	"image/color"
	"image/draw"
)

// Paper is the base view
type Paper struct {
	Background color.Color
	Bounds     image.Rectangle

	// private fields
	sub []View
}

// Image implements method of View
func (p *Paper) Image() image.Image {
	img := image.NewRGBA(p.Bounds)
	// fill background
	draw.Draw(img, img.Bounds(), image.NewUniform(p.Background), image.ZP, draw.Src)
	// draw sub views
	min := p.Bounds.Min
	for _, sv := range p.sub {
		simg := sv.Image()
		draw.Draw(img, simg.Bounds().Add(min), simg, simg.Bounds().Min, draw.Src)
	}
	return img
}

// AddSub add sub view to self
func (p *Paper) AddSub(sub View) {
	p.sub = append(p.sub, sub)
}

// RemoveSub remove sub view
func (p *Paper) RemoveSub(sub View) {
	if p.sub == nil {
		return
	}
	di := -1
	for i, s := range p.sub {
		if s == sub {
			di = i
			break
		}
	}
	if di != -1 {
		p.sub = append(p.sub[:di], p.sub[di+1:]...)
	}
}

func (p *Paper) checkDefault() {
	if p.Background == nil {
		p.Background = color.White
	}
}
