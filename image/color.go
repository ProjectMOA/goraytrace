package image

// Color defines an RGB color with floating point precision
type Color struct {
	R float64 `json:"r"`
	G float64 `json:"g"`
	B float64 `json:"b"`
}

// ColorFromMap returns the color defined in the map
func ColorFromMap(m map[string]float64) Color {
	return Color{R: m["r"], G: m["g"], B: m["b"]}
}
