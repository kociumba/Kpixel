package main

import (
	"image/color"
	"math"
)

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

func getRed(c color.RGBA) float64 {
	return float64(c.R)
}

func getGreen(c color.RGBA) float64 {
	return float64(c.G)
}

func getBlue(c color.RGBA) float64 {
	return float64(c.B)
}
