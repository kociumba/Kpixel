package main

import (
	"image"
	"image/color"
	"math"
)

type Pixel struct {
	Hue        float64
	ColorValue color.RGBA
}

func sortPixels(img image.Image, sortByColumn bool) *image.RGBA {
	bounds := img.Bounds()
	newImg := image.NewRGBA(bounds)

	if sortByColumn {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			column := extractPixels(img, x, bounds.Min.Y, x+1, bounds.Max.Y)

			quickSort(column, 0, len(column)-1)

			for y, pixel := range column {
				newImg.Set(x, bounds.Min.Y+y, pixel.ColorValue)
			}

		}
	} else {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			row := extractPixels(img, bounds.Min.X, y, bounds.Max.X, y+1)

			quickSort(row, 0, len(row)-1)

			for x, pixel := range row {
				newImg.Set(bounds.Min.X+x, y, pixel.ColorValue)
			}

		}
	}

	return newImg
}

// extractPixels extracts pixels either as a row or column slice
func extractPixels(img image.Image, minX, minY, maxX, maxY int) []Pixel {
	pixels := make([]Pixel, 0, maxX-minX+maxY-minY)

	for y := minY; y < maxY; y++ {
		for x := minX; x < maxX; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			rgba := color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)}
			hue := calculateHue(rgba)
			pixels = append(pixels, Pixel{Hue: hue, ColorValue: rgba})
		}
	}

	return pixels
}

func quickSort(pixels []Pixel, start int, end int) {
	if start < end {
		pivotIndex := partition(pixels, start, end)
		quickSort(pixels, start, pivotIndex-1)
		quickSort(pixels, pivotIndex+1, end)
	}
}

func partition(pixels []Pixel, start int, end int) int {
	pivot := pixels[end].Hue
	i := start
	for j := start; j < end; j++ {
		if pixels[j].Hue < pivot {
			pixels[i], pixels[j] = pixels[j], pixels[i]
			i++
		}
	}
	pixels[i], pixels[end] = pixels[end], pixels[i]
	return i
}

func calculateHue(c color.RGBA) float64 {
	r := float64(c.R) / 255.0
	g := float64(c.G) / 255.0
	b := float64(c.B) / 255.0
	// pixelColor := gocolor.Color{R: float64(c.R), G: float64(c.G), B: float64(c.B)} // simpler way but with a dependency produces the same result
	// hue, _, _ := pixelColor.Hsl()
	max := math.Max(r, math.Max(g, b))
	min := math.Min(r, math.Min(g, b))
	var hue float64
	if max == min {
		hue = 0 // achromatic
	} else {
		switch max {
		case r:
			hue = (g - b) / (max - min)
		case g:
			hue = 2 + (b-r)/(max-min)
		case b:
			hue = 4 + (r-g)/(max-min)
		}
		hue *= 60
		if hue < 0 {
			hue += 360
		}
	}
	return hue
}
