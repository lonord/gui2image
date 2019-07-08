package gui2image

import (
	"image"
	"image/draw"
)

// ImageView is the view display image
type ImageView struct {
	Control
	Img image.Image
}

// Image implements method of View
func (v *ImageView) Image() image.Image {
	v.checkDefault()
	img := image.NewRGBA(v.Bounds)
	// fill background
	draw.Draw(img, img.Bounds(), image.NewUniform(v.Background), image.ZP, draw.Over)
	// fill image
	bSize := v.Bounds.Size()
	sSize := v.Img.Bounds().Size()
	px, py := (bSize.X-sSize.X)/2, (bSize.Y-sSize.Y)/2
	draw.Draw(img, img.Bounds(), v.Img, v.Img.Bounds().Min.Add(image.Pt(-px, -py)), draw.Over)
	return img
}

func (v *ImageView) checkDefault() {
	v.Control.checkDefault()
}
