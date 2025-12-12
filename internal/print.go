package internal

import (
	"fmt"
	"image"
)

func PrintImage(img image.Image) {
	rect := img.Bounds()
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			fmt.Println(r, g, b, a)
		}
	}
}
