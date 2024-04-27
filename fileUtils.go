package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

func openAndDecodeImage(imgPath string) (image.Image, string, error) {
	file, err := os.Open(imgPath)
	if err != nil {
		return nil, "", err
	}
	defer file.Close()

	switch filepath.Ext(imgPath) {
	case ".jpg", ".jpeg":
		img, err := jpeg.Decode(file)
		return img, "jpeg", err
	case ".png":
		img, err := png.Decode(file)
		return img, "png", err
	default:
		return nil, "", fmt.Errorf("unsupported file format")
	}
}

func saveImage(outputPath string, img image.Image, format string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	switch format {
	case "jpeg":
		return jpeg.Encode(file, img, nil)
	case "png":
		return png.Encode(file, img)
	default:
		return fmt.Errorf("unsupported image format")
	}
}
