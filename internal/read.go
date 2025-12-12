package internal

import (
	"image"
	"image/color"
	_ "image/jpeg"
	"os"
)

func ReadImage(imgPath string) (image.Image, error) {
	file, err := os.Open(imgPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return castToRGBA8(img), nil
}

func castToRGBA8(img image.Image) image.Image {
	rect := img.Bounds()
	newImg := image.NewRGBA(rect)
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			r8, g8, b8, a8 := convToRGBA8(r, g, b, a)
			newImg.Set(x, y, color.RGBA{uint8(r8), uint8(g8), uint8(b8), uint8(a8)})
		}
	}
	return newImg
}

func convToRGBA8(r, g, b, a uint32) (uint8, uint8, uint8, uint8) {
	return uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), upperBound8
}
