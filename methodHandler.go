package main

import (
	"fmt"
	"image"
)

var chunkSize *int

// SortingMethod represents different methods of sorting an image
type SortingMethod int

const (
	ColumnSorting SortingMethod = iota
	RowSorting    SortingMethod = iota
	RandomSorting SortingMethod = iota
)

// SortingStrategy defines the behavior for sorting methods
type SortingStrategy interface {
	SortImage(img image.Image) *image.RGBA
}

// ColumnSort sorts the image by columns
type ColumnSort struct{}

// RowSort sorts the image by rows
type RowSort struct{}

// RandomSort sorts the image randomly
type RandomSort struct{}

func (cs *ColumnSort) SortImage(img image.Image) *image.RGBA {
	return sortPixels(img, ColumnSorting)
}

func (rs *RowSort) SortImage(img image.Image) *image.RGBA {
	return sortPixels(img, RowSorting)
}

func (rs *RandomSort) SortImage(img image.Image) *image.RGBA {
	return sortPixels(img, RandomSorting)
}

func parseSortingMethod(input string) (SortingStrategy, error) {
	switch input {
	case "column":
		return &ColumnSort{}, nil
	case "row":
		return &RowSort{}, nil
	case "random":
		return &RandomSort{}, nil
	default:
		return nil, fmt.Errorf("invalid sorting method: %s", input)
	}
}
