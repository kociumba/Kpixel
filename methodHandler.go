package main

import (
	"fmt"
	"image"
)

// SortingMethod represents different methods of sorting an image
type SortingMethod int

const (
	ColumnSorting SortingMethod = iota
	RowSorting    SortingMethod = iota
	RandomSorting SortingMethod = iota
)

// Define color sorting types
type ColorSortType int

const (
	Hue   ColorSortType = iota
	Red   ColorSortType = iota
	Green ColorSortType = iota
	Blue  ColorSortType = iota
)

// SortingStrategy defines the behavior for sorting methods
type SortingStrategy interface {
	SortImage(img image.Image) *image.RGBA
}

// ColumnSort sorts an image by column
type ColumnSort struct {
	colorSortType ColorSortType
}

// RowSort sorts an image by row
type RowSort struct {
	colorSortType ColorSortType
}

// RandomSort sorts an image randomly in chunks
type RandomSort struct {
	colorSortType ColorSortType
}

func (cs *ColumnSort) SortImage(img image.Image) *image.RGBA {
	return sortPixels(img, ColumnSorting, cs.colorSortType)
}

func (rs *RowSort) SortImage(img image.Image) *image.RGBA {
	return sortPixels(img, RowSorting, rs.colorSortType)
}

func (rs *RandomSort) SortImage(img image.Image) *image.RGBA {
	return sortPixels(img, RandomSorting, rs.colorSortType)
}

func setupSorting(inputMethod string, inputColorSort string) (SortingStrategy, error) {
	colorSortType, err := parseSortByValue(inputColorSort)
	if err != nil {
		return nil, err
	}

	var strategy SortingStrategy
	switch inputMethod {
	case "column":
		strategy = &ColumnSort{colorSortType: colorSortType}
	case "row":
		strategy = &RowSort{colorSortType: colorSortType}
	case "random":
		strategy = &RandomSort{colorSortType: colorSortType}
	default:
		return nil, fmt.Errorf("invalid sorting method: %s", inputMethod)
	}

	return strategy, nil
}

// func parseSortingMethod(input string) (SortingStrategy, error) {
// 	switch input {
// 	case "column":
// 		return &ColumnSort{}, nil
// 	case "row":
// 		return &RowSort{}, nil
// 	case "random":
// 		return &RandomSort{}, nil
// 	default:
// 		return nil, fmt.Errorf("invalid sorting method: %s", input)
// 	}
// }

func parseSortByValue(input string) (ColorSortType, error) {
	switch input {
	case "hue":
		return Hue, nil
	case "red":
		return Red, nil
	case "green":
		return Green, nil
	case "blue":
		return Blue, nil
	default:
		return 0, fmt.Errorf("invalid sort value: %s", input)
	}
}
