package internal

import "image"

type mirrorImage struct {
}

func (mirrorImage) Process(img image.Image) (image.Image, error) {
	return nil, nil
}

func (mirrorImage) OutPrefix() string {
	return "mirrored"
}
