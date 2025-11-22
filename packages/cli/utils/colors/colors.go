package colors

import (
	"fmt"
	"math"
	"math/rand/v2"
	"strconv"
	"strings"
)

// HexToRgb converts a hex color string to RGB values
func HexToRgb(hex string) (int, int, int, error) {
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) == 3 {
		hex = fmt.Sprintf("%c%c%c%c%c%c", hex[0], hex[0], hex[1], hex[1], hex[2], hex[2])
	}

	r, err := strconv.ParseInt(hex[0:2], 16, 0)
	if err != nil {
		return 0, 0, 0, err
	}
	g, err := strconv.ParseInt(hex[2:4], 16, 0)
	if err != nil {
		return 0, 0, 0, err
	}
	b, err := strconv.ParseInt(hex[4:6], 16, 0)
	if err != nil {
		return 0, 0, 0, err
	}

	return int(r), int(g), int(b), nil
}

// GenerateRandomHexColor generates a random hex color string using math/rand/v2.
func GenerateRandomHexColor() string {
	r := rand.Uint32N(256)
	g := rand.Uint32N(256)
	b := rand.Uint32N(256)
	return fmt.Sprintf("#%02X%02X%02X", r, g, b)
}

// HexToHSL converts a hex color string to HSL
func HexToHSL(hex string) (h, s, l float64, err error) {
	rInt, gInt, bInt, err := HexToRgb(hex)
	if err != nil {
		return 0, 0, 0, err
	}

	r := float64(rInt) / 255
	g := float64(gInt) / 255
	b := float64(bInt) / 255

	max := math.Max(r, math.Max(g, b))
	min := math.Min(r, math.Min(g, b))
	l = (max + min) / 2

	if max == min {
		h, s = 0, 0
	} else {
		d := max - min
		if l > 0.5 {
			s = d / (2 - max - min)
		} else {
			s = d / (max + min)
		}

		switch max {
		case r:
			h = (g - b) / d
			if g < b {
				h += 6
			}
		case g:
			h = (b-r)/d + 2
		case b:
			h = (r-g)/d + 4
		}
		h *= 60
	}

	return h, s * 100, l * 100, nil
}

// HexToOKLCH converts a hex color string to OKLCH
func HexToOKLCH(hex string) (l, c, h float64, err error) {
	rInt, gInt, bInt, err := HexToRgb(hex)
	if err != nil {
		return 0, 0, 0, err
	}

	// Convert sRGB to linear RGB
	toLinear := func(c float64) float64 {
		c = c / 255
		if c <= 0.04045 {
			return c / 12.92
		}
		return math.Pow((c+0.055)/1.055, 2.4)
	}
	r := toLinear(float64(rInt))
	g := toLinear(float64(gInt))
	b := toLinear(float64(bInt))

	// Linear RGB to LMS (Oklab)
	l_ := 0.4122214708*r + 0.5363325363*g + 0.0514459929*b
	m_ := 0.2119034982*r + 0.6806995451*g + 0.1073969566*b
	s_ := 0.0883024619*r + 0.2817188376*g + 0.6299787005*b

	cbrt := func(x float64) float64 { return math.Cbrt(x) }

	L := 0.2104542553*cbrt(l_) + 0.7936177850*cbrt(m_) - 0.0040720468*cbrt(s_)
	A := 1.9779984951*cbrt(l_) - 2.4285922050*cbrt(m_) + 0.4505937099*cbrt(s_)
	B := 0.0259040371*cbrt(l_) + 0.7827717662*cbrt(m_) - 0.8086757660*cbrt(s_)

	c = math.Sqrt(A*A + B*B)
	h = math.Atan2(B, A) * (180 / math.Pi)
	if h < 0 {
		h += 360
	}

	return L, c, h, nil
}

// HexToHCL converts a hex color string to HCL (Hue, Chroma, Lightness)
func HexToHCL(hex string) (h, c, l float64, err error) {
	// Convert HEX → RGB
	r, g, b, err := HexToRgb(hex)
	if err != nil {
		return 0, 0, 0, err
	}

	// Convert RGB to [0,1]
	R := float64(r) / 255
	G := float64(g) / 255
	B := float64(b) / 255

	// sRGB → Linear RGB
	toLinear := func(u float64) float64 {
		if u <= 0.04045 {
			return u / 12.92
		}
		return math.Pow((u+0.055)/1.055, 2.4)
	}
	R = toLinear(R)
	G = toLinear(G)
	B = toLinear(B)

	// Linear RGB → XYZ
	X := 0.4124564*R + 0.3575761*G + 0.1804375*B
	Y := 0.2126729*R + 0.7151522*G + 0.0721750*B
	Z := 0.0193339*R + 0.1191920*G + 0.9503041*B

	// XYZ → Lab
	refX, refY, refZ := 0.95047, 1.00000, 1.08883
	f := func(t float64) float64 {
		if t > 0.008856 {
			return math.Cbrt(t)
		}
		return 7.787*t + 16.0/116
	}
	fX := f(X / refX)
	fY := f(Y / refY)
	fZ := f(Z / refZ)

	L := 116*fY - 16
	a := 500 * (fX - fY)
	bb := 200 * (fY - fZ)

	c = math.Sqrt(a*a + bb*bb)
	h = math.Atan2(bb, a) * (180 / math.Pi)
	if h < 0 {
		h += 360
	}
	l = L

	return h, c, l, nil
}
