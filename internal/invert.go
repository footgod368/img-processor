package internal

import (
	"image"
	"image/color"
	"path/filepath"
)

func InvertImage(imgPath string) error {
	img, err := ReadImage(imgPath)
	if err != nil {
		return err
	}
	img = invert(img)
	return WriteImage(img, convToOutPath(imgPath))
}

func invert(img image.Image) image.Image {
	rect := img.Bounds()
	newImg := image.NewRGBA(rect)
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			r8, g8, b8, a8 := convToRGBA8(r, g, b, a)
			r_ := upperBound8 - r8
			g_ := upperBound8 - g8
			b_ := upperBound8 - b8
			newImg.Set(x, y, color.RGBA{uint8(r_), uint8(g_), uint8(b_), uint8(a8)})
		}
	}
	return newImg
}

func convToRGBA8(r, g, b, a uint32) (uint8, uint8, uint8, uint8) {
	return uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), upperBound8
}

func convToOutPath(path string) string {
	return "out/inverted-" + filepath.Base(path)
}
