package colors

import (
	"fmt"
	"math"
)

type HSL struct {
	H float64 // 0–360
	S float64 // 0–1
	L float64 // 0–1
}

func HSLToHex(hsl HSL) string {
	h := hsl.H
	s := hsl.S
	l := hsl.L

	c := (1 - math.Abs(2*l-1)) * s
	x := c * (1 - math.Abs(math.Mod(h/60, 2)-1))
	m := l - c/2

	var r, g, b float64

	switch {
	case h < 60:
		r, g, b = c, x, 0
	case h < 120:
		r, g, b = x, c, 0
	case h < 180:
		r, g, b = 0, c, x
	case h < 240:
		r, g, b = 0, x, c
	case h < 300:
		r, g, b = x, 0, c
	default:
		r, g, b = c, 0, x
	}

	r = (r + m) * 255
	g = (g + m) * 255
	b = (b + m) * 255

	return fmt.Sprintf("#%02x%02x%02x", clampInt(r), clampInt(g), clampInt(b))
}

func clampInt(v float64) int {
	if v < 0 {
		return 0
	}
	if v > 255 {
		return 255
	}
	return int(math.Round(v))
}
