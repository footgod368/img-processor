package internal

import (
	"image"
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
	return img, nil
}
