package internal

import (
	"image"
	"image/color"
)

type mirrorImage struct {
}

func (mirrorImage) Process(img image.Image) (image.Image, error) {
	rect := img.Bounds()
	newImg := image.NewRGBA(rect)
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for left, right := rect.Min.X, rect.Max.X; left < right; left, right = left+1, right-1 {
			rLeft, gLeft, bLeft, aLeft := img.At(left, y).RGBA()
			rRight, gRight, bRight, aRight := img.At(right, y).RGBA()
			newImg.Set(left, y, color.RGBA{uint8(rRight), uint8(gRight), uint8(bRight), uint8(aRight)})
			newImg.Set(right, y, color.RGBA{uint8(rLeft), uint8(gLeft), uint8(bLeft), uint8(aLeft)})
		}
	}
	return newImg, nil
}

func (mirrorImage) OutPrefix() string {
	return "mirrored"
}
