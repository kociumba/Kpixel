package main

import (
	"image/color"
	"math"
	"testing"
)

func Test_calculateHue(t *testing.T) {
	type args struct {
		c color.RGBA
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Red dominant",
			args: args{c: color.RGBA{R: 201, G: 25, B: 75, A: 255}},
			want: 342.96,
		},
		{
			name: "equal",
			args: args{c: color.RGBA{R: 255, G: 255, B: 255, A: 255}},
			want: 0.0,
		},
		{
			name: "qual 0",
			args: args{c: color.RGBA{R: 0, G: 0, B: 0, A: 255}},
			want: 0.0,
		},
		// Add more test cases here
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateHue(tt.args.c); !almostEqual(got, tt.want, 0.01) {
				t.Errorf("calculateHue() = %v, want %v", got, tt.want)
			}
		})
	}
}

// almostEqual helps in comparing float values with a tolerance to handle precision errors.
func almostEqual(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}
