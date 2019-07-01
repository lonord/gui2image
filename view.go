package gui2image

import "image"

// View is the interface of all gui control
type View interface {
	// Image render gui control tree to image
	Image() image.Image
}
