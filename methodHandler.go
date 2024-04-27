package main

import (
	"fmt"
	"image"
)

// SortingMethod represents different methods of sorting an image
type SortingMethod int

const (
	ColumnSorting SortingMethod = iota
	RowSorting
)

// SortingStrategy defines the behavior for sorting methods
type SortingStrategy interface {
	SortImage(img image.Image) *image.RGBA
}

// ColumnSort sorts the image by columns
type ColumnSort struct{}

// RowSort sorts the image by rows
type RowSort struct{}

func (cs *ColumnSort) SortImage(img image.Image) *image.RGBA {
	return sortPixels(img, true)
}

func (rs *RowSort) SortImage(img image.Image) *image.RGBA {
	return sortPixels(img, false)
}

func parseSortingMethod(input string) (SortingStrategy, error) {
	switch input {
	case "column":
		return &ColumnSort{}, nil
	case "row":
		return &RowSort{}, nil
	default:
		return nil, fmt.Errorf("invalid sorting method: %s", input)
	}
}
