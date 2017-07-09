package image

import "fmt"

var (
	// White color
	White = Color{R: 1, G: 1, B: 1}
	// Black color
	Black = Color{R: 0, G: 0, B: 0}
	// Red color
	Red = Color{R: 1, G: 0, B: 0}
	// Green color
	Green = Color{R: 0, G: 1, B: 0}
	// Blue color
	Blue = Color{R: 0, G: 0, B: 1}
	// Magenta color
	Magenta = Color{R: 1, G: 0, B: 1}
	// Yellow color
	Yellow = Color{R: 1, G: 1, B: 0}
	// Cyan color
	Cyan = Color{R: 0, G: 1, B: 1}
)

// Color defines an RGB color with floating point precision
type Color struct {
	R float64 `json:"r"`
	G float64 `json:"g"`
	B float64 `json:"b"`
}

// Add returns a color result of adding both colors elementwise
func (c *Color) Add(c2 *Color) *Color {
	return &Color{R: c.R + c2.R, G: c.G + c2.G, B: c.B + c2.B}
}

// Multiply returns a color result of multiplying this color by a float
func (c *Color) Multiply(v float64) *Color {
	return &Color{R: c.R * v, G: c.G * v, B: c.B * v}
}

// Divide returns a color result of dividing this color by a float
func (c *Color) Divide(v float64) *Color {
	return &Color{R: c.R / v, G: c.G / v, B: c.B / v}
}

// CMultiply returns a color result of multiplying both colors
// by channel
func (c *Color) CMultiply(c2 *Color) *Color {
	return &Color{R: c.R * c2.R, G: c.G * c2.G, B: c.B * c2.B}
}

// ColorFromMap returns the color defined in the map
func ColorFromMap(m map[string]float64) Color {
	return Color{R: m["r"], G: m["g"], B: m["b"]}
}

func (c *Color) String() string {
	return fmt.Sprintf("[R: %.2f, G: %.2f, B: %.2f]", c.R, c.G, c.B)
}
