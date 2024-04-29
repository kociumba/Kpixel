package main

import (
	"image"
	"image/color"
	"math/rand"
	"sort"
	"sync"
	"time"
)

type Pixel struct {
	SortValue  float64
	Hue        float64
	Red        float64
	Green      float64
	Blue       float64
	ColorValue color.RGBA
}

func sortPixels(img image.Image, sortingMethod SortingMethod, colorSortType ColorSortType) *image.RGBA {
	bounds := img.Bounds()
	newImg := image.NewRGBA(bounds)
	var wg sync.WaitGroup

	rand.New(rand.NewSource(time.Now().UnixNano())) // Seed random number generator

	switch sortingMethod {
	case ColumnSorting:
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			wg.Add(1)
			go func(x int) {
				defer wg.Done()
				column := extractPixels(img, x, bounds.Min.Y, x+1, bounds.Max.Y, colorSortType)
				sort.SliceStable(column, func(i, j int) bool {
					return column[i].SortValue < column[j].SortValue
				})
				for y, pixel := range column {
					newImg.Set(x, bounds.Min.Y+y, pixel.ColorValue)
				}
			}(x)
		}
	case RowSorting:
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			wg.Add(1)
			go func(y int) {
				defer wg.Done()
				row := extractPixels(img, bounds.Min.X, y, bounds.Max.X, y+1, colorSortType)
				sort.SliceStable(row, func(i, j int) bool {
					return row[i].SortValue < row[j].SortValue
				})
				for x, pixel := range row {
					newImg.Set(bounds.Min.X+x, y, pixel.ColorValue)
				}
			}(y)
		}
	case RandomSorting: // Assuming chunkSize is defined and valid. Adjust as needed.
		for x := bounds.Min.X; x < bounds.Max.X; x += *chunkSize {
			for y := bounds.Min.Y; y < bounds.Max.Y; y += *chunkSize {
				wg.Add(1)
				go func(x, y int) {
					defer wg.Done()
					// Determine the boundaries of the chunk
					endX := x + *chunkSize
					if endX > bounds.Max.X {
						endX = bounds.Max.X
					}
					endY := y + *chunkSize
					if endY > bounds.Max.Y {
						endY = bounds.Max.Y
					}

					// Extract all pixels in the current chunk
					pixels := make([]color.Color, 0, (endX-x)*(endY-y))
					for j := y; j < endY; j++ {
						for i := x; i < endX; i++ {
							pixels = append(pixels, img.At(i, j))
						}
					}

					// Shuffle the pixels in the current chunk
					rand.Shuffle(len(pixels), func(i, j int) {
						pixels[i], pixels[j] = pixels[j], pixels[i]
					})

					// Place the shuffled pixels back into the new image
					idx := 0
					for j := y; j < endY; j++ {
						for i := x; i < endX; i++ {
							newImg.Set(i, j, pixels[idx])
							idx++
						}
					}
				}(x, y)
			}
		}
	}

	wg.Wait()
	return newImg
}

// extractPixels extracts pixels either as a row or column slice
func extractPixels(img image.Image, minX, minY, maxX, maxY int, colorSortType ColorSortType) []Pixel {
	pixels := make([]Pixel, 0, maxX-minX+maxY-minY)
	var method float64

	for y := minY; y < maxY; y++ {
		for x := minX; x < maxX; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			rgba := color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)}

			switch colorSortType {
			case Hue:
				method = calculateHue(rgba)
			case Red:
				method = getRed(rgba)
			case Green:
				method = getGreen(rgba)
			case Blue:
				method = getBlue(rgba)
			case Luminosity:
				method = calculateLuminosity(rgba)
			case Saturation:
				method = calculateSaturation(rgba)
			}

			pixels = append(pixels, Pixel{SortValue: method, ColorValue: rgba})
		}
	}

	return pixels
}

// func quickSort(pixels []Pixel, start int, end int) {
// 	if start < end {
// 		pivotIndex := partition(pixels, start, end)
// 		quickSort(pixels, start, pivotIndex-1)
// 		quickSort(pixels, pivotIndex+1, end)
// 	}
// }

// func partition(pixels []Pixel, start int, end int) int {
// 	pivot := pixels[end].SortValue
// 	i := start
// 	for j := start; j < end; j++ {
// 		if pixels[j].SortValue < pivot {
// 			pixels[i], pixels[j] = pixels[j], pixels[i]
// 			i++
// 		}
// 	}
// 	pixels[i], pixels[end] = pixels[end], pixels[i]
// 	return i
// }
