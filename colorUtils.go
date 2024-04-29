package main

import (
	"image/color"
	"math"

	clog "github.com/charmbracelet/log"
)

func getMaxMin(c color.RGBA) (float64, float64) {
	// r := float64(c.R) / 255.0
	// g := float64(c.G) / 255.0
	// b := float64(c.B) / 255.0

	max := math.Max(float64(c.R), math.Max(float64(c.G), float64(c.B)))
	min := math.Min(float64(c.R), math.Min(float64(c.G), float64(c.B)))

	return max, min
}

// (A) If R ≥ G ≥ B | H = 60° x [(GB)/(RB)]

// (B) If G > R ≥ B | H = 60° x [2 - (RB)/(GB)]

// (C) If G ≥ B > R | H = 60° x [2 + (BR)/(GR)]

// (D) If B > G > R | H = 60° x [4 - (GR)/(BR)]

// (E) If B > R ≥ G | H = 60° x [4 + (RG)/(BG)]

// (F) If R ≥ B > G | H = 60° x [6 - (BG)/(RG)]

// aboe framwork for calculating hue
func calculateHue(c color.RGBA) float64 {
	r := float64(c.R) / 255.0
	g := float64(c.G) / 255.0
	b := float64(c.B) / 255.0
	// pixelColor := gocolor.Color{R: float64(c.R), G: float64(c.G), B: float64(c.B)} // simpler way but with a dependency produces the same result
	// hue, _, _ := pixelColor.Hsl()

	if c.R == c.B && c.G == c.B && c.R == c.G { // has tto be here for some reson
		// clog.Info("equal") 					// if this is not here the function doesn't detect equals 💀
		return 0.0
	}

	// max, min := getMaxMin(c)

	// clog.Info(max)
	// clog.Info(min)

	var hue float64

	switch {
	case r >= g && g >= b: // A
		hue = 60 * ((g - b) / (r - b))
	// clog.Info("A")
	case g > r && r >= b: // B
		hue = 60 * (2 - (r-b)/(g-b))
	// clog.Info("B")
	case g >= b && b > r: // C
		hue = 60 * (2 + (b-r)/(g-r))
	// clog.Info("C")
	case b > g && g > r: // D
		hue = 60 * (4 - (g-r)/(b-r))
	// clog.Info("D")
	case b > r && r >= g: // E
		hue = 60 * (4 + (r-g)/(b-g))
	// clog.Info("E")
	case r >= b && b > g: // F
		hue = 60 * (6 - (b-g)/(r-g))
		// clog.Info("F")
	}

	// if math.IsNaN(hue) {
	// 	clog.Warn("Hue is NaN")
	// 	hue = 0
	// }
	// fmt.Println("pass start: ", "r: ", c.R, "g: ", c.G, "b: ", c.B, "Is nan: ", math.IsNaN(hue), "end of pass")
	// clog.Info(hue)

	if hue < 0 {
		clog.Warn("Hue is negative")
	}

	if hue > 360 {
		clog.Warn("Hue is greater than 360")
	}

	// panic("implement me")
	return hue
}

func calculateLuminosity(c color.RGBA) float64 {
	max, min := getMaxMin(c)

	L := (1.0 / 2.0) * (max + min)

	return L
}

func calculateSaturation(c color.RGBA) float64 {
	max, min := getMaxMin(c)
	L := calculateLuminosity(c)
	var S float64

	if L < 1.0 {
		if (1.0 - math.Abs(2.0*L-1.0)) == 0.0 {
			S = 0.0
		} else {
			S = (max - min) / (1.0 - math.Abs(2.0*L-1.0))
		}
	} else {
		S = 0.0
	}

	return S
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
