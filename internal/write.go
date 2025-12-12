package internal

import (
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
)

func WriteImage(img image.Image, path string) error {
	// 1. 提取文件路径的目录部分（比如 "output/imgs/test.jpg" → "output/imgs"）
	dir := filepath.Dir(path)

	// 2. 递归创建目录（mkdir -p），权限0755（通用目录权限）
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err // 目录创建失败则返回错误
	}
	
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	err = jpeg.Encode(file, img, nil)
	if err != nil {
		return err
	}
	return nil
}
