package internal

import (
	"image"
	"image/color"
)

type invertImage struct {
}

func (invertImage) Process(img image.Image) (image.Image, error) {
	rect := img.Bounds()
	newImg := image.NewRGBA(rect)
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			r_ := uint32(upperBound8) - r
			g_ := uint32(upperBound8) - g
			b_ := uint32(upperBound8) - b
			newImg.Set(x, y, color.RGBA{uint8(r_), uint8(g_), uint8(b_), uint8(a)})
		}
	}
	return newImg, nil
}

func (invertImage) OutPrefix() string {
	return "inverted"
}
