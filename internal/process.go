package internal

import (
	"fmt"
	"image"
	"path/filepath"
)

type processor interface {
	Process(img image.Image) (image.Image, error)
}

type outPrefixer interface {
	OutPrefix() string
}

var processorMap = map[Action]processor{
	Invert: invertImage{},
	Print:  printImage{},
	Mirror: mirrorImage{},
}

func ProcessImage(imgPath string, action Action) error {
	processor, exist := processorMap[action]
	if !exist {
		return fmt.Errorf("unsupported action: %d", action)
	}
	img, err := ReadImage(imgPath)
	if err != nil {
		return fmt.Errorf("read image failed: %w", err)
	}
	img, err = processor.Process(img)
	if err != nil {
		return fmt.Errorf("process image failed: %w", err)
	}
	if img != nil {
		outPath := getOutPath(imgPath, processor)
		err = WriteImage(img, outPath)
		if err != nil {
			return fmt.Errorf("write image failed: %w", err)
		}
		if IndicateOutPath {
			fmt.Println(outPath)
		}
	}
	return nil
}

func getOutPath(imgPath string, maybePrefixer interface{}) string {
	prefix := ""
	if prefixer, ok := maybePrefixer.(outPrefixer); ok {
		prefix = prefixer.OutPrefix() + "-"
	}
	return filepath.Join(
		OutPath,
		prefix+filepath.Base(imgPath),
	)
}
